package business

import (
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.business/src/config"
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
		BusinessUUID string `json:"businessUUID"`
		User         *User  `json:"user"`
	}
	EventBusinessUserRemoved struct {
		BusinessUUID   string    `json:"businessUUID"`
		AccessUserName string    `json:"accessUserName"`
		AccessUserUUID string    `json:"accessUserUUID"`
		User           EventUser `json:"user"`
	}
	EventBusinessPermissionRemoved struct {
		BusinessUUID   string    `json:"businessUUID"`
		AccessUserUUID string    `json:"accessUserUUID"`
		AccessUserName string    `json:"accessUserName"`
		User           EventUser `json:"user"`
		PermissionName string    `json:"permission"`
	}
	EventBusinessPermissionAdded struct {
		BusinessUUID   string    `json:"businessUUID"`
		AccessUserUUID string    `json:"accessUserUUID"`
		AccessUserName string    `json:"accessUserName"`
		User           EventUser `json:"user"`
		PermissionName string    `json:"permission"`
	}
	EventBusinessVerifiedByAdmin struct {
		BusinessUUID string `json:"businessUUID"`
		AdminUUID    string `json:"adminUUID"`
	}
	EventBusinessRejectedByAdmin struct {
		BusinessUUID string `json:"businessUUID"`
		Reason       string `json:"reason"`
		AdminUUID    string `json:"adminUUID"`
	}
	EventBusinessDeletedByAdmin struct {
		BusinessUUID string `json:"businessUUID"`
		AdminUUID    string `json:"adminUUID"`
	}
	EventBusinessRecoverByAdmin struct {
		BusinessUUID string `json:"businessUUID"`
		AdminUUID    string `json:"adminUUID"`
	}
	EventBusinessDisabled struct {
		UserName     string `json:"nickName"`
		UserUUID     string `json:"userUUID"`
		UserCode     string `json:"userCode"`
		BusinessUUID string `json:"businessUUID"`
	}
	EventBusinessEnabled struct {
		UserName     string `json:"nickName"`
		UserUUID     string `json:"userUUID"`
		UserCode     string `json:"userCode"`
		BusinessUUID string `json:"businessUUID"`
	}
)

type businessEvents struct {
	publisher events.Publisher
	topics    config.Topics
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
}

func NewEvents(config EventConfig) Events {
	return &businessEvents{
		publisher: config.Publisher,
		topics:    config.Topics,
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
}

func (e *businessEvents) DeletedByAdmin(event *EventBusinessDeletedByAdmin) {
	_ = e.publisher.Publish(e.topics.Business.DeletedByAdmin, event)
}

func (e *businessEvents) RejectedByAdmin(event *EventBusinessRejectedByAdmin) {
	_ = e.publisher.Publish(e.topics.Business.RejectedByAdmin, event)
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
