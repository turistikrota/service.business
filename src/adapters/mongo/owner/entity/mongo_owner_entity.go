package entity

import (
	"time"

	"github.com/turistikrota/service.owner/src/domain/owner"
)

type MongoOwner struct {
	UUID        string                 `bson:"_id,omitempty"`
	NickName    string                 `bson:"nick_name"`
	RealName    string                 `bson:"real_name"`
	AvatarURL   string                 `bson:"avatar_url"`
	CoverURL    string                 `bson:"cover_url"`
	OwnerType   string                 `bson:"owner_type"`
	Individual  *MongoOwnerIndividual  `bson:"individual"`
	Corporation *MongoOwnerCorporation `bson:"corporation"`
	Users       []*MongoOwnerUser      `bson:"users"`
	IsEnabled   bool                   `bson:"is_enabled"`
	IsVerified  bool                   `bson:"is_verified"`
	IsDeleted   bool                   `bson:"is_deleted"`
	VerifiedAt  *time.Time             `bson:"verified_at"`
	CreatedAt   *time.Time             `bson:"created_at"`
	UpdatedAt   *time.Time             `bson:"updated_at"`
}

type MongoOwnerIndividual struct {
	FirstName      string    `bson:"first_name"`
	LastName       string    `bson:"last_name"`
	IdentityNumber []byte    `bson:"identity_number"`
	Province       string    `bson:"province"`
	District       string    `bson:"district"`
	Address        string    `bson:"address"`
	SerialNumber   []byte    `bson:"serial_number"`
	DateOfBirth    time.Time `bson:"date_of_birth"`
}

type MongoOwnerCorporation struct {
	TaxNumber []byte `bson:"tax_number"`
	Province  string `bson:"province"`
	District  string `bson:"district"`
	Address   string `bson:"address"`
	TaxOffice string `bson:"tax_office"`
	Title     string `bson:"title"`
	Type      string `bson:"type"`
}

type MongoOwnerUser struct {
	UUID   string    `bson:"uuid"`
	Name   string    `bson:"name"`
	Code   string    `bson:"code"`
	Roles  []string  `bson:"roles"`
	JoinAt time.Time `bson:"join_at"`
}

func (m *MongoOwner) FromOwner(owner *owner.Entity) *MongoOwner {
	m.NickName = owner.NickName
	m.RealName = owner.RealName
	m.AvatarURL = owner.AvatarURL
	m.CoverURL = owner.CoverURL
	m.OwnerType = string(owner.OwnerType)
	m.Individual = &MongoOwnerIndividual{
		FirstName:      owner.Individual.FirstName,
		LastName:       owner.Individual.LastName,
		IdentityNumber: owner.Individual.IdentityNumber,
		Province:       owner.Individual.Province,
		District:       owner.Individual.District,
		Address:        owner.Individual.Address,
		SerialNumber:   owner.Individual.SerialNumber,
		DateOfBirth:    owner.Individual.DateOfBirth,
	}
	m.Corporation = &MongoOwnerCorporation{
		TaxNumber: owner.Corporation.TaxNumber,
		Province:  owner.Corporation.Province,
		District:  owner.Corporation.District,
		Address:   owner.Corporation.Address,
		Type:      string(owner.Corporation.Type),
		TaxOffice: owner.Corporation.TaxOffice,
		Title:     owner.Corporation.Title,
	}
	m.Users = m.fromOwnerUsers(owner.Users)
	m.IsEnabled = owner.IsEnabled
	m.IsVerified = owner.IsVerified
	m.VerifiedAt = owner.VerifiedAt
	m.CreatedAt = owner.CreatedAt
	m.UpdatedAt = owner.UpdatedAt
	return m
}

func (m *MongoOwner) ToOwner() *owner.Entity {
	e := &owner.Entity{
		UUID:       m.UUID,
		NickName:   m.NickName,
		RealName:   m.RealName,
		AvatarURL:  m.AvatarURL,
		CoverURL:   m.CoverURL,
		OwnerType:  owner.Type(m.OwnerType),
		IsEnabled:  m.IsEnabled,
		IsVerified: m.IsVerified,
		VerifiedAt: m.VerifiedAt,
		CreatedAt:  m.CreatedAt,
		UpdatedAt:  m.UpdatedAt,
	}
	if m.Individual != nil {
		e.Individual = m.Individual.ToOwnerIndividual()
	}
	if m.Corporation != nil {
		e.Corporation = m.Corporation.ToOwnerCorporation()
	}
	if m.Users != nil {
		e.Users = m.ToOwnerUsers()
	}

	return e
}

func (m *MongoOwner) ToOwnerWithUser(u owner.UserDetail) *owner.EntityWithUser {
	for _, user := range m.Users {
		if user.Name == u.Name && user.Code == u.Code {
			return &owner.EntityWithUser{
				Entity: *m.ToOwner(),
				User:   user.ToOwnerUser(),
			}
		}
	}
	return nil
}

func (i *MongoOwnerIndividual) ToOwnerIndividual() owner.Individual {
	return owner.Individual{
		FirstName:      i.FirstName,
		LastName:       i.LastName,
		IdentityNumber: i.IdentityNumber,
		Province:       i.Province,
		District:       i.District,
		Address:        i.Address,
		SerialNumber:   i.SerialNumber,
		DateOfBirth:    i.DateOfBirth,
	}
}

func (c *MongoOwnerCorporation) ToOwnerCorporation() owner.Corporation {
	return owner.Corporation{
		TaxNumber: c.TaxNumber,
		Province:  c.Province,
		District:  c.District,
		Address:   c.Address,
		TaxOffice: c.TaxOffice,
		Title:     c.Title,
		Type:      owner.CorporationType(c.Type),
	}
}

func (m *MongoOwner) ToOwnerUsers() []owner.User {
	var users []owner.User
	for _, user := range m.Users {
		users = append(users, user.ToOwnerUser())
	}
	return users
}

func (u *MongoOwnerUser) ToOwnerUser() owner.User {
	return owner.User{
		UUID:   u.UUID,
		Name:   u.Name,
		Code:   u.Code,
		Roles:  u.Roles,
		JoinAt: u.JoinAt,
	}
}

func (u *MongoOwnerUser) FromOwnerUser(user *owner.User) *MongoOwnerUser {
	u.UUID = user.UUID
	u.Name = user.Name
	u.Code = user.Code
	u.Roles = user.Roles
	u.JoinAt = user.JoinAt
	return u
}

func (m *MongoOwner) fromOwnerUsers(users []owner.User) []*MongoOwnerUser {
	var mongoUsers []*MongoOwnerUser
	for _, user := range users {
		mongoUsers = append(mongoUsers, &MongoOwnerUser{
			UUID:   user.UUID,
			Name:   user.Name,
			Code:   user.Code,
			Roles:  user.Roles,
			JoinAt: user.JoinAt,
		})
	}
	return mongoUsers
}
