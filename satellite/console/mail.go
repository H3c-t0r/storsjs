// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information

package console

import (
	"net/mail"
)

const (
	// ActivationSubject activation email subject
	ActivationSubject = "Activate your email"
	// InvitationSubject invitation email subject
	InvitationSubject = ""
	// ForgotPasswordSubject forgot password email subject
	ForgotPasswordSubject = ""
)

// MailTemplate is implementation of satellite/mailservice.Template interface
type MailTemplate struct {
	to            mail.Address
	subject       string
	htmlPath      string
	plainTextPath string
}

// NewMailTemplate creates new instance of MailTemplate
func NewMailTemplate(to mail.Address, subject, prefix string) MailTemplate {
	return MailTemplate{
		to:            to,
		subject:       subject,
		htmlPath:      prefix + ".html",
		plainTextPath: prefix + ".txt",
	}
}

// To gets recipients mailservice addresses
func (tmpl *MailTemplate) To() []mail.Address {
	return []mail.Address{tmpl.to}
}

// Subject gets email subject
func (tmpl *MailTemplate) Subject() string {
	return tmpl.subject
}

// HTMLPath gets path to html template
func (tmpl *MailTemplate) HTMLPath() string {
	return tmpl.htmlPath
}

// PainTextPath gets path to text template
func (tmpl *MailTemplate) PainTextPath() string {
	return tmpl.plainTextPath
}

// AccountActivationEmail is mailservice template with activation data
type AccountActivationEmail struct {
	MailTemplate
	ActivationLink string
}

// ForgotPasswordEmail is mailservice template with reset password data
type ForgotPasswordEmail struct {
	MailTemplate
	UserName  string
	ResetLink string
}

// ProjectInvitationEmail is mailservice template for project invitation email
type ProjectInvitationEmail struct {
	MailTemplate
	UserName    string
	ProjectName string
}
