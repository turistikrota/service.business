package owner

import (
	"github.com/mixarchitecture/microp/events"
	"github.com/turistikrota/service.owner/src/config"
)

type Events interface {
	Created(event *EventOwnerCreated)
	UserAdded(event *EventOwnerUserAdded)
	UserRemoved(event *EventOwnerUserRemoved)
	UserPermissionRemoved(event *EventOwnerPermissionRemoved)
	UserPermissionAdded(event *EventOwnerPermissionAdded)
	VerifiedByAdmin(event *EventOwnerVerifiedByAdmin)
	DeletedByAdmin(event *EventOwnerDeletedByAdmin)
	RecoverByAdmin(event *EventOwnerRecoverByAdmin)
	RejectedByAdmin(event *EventOwnerRejectedByAdmin)
	Disabled(event *EventOwnerDisabled)
	Enabled(event *EventOwnerEnabled)
}

type (
	EventOwnerCreated struct {
		Owner *Entity `json:"owner"`
	}
	EventUser struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}
	EventOwnerUserAdded struct {
		OwnerNickName string `json:"nickName"`
		User          *User  `json:"user"`
	}
	EventOwnerUserRemoved struct {
		OwnerNickName  string    `json:"nickName"`
		AccessUserUUID string    `json:"accessUserUUID"`
		User           EventUser `json:"user"`
	}
	EventOwnerPermissionRemoved struct {
		OwnerNickName  string    `json:"nickName"`
		AccessUserUUID string    `json:"accessUserUUID"`
		User           EventUser `json:"user"`
		PermissionName string    `json:"permission"`
	}
	EventOwnerPermissionAdded struct {
		OwnerNickName  string    `json:"nickName"`
		AccessUserUUID string    `json:"accessUserUUID"`
		User           EventUser `json:"user"`
		PermissionName string    `json:"permission"`
	}
	EventOwnerVerifiedByAdmin struct {
		OwnerNickName string `json:"nickName"`
		AdminUUID     string `json:"adminUUID"`
	}
	EventOwnerRejectedByAdmin struct {
		OwnerNickName string `json:"nickName"`
		Reason        string `json:"reason"`
		AdminUUID     string `json:"adminUUID"`
	}
	EventOwnerDeletedByAdmin struct {
		OwnerNickName string `json:"nickName"`
		Reason        string `json:"reason"`
		AdminUUID     string `json:"adminUUID"`
	}
	EventOwnerRecoverByAdmin struct {
		OwnerNickName string `json:"nickName"`
		Reason        string `json:"reason"`
		AdminUUID     string `json:"adminUUID"`
	}
	EventOwnerDisabled struct {
		UserName      string `json:"nickName"`
		UserUUID      string `json:"userUUID"`
		UserCode      string `json:"userCode"`
		OwnerNickName string `json:"ownership"`
	}
	EventOwnerEnabled struct {
		UserName      string `json:"nickName"`
		UserUUID      string `json:"userUUID"`
		UserCode      string `json:"userCode"`
		OwnerNickName string `json:"ownership"`
	}
)

type ownerEvents struct {
	publisher events.Publisher
	topics    config.Topics
}

type EventConfig struct {
	Topics    config.Topics
	Publisher events.Publisher
}

func NewEvents(config EventConfig) Events {
	return &ownerEvents{
		publisher: config.Publisher,
		topics:    config.Topics,
	}
}

func (e *ownerEvents) Created(event *EventOwnerCreated) {
	_ = e.publisher.Publish(e.topics.Owner.Created, event)
}

func (e *ownerEvents) UserRemoved(event *EventOwnerUserRemoved) {
	_ = e.publisher.Publish(e.topics.Owner.UserRemoved, event)
}

func (e *ownerEvents) UserAdded(event *EventOwnerUserAdded) {
	_ = e.publisher.Publish(e.topics.Owner.UserAdded, event)
}

func (e *ownerEvents) UserPermissionRemoved(event *EventOwnerPermissionRemoved) {
	_ = e.publisher.Publish(e.topics.Owner.UserPermissionRemoved, event)
}

func (e *ownerEvents) UserPermissionAdded(event *EventOwnerPermissionAdded) {
	_ = e.publisher.Publish(e.topics.Owner.UserPermissionAdded, event)
}

func (e *ownerEvents) VerifiedByAdmin(event *EventOwnerVerifiedByAdmin) {
	_ = e.publisher.Publish(e.topics.Owner.VerifiedByAdmin, event)
}

func (e *ownerEvents) DeletedByAdmin(event *EventOwnerDeletedByAdmin) {
	_ = e.publisher.Publish(e.topics.Owner.DeletedByAdmin, event)
}

func (e *ownerEvents) RejectedByAdmin(event *EventOwnerRejectedByAdmin) {
	_ = e.publisher.Publish(e.topics.Owner.RejectedByAdmin, event)
}

func (e *ownerEvents) RecoverByAdmin(event *EventOwnerRecoverByAdmin) {
	_ = e.publisher.Publish(e.topics.Owner.RecoverByAdmin, event)
}

func (e *ownerEvents) Disabled(event *EventOwnerDisabled) {
	_ = e.publisher.Publish(e.topics.Owner.Disabled, event)
}

func (e *ownerEvents) Enabled(event *EventOwnerEnabled) {
	_ = e.publisher.Publish(e.topics.Owner.Enabled, event)
}
