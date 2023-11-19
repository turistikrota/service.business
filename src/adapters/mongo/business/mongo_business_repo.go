package business

import (
	"github.com/turistikrota/service.business/src/adapters/mongo/business/entity"
	"github.com/turistikrota/service.business/src/domain/business"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	factory    business.Factory
	collection *mongo.Collection
	helper     mongo2.Helper[entity.MongoBusiness, *business.Entity]
}

func New(factory business.Factory, collection *mongo.Collection) business.Repository {
	validate(factory, collection)
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[entity.MongoBusiness, *business.Entity](collection, createEntity),
	}
}

func createEntity() *entity.MongoBusiness {
	return &entity.MongoBusiness{}
}

func validate(factory business.Factory, collection *mongo.Collection) {
	if factory.IsZero() {
		panic("business Factory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}
