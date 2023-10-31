package entity

import (
	"time"

	"github.com/turistikrota/service.owner/src/domain/invite"
)

type MongoInvite struct {
	UUID      string     `bson:"_id,omitempty"`
	OwnerUUID string     `bson:"owner_uuid"`
	Email     string     `bson:"email"`
	IsUsed    bool       `bson:"is_used"`
	IsDeleted bool       `bson:"is_deleted"`
	CreatedAt *time.Time `bson:"created_at"`
	UpdatedAt *time.Time `bson:"updated_at"`
}

func (m *MongoInvite) FromInvite(invite *invite.Entity) *MongoInvite {
	m.UUID = invite.UUID
	m.OwnerUUID = invite.OwnerUUID
	m.Email = invite.Email
	m.IsUsed = invite.IsUsed
	m.IsDeleted = invite.IsDeleted
	m.CreatedAt = invite.CreatedAt
	m.UpdatedAt = invite.UpdatedAt
	return m
}

func (m *MongoInvite) ToInvite() *invite.Entity {
	return &invite.Entity{
		UUID:      m.UUID,
		OwnerUUID: m.OwnerUUID,
		Email:     m.Email,
		IsUsed:    m.IsUsed,
		IsDeleted: m.IsDeleted,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
	}
}
