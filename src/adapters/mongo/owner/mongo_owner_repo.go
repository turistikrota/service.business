package owner

import (
	"github.com/turistikrota/service.owner/src/adapters/mongo/owner/entity"
	"github.com/turistikrota/service.owner/src/domain/owner"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	factory    owner.Factory
	collection *mongo.Collection
	helper     mongo2.Helper[entity.MongoOwner, *owner.Entity]
}

func New(factory owner.Factory, collection *mongo.Collection) owner.Repository {
	validate(factory, collection)
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[entity.MongoOwner, *owner.Entity](collection, createEntity),
	}
}

func createEntity() *entity.MongoOwner {
	return &entity.MongoOwner{}
}

func validate(factory owner.Factory, collection *mongo.Collection) {
	if factory.IsZero() {
		panic("owner Factory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}
