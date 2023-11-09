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
		Owner    *Entity `json:"owner"`
		UserUUID string  `json:"userUUID"`
		UserName string  `json:"userName"`
	}
	EventUser struct {
		UUID string `json:"uuid"`
		Name string `json:"name"`
	}
	EventOwnerUserAdded struct {
		OwnerUUID string `json:"ownerUUID"`
		User      *User  `json:"user"`
	}
	EventOwnerUserRemoved struct {
		OwnerUUID      string    `json:"ownerUUID"`
		AccessUserUUID string    `json:"accessUserUUID"`
		User           EventUser `json:"user"`
	}
	EventOwnerPermissionRemoved struct {
		OwnerUUID      string    `json:"ownerUUID"`
		AccessUserUUID string    `json:"accessUserUUID"`
		AccessUserName string    `json:"accessUserName"`
		User           EventUser `json:"user"`
		PermissionName string    `json:"permission"`
	}
	EventOwnerPermissionAdded struct {
		OwnerUUID      string    `json:"ownerUUID"`
		AccessUserUUID string    `json:"accessUserUUID"`
		AccessUserName string    `json:"accessUserName"`
		User           EventUser `json:"user"`
		PermissionName string    `json:"permission"`
	}
	EventOwnerVerifiedByAdmin struct {
		OwnerUUID string `json:"ownerUUID"`
		AdminUUID string `json:"adminUUID"`
	}
	EventOwnerRejectedByAdmin struct {
		OwnerUUID string `json:"ownerUUID"`
		Reason    string `json:"reason"`
		AdminUUID string `json:"adminUUID"`
	}
	EventOwnerDeletedByAdmin struct {
		OwnerUUID string `json:"ownerUUID"`
		AdminUUID string `json:"adminUUID"`
	}
	EventOwnerRecoverByAdmin struct {
		OwnerUUID string `json:"ownerUUID"`
		AdminUUID string `json:"adminUUID"`
	}
	EventOwnerDisabled struct {
		UserName  string `json:"nickName"`
		UserUUID  string `json:"userUUID"`
		UserCode  string `json:"userCode"`
		OwnerUUID string `json:"ownerUUID"`
	}
	EventOwnerEnabled struct {
		UserName  string `json:"nickName"`
		UserUUID  string `json:"userUUID"`
		UserCode  string `json:"userCode"`
		OwnerUUID string `json:"ownerUUID"`
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
