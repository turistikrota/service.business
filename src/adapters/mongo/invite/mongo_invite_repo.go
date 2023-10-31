package invite

import (
	"github.com/turistikrota/service.owner/src/adapters/mongo/invite/entity"
	"github.com/turistikrota/service.owner/src/domain/invite"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

type repo struct {
	factory    invite.Factory
	collection *mongo.Collection
	helper     mongo2.Helper[entity.MongoInvite, *invite.Entity]
}

func New(factory invite.Factory, collection *mongo.Collection) invite.Repository {
	validate(factory, collection)
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[entity.MongoInvite, *invite.Entity](collection, createEntity),
	}
}

func createEntity() *entity.MongoInvite {
	return &entity.MongoInvite{}
}

func validate(factory invite.Factory, collection *mongo.Collection) {
	if factory.IsZero() {
		panic("invite Factory is zero")
	}
	if collection == nil {
		panic("collection is nil")
	}
}
