package mongo

import (
	mongo_business "github.com/turistikrota/service.business/src/adapters/mongo/business"
	mongo_invite "github.com/turistikrota/service.business/src/adapters/mongo/invite"
	"github.com/turistikrota/service.business/src/domain/business"
	"github.com/turistikrota/service.business/src/domain/invite"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewBusiness(factory business.Factory, collection *mongo.Collection) business.Repository
	NewInvite(factory invite.Factory, collection *mongo.Collection) invite.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewBusiness(factory business.Factory, collection *mongo.Collection) business.Repository {
	return mongo_business.New(factory, collection)
}

func (m *mongodb) NewInvite(factory invite.Factory, collection *mongo.Collection) invite.Repository {
	return mongo_invite.New(factory, collection)
}
