package mongo

import (
	mongo_owner "github.com/turistikrota/service.owner/src/adapters/mongo/owner"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo interface {
	NewOwner(factory owner.Factory, collection *mongo.Collection) owner.Repository
}

type mongodb struct{}

func New() Mongo {
	return &mongodb{}
}

func (m *mongodb) NewOwner(factory owner.Factory, collection *mongo.Collection) owner.Repository {
	return mongo_owner.New(factory, collection)
}
