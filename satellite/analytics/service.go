// Copyright (C) 2021 Storj Labs, Inc.
// See LICENSE for copying information.

package analytics

import (
	"context"

	"go.uber.org/zap"
	segment "gopkg.in/segmentio/analytics-go.v3"

	"storj.io/common/uuid"
)

const (
	eventAccountCreated            = "Account Created"
	eventSignedIn                  = "Signed In"
	eventProjectCreated            = "Project Created"
	eventAccessGrantCreated        = "Access Grant Created"
	eventAccountVerified           = "Account Verified"
	eventGatewayCredentialsCreated = "Credentials Created"
	eventPassphraseCreated         = "Passphrase Created"
	eventExternalLinkClicked       = "External Link Clicked"
	eventPathSelected              = "Path Selected"
	eventLinkShared                = "Link Shared"
)

// Config is a configuration struct for analytics Service.
type Config struct {
	SegmentWriteKey string `help:"segment write key" default:""`
	Enabled         bool   `help:"enable analytics reporting" default:"false"`
	HubSpot         HubSpotConfig
}

// Service for sending analytics.
//
// architecture: Service
type Service struct {
	log           *zap.Logger
	config        Config
	satelliteName string
	clientEvents  map[string]bool

	segment segment.Client
	hubspot *HubSpotEvents
}

// NewService creates new service for creating sending analytics.
func NewService(log *zap.Logger, config Config, satelliteName string) *Service {
	service := &Service{
		log:           log,
		config:        config,
		satelliteName: satelliteName,
		clientEvents:  make(map[string]bool),
		hubspot:       NewHubSpotEvents(config.HubSpot, satelliteName),
	}
	if config.Enabled {
		service.segment = segment.New(config.SegmentWriteKey)
	}
	for _, name := range []string{eventGatewayCredentialsCreated, eventPassphraseCreated, eventExternalLinkClicked, eventPathSelected, eventLinkShared} {
		service.clientEvents[name] = true
	}

	return service
}

// Run runs the service and use the context in new requests.
func (service *Service) Run(ctx context.Context) error {
	if !service.config.Enabled {
		return nil
	}

	return service.hubspot.Run(ctx)
}

// Close closes the Segment client.
func (service *Service) Close() error {
	if !service.config.Enabled {
		return nil
	}
	return service.segment.Close()
}

// UserType is a type for distinguishing personal vs. professional users.
type UserType string

const (
	// Professional defines a "professional" user type.
	Professional UserType = "Professional"
	// Personal defines a "personal" user type.
	Personal UserType = "Personal"
)

// TrackCreateUserFields contains input data for tracking a create user event.
type TrackCreateUserFields struct {
	ID               uuid.UUID
	AnonymousID      string
	FullName         string
	Email            string
	Type             UserType
	EmployeeCount    string
	CompanyName      string
	JobTitle         string
	HaveSalesContact bool
}

func (service *Service) enqueueMessage(message segment.Message) {
	if !service.config.Enabled {
		return
	}

	err := service.segment.Enqueue(message)
	if err != nil {
		service.log.Error("Error enqueueing message", zap.Error(err))
	}
}

// TrackCreateUser sends an "Account Created" event to Segment.
func (service *Service) TrackCreateUser(fields TrackCreateUserFields) {
	traits := segment.NewTraits()
	traits.SetName(fields.FullName)
	traits.SetEmail(fields.Email)

	service.enqueueMessage(segment.Identify{
		UserId:      fields.ID.String(),
		AnonymousId: fields.AnonymousID,
		Traits:      traits,
	})

	props := segment.NewProperties()
	props.Set("email", fields.Email)
	props.Set("name", fields.FullName)
	props.Set("satellite_selected", service.satelliteName)
	props.Set("account_type", fields.Type)

	if fields.Type == Professional {
		props.Set("company_size", fields.EmployeeCount)
		props.Set("company_name", fields.CompanyName)
		props.Set("job_title", fields.JobTitle)
		props.Set("have_sales_contact", fields.HaveSalesContact)
	}

	service.enqueueMessage(segment.Track{
		UserId:      fields.ID.String(),
		AnonymousId: fields.AnonymousID,
		Event:       eventAccountCreated,
		Properties:  props,
	})

	service.hubspot.EnqueueCreateUser(fields)

}

// TrackSignedIn sends an "Signed In" event to Segment.
func (service *Service) TrackSignedIn(userID uuid.UUID, email string) {
	traits := segment.NewTraits()
	traits.SetEmail(email)

	service.enqueueMessage(segment.Identify{
		UserId: userID.String(),
		Traits: traits,
	})

	props := segment.NewProperties()
	props.Set("email", email)

	service.enqueueMessage(segment.Track{
		UserId:     userID.String(),
		Event:      eventSignedIn,
		Properties: props,
	})

	if service.config.Enabled {
		service.hubspot.EnqueueEvent(email, eventSignedIn, map[string]interface{}{
			"userid": userID.String(),
		})
	}
}

// TrackProjectCreated sends an "Project Created" event to Segment.
func (service *Service) TrackProjectCreated(userID uuid.UUID, email string, projectID uuid.UUID, currentProjectCount int) {
	props := segment.NewProperties()
	props.Set("project_count", currentProjectCount)
	props.Set("project_id", projectID.String())

	service.enqueueMessage(segment.Track{
		UserId:     userID.String(),
		Event:      eventProjectCreated,
		Properties: props,
	})

	if service.config.Enabled {
		service.hubspot.EnqueueEvent(email, eventProjectCreated, map[string]interface{}{
			"userid":        userID.String(),
			"project_count": currentProjectCount,
			"project_id":    projectID.String(),
		})
	}
}

// TrackAccessGrantCreated sends an "Access Grant Created" event to Segment.
func (service *Service) TrackAccessGrantCreated(userID uuid.UUID, email string) {
	service.enqueueMessage(segment.Track{
		UserId: userID.String(),
		Event:  eventAccessGrantCreated,
	})

	if service.config.Enabled {
		service.hubspot.EnqueueEvent(email, eventAccessGrantCreated, map[string]interface{}{
			"userid": userID.String(),
		})
	}
}

// TrackAccountVerified sends an "Account Verified" event to Segment.
func (service *Service) TrackAccountVerified(userID uuid.UUID, email string) {
	traits := segment.NewTraits()
	traits.SetEmail(email)

	service.enqueueMessage(segment.Identify{
		UserId: userID.String(),
		Traits: traits,
	})

	props := segment.NewProperties()
	props.Set("email", email)

	service.enqueueMessage(segment.Track{
		UserId:     userID.String(),
		Event:      eventAccountVerified,
		Properties: props,
	})

	if service.config.Enabled {
		service.hubspot.EnqueueEvent(email, eventAccountVerified, map[string]interface{}{
			"userid": userID.String(),
		})
	}
}

// TrackEvent sends an arbitrary event associated with user ID to Segment.
// It is used for tracking occurrences of client-side events.
func (service *Service) TrackEvent(eventName string, userID uuid.UUID, email string) {
	// do not track if the event name is an invalid client-side event
	if !service.clientEvents[eventName] {
		service.log.Error("Invalid client-triggered event", zap.String("eventName", eventName))
		return
	}
	service.enqueueMessage(segment.Track{
		UserId: userID.String(),
		Event:  eventName,
	})

	if service.config.Enabled {
		service.hubspot.EnqueueEvent(email, eventName, map[string]interface{}{
			"userid": userID.String(),
		})
	}
}

// TrackLinkEvent sends an arbitrary event and link associated with user ID to Segment.
// It is used for tracking occurrences of client-side events.
func (service *Service) TrackLinkEvent(eventName string, userID uuid.UUID, email, link string) {
	// do not track if the event name is an invalid client-side event
	if !service.clientEvents[eventName] {
		service.log.Error("Invalid client-triggered event", zap.String("eventName", eventName))
		return
	}

	props := segment.NewProperties()
	props.Set("link", link)

	service.enqueueMessage(segment.Track{
		UserId:     userID.String(),
		Event:      eventName,
		Properties: props,
	})

	if service.config.Enabled {
		service.hubspot.EnqueueEvent(email, eventName, map[string]interface{}{
			"userid": userID.String(),
			"link":   link,
		})
	}
}
