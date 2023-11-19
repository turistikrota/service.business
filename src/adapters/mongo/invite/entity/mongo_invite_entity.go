package entity

import (
	"time"

	"github.com/turistikrota/service.business/src/domain/invite"
)

type MongoInvite struct {
	UUID            string     `bson:"_id,omitempty"`
	BusinessUUID    string     `bson:"business_uuid"`
	BusinessName    string     `bson:"business_name"`
	CreatorUserName string     `bson:"creator_user_name"`
	Email           string     `bson:"email"`
	IsUsed          bool       `bson:"is_used"`
	IsDeleted       bool       `bson:"is_deleted"`
	CreatedAt       *time.Time `bson:"created_at"`
	UpdatedAt       *time.Time `bson:"updated_at"`
}

func (m *MongoInvite) FromInvite(invite *invite.Entity) *MongoInvite {
	m.UUID = invite.UUID
	m.BusinessUUID = invite.BusinessUUID
	m.BusinessName = invite.BusinessNickName
	m.CreatorUserName = invite.CreatorUserName
	m.Email = invite.Email
	m.IsUsed = invite.IsUsed
	m.IsDeleted = invite.IsDeleted
	m.CreatedAt = invite.CreatedAt
	m.UpdatedAt = invite.UpdatedAt
	return m
}

func (m *MongoInvite) ToInvite() *invite.Entity {
	return &invite.Entity{
		UUID:             m.UUID,
		BusinessNickName: m.BusinessName,
		BusinessUUID:     m.BusinessUUID,
		Email:            m.Email,
		IsUsed:           m.IsUsed,
		CreatorUserName:  m.CreatorUserName,
		IsDeleted:        m.IsDeleted,
		CreatedAt:        m.CreatedAt,
		UpdatedAt:        m.UpdatedAt,
	}
}
