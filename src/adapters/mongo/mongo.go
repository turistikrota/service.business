package mongo

import (
	mongo_invite "github.com/turistikrota/service.owner/src/adapters/mongo/invite"
	mongo_owner "github.com/turistikrota/service.owner/src/adapters/mongo/owner"
	"github.com/turistikrota/service.owner/src/domain/invite"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewOwner(factory owner.Factory, collection *mongo.Collection) owner.Repository
	NewInvite(factory invite.Factory, collection *mongo.Collection) invite.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewOwner(factory owner.Factory, collection *mongo.Collection) owner.Repository {
	return mongo_owner.New(factory, collection)
}

func (m *mongodb) NewInvite(factory invite.Factory, collection *mongo.Collection) invite.Repository {
	return mongo_invite.New(factory, collection)
}
