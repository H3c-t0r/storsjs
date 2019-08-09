// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package consoleweb

import (
	"context"
	"encoding/json"
	"fmt"
	"html/template"
	"mime"
	"net"
	"net/http"
	"net/url"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sync/errgroup"
	"github.com/graphql-go/graphql"
	"github.com/skyrings/skyring-common/tools/uuid"
	"github.com/zeebo/errs"
	"go.uber.org/zap"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/auth"
	"storj.io/storj/satellite/console"
	"storj.io/storj/satellite/console/consoleweb/consoleql"
	"storj.io/storj/satellite/mailservice"
)

const (
	authorization = "Authorization"
	contentType   = "Content-Type"

	authorizationBearer = "Bearer "

	applicationJSON    = "application/json"
	applicationGraphql = "application/graphql"
)

var (
	// Error is satellite console error type
	Error = errs.Class("satellite console error")

	mon = monkit.Package()
)

// Config contains configuration for console web server
type Config struct {
	Address         string `help:"server address of the graphql api gateway and frontend app" devDefault:"127.0.0.1:8081" releaseDefault:":10100"`
	StaticDir       string `help:"path to static resources" default:""`
	ExternalAddress string `help:"external endpoint of the satellite if hosted" default:""`
	StripeKey       string `help:"stripe api key" default:""`

	// TODO: remove after Vanguard release
	AuthToken       string `help:"auth token needed for access to registration token creation endpoint" default:""`
	AuthTokenSecret string `help:"secret used to sign auth tokens" releaseDefault:"" devDefault:"my-suppa-secret-key"`

	PasswordCost int `internal:"true" help:"password hashing cost (0=automatic)" default:"0"`

	CustomerSupportLink string `help:"link that leads to some resource for customer support" devDefault:"https://storjlabs.atlassian.net/servicedesk/customer/portals" releaseDefault:""`
}

// consoleTemplate - is used to enumerate template list used in server
type consoleTemplate int

const (
	notFound             consoleTemplate = 0
	usageReport          consoleTemplate = 1
	resetPassword        consoleTemplate = 2
	resetPasswordSuccess consoleTemplate = 3
	accountActivated     consoleTemplate = 4
	index                consoleTemplate = 5
)

// Server represents console web server
type Server struct {
	log *zap.Logger

	config      Config
	service     *console.Service
	mailService *mailservice.Service

	listener net.Listener
	server   http.Server

	schema    graphql.Schema
	templates map[consoleTemplate]*template.Template
}

// NewServer creates new instance of console server
func NewServer(logger *zap.Logger, config Config, service *console.Service, mailService *mailservice.Service, listener net.Listener) *Server {
	server := Server{
		log:         logger,
		config:      config,
		listener:    listener,
		service:     service,
		mailService: mailService,
	}

	logger.Sugar().Debugf("Starting Satellite UI on %s...", server.listener.Addr().String())

	if server.config.ExternalAddress != "" {
		if !strings.HasSuffix(server.config.ExternalAddress, "/") {
			server.config.ExternalAddress += "/"
		}
	} else {
		server.config.ExternalAddress = "http://" + server.listener.Addr().String() + "/"
	}

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir(server.config.StaticDir))

	mux.Handle("/api/graphql/v0", http.HandlerFunc(server.grapqlHandler))

	if server.config.StaticDir != "" {
		mux.Handle("/activation/", http.HandlerFunc(server.accountActivationHandler))
		mux.Handle("/password-recovery/", http.HandlerFunc(server.passwordRecoveryHandler))
		mux.Handle("/cancel-password-recovery/", http.HandlerFunc(server.cancelPasswordRecoveryHandler))
		mux.Handle("/registrationToken/", http.HandlerFunc(server.createRegistrationTokenHandler))
		mux.Handle("/usage-report/", http.HandlerFunc(server.bucketUsageReportHandler))
		mux.Handle("/static/", server.gzipHandler(http.StripPrefix("/static", fs)))
		mux.Handle("/", http.HandlerFunc(server.appHandler))
	}

	server.server = http.Server{
		Handler: mux,
	}

	return &server
}

// Run starts the server that host webapp and api endpoint
func (server *Server) Run(ctx context.Context) (err error) {
	defer mon.Task()(&ctx)(&err)

	server.schema, err = consoleql.CreateSchema(server.log, server.service, server.mailService)
	if err != nil {
		return Error.Wrap(err)
	}

	err = server.initializeTemplates()
	if err != nil {
		// TODO: should it return error if some template can not be initialized or just log about it?
		return Error.Wrap(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	var group errgroup.Group
	group.Go(func() error {
		<-ctx.Done()
		return server.server.Shutdown(nil)
	})
	group.Go(func() error {
		defer cancel()
		return server.server.Serve(server.listener)
	})

	return group.Wait()
}

// Close closes server and underlying listener
func (server *Server) Close() error {
	return server.server.Close()
}

// appHandler is web app http handler function
func (server *Server) appHandler(w http.ResponseWriter, r *http.Request) {
	header := w.Header()

	cspValues := []string{
		"default-src 'self'",
		"script-src 'self' *.stripe.com cdn.segment.com",
		"frame-src 'self' *.stripe.com",
		"img-src 'self' data:",
	}

	header.Set("Content-Type", "text/html; charset=UTF-8")
	header.Set("Content-Security-Policy", strings.Join(cspValues, "; "))

	if err := server.executeTemplate(w, r, index, nil); err != nil {
		server.serveError(w, r, http.StatusNotFound)
	}
}

// bucketUsageReportHandler generate bucket usage report page for project
func (server *Server) bucketUsageReportHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var err error
	defer mon.Task()(&ctx)(&err)

	var projectID *uuid.UUID
	var since, before time.Time

	tokenCookie, err := r.Cookie("tokenKey")
	if err != nil {
		server.log.Error("bucket usage report error", zap.Error(err))

		w.WriteHeader(http.StatusUnauthorized)
		http.ServeFile(w, r, filepath.Join(server.config.StaticDir, "static", "errors", "404.html"))
		return
	}

	auth, err := server.service.Authorize(auth.WithAPIKey(ctx, []byte(tokenCookie.Value)))
	if err != nil {
		server.log.Error("bucket usage report error", zap.Error(err))

		//TODO: when new error pages will be created - change http.StatusNotFound on http.StatusUnauthorized
		server.serveError(w, r, http.StatusNotFound)
		return
	}

	defer func() {
		if err != nil {
			server.log.Error("bucket usage report error", zap.Error(err))

			w.WriteHeader(http.StatusNotFound)
			http.ServeFile(w, r, filepath.Join(server.config.StaticDir, "static", "errors", "404.html"))
		}
	}()

	// parse query params
	projectID, err = uuid.Parse(r.URL.Query().Get("projectID"))
	if err != nil {
		return
	}
	sinceStamp, err := strconv.ParseInt(r.URL.Query().Get("since"), 10, 64)
	if err != nil {
		return
	}
	beforeStamp, err := strconv.ParseInt(r.URL.Query().Get("before"), 10, 64)
	if err != nil {
		return
	}

	since = time.Unix(sinceStamp, 0)
	before = time.Unix(beforeStamp, 0)

	server.log.Debug("querying bucket usage report",
		zap.Stringer("projectID", projectID),
		zap.Stringer("since", since),
		zap.Stringer("before", before))

	ctx = console.WithAuth(ctx, auth)
	bucketRollups, err := server.service.GetBucketUsageRollups(ctx, *projectID, since, before)
	if err != nil {
		return
	}

	if err = server.executeTemplate(w, r, usageReport, bucketRollups); err != nil {
		server.serveError(w, r, http.StatusNotFound)
	}
}

// accountActivationHandler is web app http handler function
func (server *Server) createRegistrationTokenHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer mon.Task()(&ctx)(nil)
	w.Header().Set(contentType, applicationJSON)

	var response struct {
		Secret string `json:"secret"`
		Error  string `json:"error,omitempty"`
	}

	defer func() {
		err := json.NewEncoder(w).Encode(&response)
		if err != nil {
			server.log.Error(err.Error())
		}
	}()

	authToken := r.Header.Get("Authorization")
	if authToken != server.config.AuthToken {
		w.WriteHeader(401)
		response.Error = "unauthorized"
		return
	}

	projectsLimitInput := r.URL.Query().Get("projectsLimit")

	projectsLimit, err := strconv.Atoi(projectsLimitInput)
	if err != nil {
		response.Error = err.Error()
		return
	}

	token, err := server.service.CreateRegToken(ctx, projectsLimit)
	if err != nil {
		response.Error = err.Error()
		return
	}

	response.Secret = token.Secret.String()
}

// accountActivationHandler is web app http handler function
func (server *Server) accountActivationHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer mon.Task()(&ctx)(nil)
	activationToken := r.URL.Query().Get("token")

	err := server.service.ActivateAccount(ctx, activationToken)
	if err != nil {
		server.log.Error("activation: failed to activate account",
			zap.String("token", activationToken),
			zap.Error(err))

		// TODO: when new error pages will be created - change http.StatusNotFound on appropriate one
		server.serveError(w, r, http.StatusNotFound)
		return
	}

	if err = server.executeTemplate(w, r, accountActivated, nil); err != nil {
		server.serveError(w, r, http.StatusNotFound)
	}
}

func (server *Server) passwordRecoveryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer mon.Task()(&ctx)(nil)

	recoveryToken := r.URL.Query().Get("token")
	if len(recoveryToken) == 0 {
		server.serveError(w, r, http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			server.serveError(w, r, http.StatusNotFound)
			return
		}

		password := r.FormValue("password")
		passwordRepeat := r.FormValue("passwordRepeat")
		if strings.Compare(password, passwordRepeat) != 0 {
			server.serveError(w, r, http.StatusNotFound)
			return
		}

		err = server.service.ResetPassword(ctx, recoveryToken, password)
		if err != nil {
			server.serveError(w, r, http.StatusNotFound)
			return
		}

		if err := server.executeTemplate(w, r, resetPasswordSuccess, nil); err != nil {
			server.serveError(w, r, http.StatusNotFound)
		}
	case http.MethodGet:
		if err := server.executeTemplate(w, r, resetPassword, nil); err != nil {
			server.serveError(w, r, http.StatusNotFound)
		}
	default:
		server.serveError(w, r, http.StatusNotFound)
		return
	}
}

func (server *Server) cancelPasswordRecoveryHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer mon.Task()(&ctx)(nil)
	recoveryToken := r.URL.Query().Get("token")

	// No need to check error as we anyway redirect user to support page
	_ = server.service.RevokeResetPasswordToken(ctx, recoveryToken)

	http.Redirect(w, r, "https://storjlabs.atlassian.net/servicedesk/customer/portals", http.StatusSeeOther)
}

func (server *Server) serveError(w http.ResponseWriter, r *http.Request, status int) {
	// TODO: show different error pages depend in status
	// F.e. switch(status)
	//      case http.StatusNotFound: server.executeTemplate(w, r, notFound, nil)
	//      case http.StatusInternalServerError: server.executeTemplate(w, r, internalError, nil)
	w.WriteHeader(status)

	if err := server.executeTemplate(w, r, notFound, nil); err != nil {
		server.log.Error("error occurred in console/server", zap.Error(err))
	}
}

// grapqlHandler is graphql endpoint http handler function
func (server *Server) grapqlHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	defer mon.Task()(&ctx)(nil)
	w.Header().Set(contentType, applicationJSON)

	token := getToken(r)
	query, err := getQuery(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx = auth.WithAPIKey(ctx, []byte(token))
	auth, err := server.service.Authorize(ctx)
	if err != nil {
		ctx = console.WithAuthFailure(ctx, err)
	} else {
		ctx = console.WithAuth(ctx, auth)
	}

	rootObject := make(map[string]interface{})

	rootObject["origin"] = server.config.ExternalAddress
	rootObject[consoleql.ActivationPath] = "activation/?token="
	rootObject[consoleql.PasswordRecoveryPath] = "password-recovery/?token="
	rootObject[consoleql.CancelPasswordRecoveryPath] = "cancel-password-recovery/?token="
	rootObject[consoleql.SignInPath] = "login"

	result := graphql.Do(graphql.Params{
		Schema:         server.schema,
		Context:        ctx,
		RequestString:  query.Query,
		VariableValues: query.Variables,
		OperationName:  query.OperationName,
		RootObject:     rootObject,
	})

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		server.log.Error(err.Error())
		return
	}

	sugar := server.log.Sugar()
	sugar.Debug(result)
}

// gzipHandler is used to gzip static content to minify resources if browser support such decoding
func (server *Server) gzipHandler(fn http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		isGzipSupported := strings.Contains(r.Header.Get("Accept-Encoding"), "gzip")
		extension := filepath.Ext(r.RequestURI)
		// we have gzipped only fonts, js and css bundles
		formats := map[string]bool{
			".js":  true,
			".ttf": true,
			".css": true,
		}
		isNeededFormatToGzip := formats[extension]

		// because we have some static content outside of console frontend app.
		// for example: 404 page, account activation, passsrowd reset, etc.
		// TODO: find better solution, its a temporary fix
		isFromStaticDir := strings.Contains(r.URL.Path, "/static/dist/")

		// in case if old browser doesn't support gzip decoding or if file extension is not recommended to gzip
		// just return original file
		if !isGzipSupported || !isNeededFormatToGzip || !isFromStaticDir {
			fn.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Type", mime.TypeByExtension(extension))
		w.Header().Set("Content-Encoding", "gzip")

		// updating request URL
		newRequest := new(http.Request)
		*newRequest = *r
		newRequest.URL = new(url.URL)
		*newRequest.URL = *r.URL
		newRequest.URL.Path += ".gz"

		fn.ServeHTTP(w, newRequest)
	})
}

// initializeTemplates is used to initialize all templates
func (server *Server) initializeTemplates() (err error) {
	server.templates = make(map[consoleTemplate]*template.Template)

	notFoundErr := server.parseTemplate(notFound, path.Join(server.config.StaticDir, "static", "errors", "404.html"))
	if notFoundErr != nil {
		err = errs.Combine(err, notFoundErr)
	}

	usageReportErr := server.parseTemplate(usageReport, path.Join(server.config.StaticDir, "static", "reports", "UsageReport.html"))
	if usageReportErr != nil {
		err = errs.Combine(err, usageReportErr)
	}

	resetPasswordErr := server.parseTemplate(resetPassword, filepath.Join(server.config.StaticDir, "static", "resetPassword", "resetPassword.html"))
	if resetPasswordErr != nil {
		err = errs.Combine(err, resetPasswordErr)
	}

	resetPasswordSuccessErr := server.parseTemplate(resetPasswordSuccess, filepath.Join(server.config.StaticDir, "static", "resetPassword", "success.html"))
	if resetPasswordSuccessErr != nil {
		err = errs.Combine(err, resetPasswordSuccessErr)
	}

	accountActivatedErr := server.parseTemplate(accountActivated, filepath.Join(server.config.StaticDir, "static", "activation", "success.html"))
	if accountActivatedErr != nil {
		err = errs.Combine(err, accountActivatedErr)
	}

	indexErr := server.parseTemplate(index, filepath.Join(server.config.StaticDir, "dist", "index.html"))
	if indexErr != nil {
		err = errs.Combine(err, indexErr)
	}

	return err
}

// parseTemplate is used to parse template and combine errors while parsing multiple
func (server *Server) parseTemplate(templateIndex consoleTemplate, templatePath string) error {
	template, err := template.ParseFiles(templatePath)
	if err != nil {
		return err
	}

	server.templates[templateIndex] = template

	return nil
}

// executeTemplate is used to find and execute specified template
func (server *Server) executeTemplate(w http.ResponseWriter, r *http.Request, template consoleTemplate, data interface{}) error {
	t, ok := server.templates[template]
	if !ok {
		return errs.New(fmt.Sprintf("can not find specified error by URI: %s", r.URL.Path))
	}

	if err := t.Execute(w, data); err != nil {
		return err
	}

	return nil
}
