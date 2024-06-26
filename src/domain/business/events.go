package business

import (
	"fmt"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.business/src/config"
	"github.com/turistikrota/service.business/src/domain/notify"
)

type Events interface {
	Created(event *EventBusinessCreated)
	UserAdded(event *EventBusinessUserAdded)
	UserRemoved(event *EventBusinessUserRemoved)
	UserPermissionRemoved(event *EventBusinessPermissionRemoved)
	UserPermissionAdded(event *EventBusinessPermissionAdded)
	VerifiedByAdmin(event *EventBusinessVerifiedByAdmin)
	DeletedByAdmin(event *EventBusinessDeletedByAdmin)
	RecoverByAdmin(event *EventBusinessRecoverByAdmin)
	RejectedByAdmin(event *EventBusinessRejectedByAdmin)
	Disabled(event *EventBusinessDisabled)
	Enabled(event *EventBusinessEnabled)
}

type (
	EventBusinessCreated struct {
		Business *Entity `json:"business"`
		UserUUID string  `json:"userUUID"`
		UserName string  `json:"userName"`
	}
	EventUser struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}
	EventBusinessUserAdded struct {
		BusinessNickName string `json:"businessNickName"`
		User             *User  `json:"user"`
	}
	EventBusinessUserRemoved struct {
		BusinessNickName string    `json:"businessNickName"`
		AccessUserName   string    `json:"accessUserName"`
		AccessUserUUID   string    `json:"accessUserUUID"`
		User             EventUser `json:"user"`
	}
	EventBusinessPermissionRemoved struct {
		BusinessNickName string    `json:"businessNickName"`
		AccessUserUUID   string    `json:"accessUserUUID"`
		AccessUserName   string    `json:"accessUserName"`
		User             EventUser `json:"user"`
		PermissionName   string    `json:"permission"`
	}
	EventBusinessPermissionAdded struct {
		BusinessNickName string    `json:"businessNickName"`
		AccessUserUUID   string    `json:"accessUserUUID"`
		AccessUserName   string    `json:"accessUserName"`
		User             EventUser `json:"user"`
		PermissionName   string    `json:"permission"`
	}
	EventBusinessVerifiedByAdmin struct {
		BusinessNickName string   `json:"businessNickName"`
		BusinessLocale   string   `json:"businessLocale"`
		Users            []string `json:"users"`
		AdminUUID        string   `json:"adminUUID"`
	}
	EventBusinessRejectedByAdmin struct {
		BusinessNickName string   `json:"businessNickName"`
		Reason           string   `json:"reason"`
		BusinessLocale   string   `json:"businessLocale"`
		Users            []string `json:"users"`
		AdminUUID        string   `json:"adminUUID"`
	}
	EventBusinessDeletedByAdmin struct {
		BusinessNickName string `json:"businessNickName"`
		AdminUUID        string `json:"adminUUID"`
	}
	EventBusinessRecoverByAdmin struct {
		BusinessNickName string `json:"businessNickName"`
		AdminUUID        string `json:"adminUUID"`
	}
	EventBusinessDisabled struct {
		UserName         string `json:"nickName"`
		UserUUID         string `json:"userUUID"`
		UserCode         string `json:"userCode"`
		BusinessNickName string `json:"businessNickName"`
	}
	EventBusinessEnabled struct {
		UserName         string `json:"nickName"`
		UserUUID         string `json:"userUUID"`
		UserCode         string `json:"userCode"`
		BusinessNickName string `json:"businessNickName"`
	}
)

type businessEvents struct {
	publisher events.Publisher
	topics    config.Topics
	i18n      *i18np.I18n
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
	I18n      *i18np.I18n
}

func NewEvents(config EventConfig) Events {
	return &businessEvents{
		publisher: config.Publisher,
		topics:    config.Topics,
		i18n:      config.I18n,
	}
}

func (e *businessEvents) Created(event *EventBusinessCreated) {
	_ = e.publisher.Publish(e.topics.Business.Created, event)
}

func (e *businessEvents) UserRemoved(event *EventBusinessUserRemoved) {
	_ = e.publisher.Publish(e.topics.Business.UserRemoved, event)
}

func (e *businessEvents) UserAdded(event *EventBusinessUserAdded) {
	_ = e.publisher.Publish(e.topics.Business.UserAdded, event)
}

func (e *businessEvents) UserPermissionRemoved(event *EventBusinessPermissionRemoved) {
	_ = e.publisher.Publish(e.topics.Business.UserPermissionRemoved, event)
}

func (e *businessEvents) UserPermissionAdded(event *EventBusinessPermissionAdded) {
	_ = e.publisher.Publish(e.topics.Business.UserPermissionAdded, event)
}

func (e *businessEvents) VerifiedByAdmin(event *EventBusinessVerifiedByAdmin) {
	_ = e.publisher.Publish(e.topics.Business.VerifiedByAdmin, event)
	if event.BusinessLocale == "" {
		event.BusinessLocale = "tr"
	}
	subject := e.i18n.Translate(I18nMessages.NotifySubjectVerified, event.BusinessLocale)
	smsContent := fmt.Sprintf(e.i18n.Translate(I18nMessages.NotifyVerifiedContent, event.BusinessLocale), event.BusinessNickName)
	template := fmt.Sprintf("business/verified.%s", event.BusinessLocale)
	for _, user := range event.Users {
		_ = e.publisher.Publish(e.topics.Notify.SendNotification, notify.NotifySendToAllChannelsCmd{
			ActorName: user,
			Content:   smsContent,
			TemplateData: i18np.P{
				"BusinessName": event.BusinessNickName,
			},
			Template:  template,
			Subject:   subject,
			Locale:    event.BusinessLocale,
			Translate: false,
		})
	}
}

func (e *businessEvents) DeletedByAdmin(event *EventBusinessDeletedByAdmin) {
	_ = e.publisher.Publish(e.topics.Business.DeletedByAdmin, event)
}

func (e *businessEvents) RejectedByAdmin(event *EventBusinessRejectedByAdmin) {
	_ = e.publisher.Publish(e.topics.Business.RejectedByAdmin, event)
	if event.BusinessLocale == "" {
		event.BusinessLocale = "tr"
	}
	subject := e.i18n.Translate(I18nMessages.NotifySubjectRejected, event.BusinessLocale)
	smsContent := fmt.Sprintf(e.i18n.Translate(I18nMessages.NotifyRejectContent, event.BusinessLocale), event.BusinessNickName)
	template := fmt.Sprintf("business/rejected.%s", event.BusinessLocale)
	for _, user := range event.Users {
		_ = e.publisher.Publish(e.topics.Notify.SendNotification, notify.NotifySendToAllChannelsCmd{
			ActorName: user,
			Content:   smsContent,
			TemplateData: i18np.P{
				"BusinessName": event.BusinessNickName,
				"Reason":       event.Reason,
			},
			Template:  template,
			Subject:   subject,
			Locale:    event.BusinessLocale,
			Translate: false,
		})
	}
}

func (e *businessEvents) RecoverByAdmin(event *EventBusinessRecoverByAdmin) {
	_ = e.publisher.Publish(e.topics.Business.RecoverByAdmin, event)
}

func (e *businessEvents) Disabled(event *EventBusinessDisabled) {
	_ = e.publisher.Publish(e.topics.Business.Disabled, event)
}

func (e *businessEvents) Enabled(event *EventBusinessEnabled) {
	_ = e.publisher.Publish(e.topics.Business.Enabled, event)
}
