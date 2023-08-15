package entity

import (
	"time"

	"github.com/turistikrota/service.owner/src/domain/account"
)

type MongoAccount struct {
	UserUUID   string     `bson:"user_uuid"`
	UserName   string     `bson:"user_name"`
	FullName   string     `bson:"full_name"`
	IsActive   bool       `bson:"is_active"`
	IsDeleted  bool       `bson:"is_deleted"`
	IsVerified bool       `bson:"is_verified"`
	BirthDate  *time.Time `bson:"birth_date"`
	CreatedAt  *time.Time `bson:"created_at"`
}

func (e *MongoAccount) ToEntity() *account.Entity {
	return &account.Entity{
		UserUUID:   e.UserUUID,
		UserName:   e.UserName,
		FullName:   e.FullName,
		IsActive:   e.IsActive,
		IsDeleted:  e.IsDeleted,
		IsVerified: e.IsVerified,
		BirthDate:  e.BirthDate,
		CreatedAt:  e.CreatedAt,
	}
}

func (e *MongoAccount) FromEntity(entity *account.Entity) *MongoAccount {
	e.UserUUID = entity.UserUUID
	e.UserName = entity.UserName
	e.FullName = entity.FullName
	e.IsActive = entity.IsActive
	e.IsDeleted = entity.IsDeleted
	e.IsVerified = entity.IsVerified
	e.BirthDate = entity.BirthDate
	e.CreatedAt = entity.CreatedAt
	return e
}
