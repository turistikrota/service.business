package mongo

import (
	mongo_account "github.com/turistikrota/service.owner/src/adapters/mongo/account"
	mongo_owner "github.com/turistikrota/service.owner/src/adapters/mongo/owner"
	"github.com/turistikrota/service.owner/src/domain/account"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewOwner(factory owner.Factory, collection *mongo.Collection) owner.Repository
	NewAccount(factory account.Factory, collection *mongo.Collection) account.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewOwner(factory owner.Factory, collection *mongo.Collection) owner.Repository {
	return mongo_owner.New(factory, collection)
}

func (m *mongodb) NewAccount(factory account.Factory, collection *mongo.Collection) account.Repository {
	return mongo_account.New(factory, collection)
}
