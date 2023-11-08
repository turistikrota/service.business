package invite

import (
	"fmt"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.owner/src/config"
	"github.com/turistikrota/service.shared/helper"
)

type Events interface {
	Invite(event InviteEvent)
	Delete(event InviteDeleteEvent)
	Use(event InviteUseEvent)
}

type (
	InviteEvent struct {
		Locale     string `json:"locale"`
		Email      string `json:"email"`
		InviteUUID string `json:"inviteUUID"`
		OwnerName  string `json:"ownerName"`
		UserUUID   string `json:"userUUID"`
		UserName   string `json:"userName"`
	}
	InviteDeleteEvent struct {
		InviteUUID string `json:"inviteUUID"`
		UserUUID   string `json:"userUUID"`
		UserName   string `json:"userName"`
	}
	InviteUseEvent struct {
		InviteUUID string `json:"inviteUUID"`
		UserEmail  string `json:"userEmail"`
		UserUUID   string `json:"userUUID"`
		UserName   string `json:"userName"`
	}
)

type inviteEvents struct {
	publisher events.Publisher
	topics    config.Topics
	i18n      *i18np.I18n
	urls      config.Urls
}

type EventConfig struct {
	Publisher events.Publisher
	Topics    config.Topics
	Urls      config.Urls
	I18n      *i18np.I18n
}

func NewEvents(cnf EventConfig) Events {
	return &inviteEvents{
		publisher: cnf.Publisher,
		topics:    cnf.Topics,
		urls:      cnf.Urls,
		i18n:      cnf.I18n,
	}
}

func (e inviteEvents) Invite(event InviteEvent) {
	subject := e.i18n.Translate(I18nMessages.InviteSubject, event.Locale)
	template := fmt.Sprintf("owner/invite.%s", event.Locale)
	_ = e.publisher.Publish(e.topics.Notify.SendMail, helper.Notify.BuildEmail(event.Email, subject, i18np.P{
		"OwnerName":  event.OwnerName,
		"InviteUUID": event.InviteUUID,
	}, event.Email, template))
	_ = e.publisher.Publish(e.topics.Owner.InviteCreate, event)
}

func (e inviteEvents) Delete(event InviteDeleteEvent) {
	_ = e.publisher.Publish(e.topics.Owner.InviteDelete, event)
}

func (e inviteEvents) Use(event InviteUseEvent) {
	_ = e.publisher.Publish(e.topics.Owner.InviteUse, event)
}
