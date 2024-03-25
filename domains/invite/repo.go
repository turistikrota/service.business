package invite

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error)
	GetByBusinessUUID(ctx context.Context, businessUUID string) ([]*Entity, *i18np.Error)
	GetByEmail(ctx context.Context, email string) ([]*Entity, *i18np.Error)
	Use(ctx context.Context, uuid string) *i18np.Error
	Delete(ctx context.Context, uuid string) *i18np.Error
}

type repo struct {
	factory    Factory
	collection *mongo.Collection
	helper     mongo2.Helper[*Entity, *Entity]
}

func NewRepository(factory Factory, collection *mongo.Collection) Repository {
	return &repo{
		factory:    factory,
		collection: collection,
		helper:     mongo2.NewHelper[*Entity, *Entity](collection, createEntity),
	}
}

func createEntity() **Entity {
	return new(*Entity)
}

func (r *repo) Create(ctx context.Context, e *Entity) (*Entity, *i18np.Error) {
	res, err := r.collection.InsertOne(ctx, e)
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	e.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return e, nil
}

func (r *repo) GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error) {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	o, exist, error := r.helper.GetFilter(ctx, filter)
	if error != nil {
		return nil, r.factory.Errors.Failed("get by uuid")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *o, nil
}

func (r *repo) GetByBusinessUUID(ctx context.Context, businessUUID string) ([]*Entity, *i18np.Error) {
	filter := bson.M{
		fields.BusinessUUID: businessUUID,
	}
	return r.helper.GetListFilter(ctx, filter)
}

func (r *repo) GetByEmail(ctx context.Context, email string) ([]*Entity, *i18np.Error) {
	filter := bson.M{
		fields.Email: email,
	}
	return r.helper.GetListFilter(ctx, filter)
}

func (r *repo) Use(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsUsed:    true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo2.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		fields.UUID: id,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsDeleted: true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}
