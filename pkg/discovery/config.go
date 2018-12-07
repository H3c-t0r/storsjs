package discovery

import (
	"context"
	"time"

	"github.com/zeebo/errs"
	"go.uber.org/zap"
	monkit "gopkg.in/spacemonkeygo/monkit.v2"

	"storj.io/storj/pkg/kademlia"
	"storj.io/storj/pkg/overlay"
	"storj.io/storj/pkg/pb"
	"storj.io/storj/pkg/provider"
	"storj.io/storj/pkg/statdb"
)

var (
	mon = monkit.Package()
	// Error represents an overlay error
	Error = errs.Class("discovery error")
)

// CtxKey used for assigning a key to Discovery server
type CtxKey int

const (
	ctxKeyDiscovery CtxKey = iota
)

// Config loads on the configuration values from run flags
type Config struct {
	RefreshInterval time.Duration `help:"the interval at which the cache refreshes itself in seconds" default:"1s"`
}

// Run runs the Discovery boot up and initialization
func (c Config) Run(ctx context.Context, server *provider.Provider) (err error) {
	defer mon.Task()(&ctx)(&err)

	srv := NewServer(zap.L())
	pb.RegisterDiscoveryServer(server.GRPC(), srv)

	ol := overlay.LoadFromContext(ctx)
	kad := kademlia.LoadFromContext(ctx)
	stat := statdb.LoadFromContext(ctx)
	discovery := NewDiscovery(ol, kad, stat)

	return server.Run(context.WithValue(ctx, ctxKeyDiscovery, discovery))
}
