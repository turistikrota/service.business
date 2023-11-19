package entity

import (
	"time"

	"github.com/turistikrota/service.business/src/domain/business"
)

type MongoBusiness struct {
	UUID         string                    `bson:"_id,omitempty"`
	NickName     string                    `bson:"nick_name"`
	RealName     string                    `bson:"real_name"`
	BusinessType string                    `bson:"business_type"`
	Individual   *MongoBusinessIndividual  `bson:"individual"`
	Corporation  *MongoBusinessCorporation `bson:"corporation"`
	Users        []*MongoBusinessUser      `bson:"users"`
	RejectReason *string                   `bson:"reject_reason,omitempty"`
	IsEnabled    bool                      `bson:"is_enabled"`
	IsVerified   bool                      `bson:"is_verified"`
	IsDeleted    bool                      `bson:"is_deleted"`
	VerifiedAt   *time.Time                `bson:"verified_at"`
	CreatedAt    *time.Time                `bson:"created_at"`
	UpdatedAt    *time.Time                `bson:"updated_at"`
}

type MongoBusinessIndividual struct {
	FirstName      string    `bson:"first_name"`
	LastName       string    `bson:"last_name"`
	IdentityNumber []byte    `bson:"identity_number"`
	Province       string    `bson:"province"`
	District       string    `bson:"district"`
	Address        string    `bson:"address"`
	SerialNumber   []byte    `bson:"serial_number"`
	DateOfBirth    time.Time `bson:"date_of_birth"`
}

type MongoBusinessCorporation struct {
	TaxNumber []byte `bson:"tax_number"`
	Province  string `bson:"province"`
	District  string `bson:"district"`
	Address   string `bson:"address"`
	TaxOffice string `bson:"tax_office"`
	Title     string `bson:"title"`
	Type      string `bson:"type"`
}

type MongoBusinessUser struct {
	UUID   string    `bson:"uuid"`
	Name   string    `bson:"name"`
	Roles  []string  `bson:"roles"`
	JoinAt time.Time `bson:"join_at"`
}

func (m *MongoBusiness) FromBusiness(business *business.Entity) *MongoBusiness {
	m.NickName = business.NickName
	m.RealName = business.RealName
	m.BusinessType = string(business.BusinessType)
	m.Individual = &MongoBusinessIndividual{
		FirstName:      business.Individual.FirstName,
		LastName:       business.Individual.LastName,
		IdentityNumber: business.Individual.IdentityNumber,
		Province:       business.Individual.Province,
		District:       business.Individual.District,
		Address:        business.Individual.Address,
		SerialNumber:   business.Individual.SerialNumber,
		DateOfBirth:    business.Individual.DateOfBirth,
	}
	m.Corporation = &MongoBusinessCorporation{
		TaxNumber: business.Corporation.TaxNumber,
		Province:  business.Corporation.Province,
		District:  business.Corporation.District,
		Address:   business.Corporation.Address,
		Type:      string(business.Corporation.Type),
		TaxOffice: business.Corporation.TaxOffice,
		Title:     business.Corporation.Title,
	}
	m.Users = m.fromBusinessUsers(business.Users)
	m.RejectReason = business.RejectReason
	m.IsEnabled = business.IsEnabled
	m.IsVerified = business.IsVerified
	m.VerifiedAt = business.VerifiedAt
	m.CreatedAt = business.CreatedAt
	m.UpdatedAt = business.UpdatedAt
	return m
}

func (m *MongoBusiness) ToBusiness() *business.Entity {
	e := &business.Entity{
		UUID:         m.UUID,
		NickName:     m.NickName,
		RealName:     m.RealName,
		BusinessType: business.Type(m.BusinessType),
		IsEnabled:    m.IsEnabled,
		RejectReason: m.RejectReason,
		IsVerified:   m.IsVerified,
		VerifiedAt:   m.VerifiedAt,
		IsDeleted:    m.IsDeleted,
		CreatedAt:    m.CreatedAt,
		UpdatedAt:    m.UpdatedAt,
	}
	if m.Individual != nil {
		e.Individual = m.Individual.ToBusinessIndividual()
	}
	if m.Corporation != nil {
		e.Corporation = m.Corporation.ToBusinessCorporation()
	}
	if m.Users != nil {
		e.Users = m.ToBusinessUsers()
	}

	return e
}

func (m *MongoBusiness) ToBusinessWithUser(u business.UserDetail) *business.EntityWithUser {
	for _, user := range m.Users {
		if user.Name == u.Name {
			return &business.EntityWithUser{
				Entity: *m.ToBusiness(),
				User:   user.ToBusinessUser(),
			}
		}
	}
	return nil
}

func (i *MongoBusinessIndividual) ToBusinessIndividual() business.Individual {
	return business.Individual{
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

func (c *MongoBusinessCorporation) ToBusinessCorporation() business.Corporation {
	return business.Corporation{
		TaxNumber: c.TaxNumber,
		Province:  c.Province,
		District:  c.District,
		Address:   c.Address,
		TaxOffice: c.TaxOffice,
		Title:     c.Title,
		Type:      business.CorporationType(c.Type),
	}
}

func (m *MongoBusiness) ToBusinessUsers() []business.User {
	var users []business.User
	for _, user := range m.Users {
		users = append(users, user.ToBusinessUser())
	}
	return users
}

func (u *MongoBusinessUser) ToBusinessUser() business.User {
	return business.User{
		UUID:   u.UUID,
		Name:   u.Name,
		Roles:  u.Roles,
		JoinAt: u.JoinAt,
	}
}

func (u *MongoBusinessUser) FromBusinessUser(user *business.User) *MongoBusinessUser {
	u.UUID = user.UUID
	u.Name = user.Name
	u.Roles = user.Roles
	u.JoinAt = user.JoinAt
	return u
}

func (m *MongoBusiness) fromBusinessUsers(users []business.User) []*MongoBusinessUser {
	var mongoUsers []*MongoBusinessUser
	for _, user := range users {
		mongoUsers = append(mongoUsers, &MongoBusinessUser{
			UUID:   user.UUID,
			Name:   user.Name,
			Roles:  user.Roles,
			JoinAt: user.JoinAt,
		})
	}
	return mongoUsers
}
