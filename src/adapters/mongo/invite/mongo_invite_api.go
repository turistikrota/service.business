package invite

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.business/src/adapters/mongo/invite/entity"
	"github.com/turistikrota/service.business/src/domain/invite"
	"github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *repo) Create(ctx context.Context, e *invite.Entity) (*invite.Entity, *i18np.Error) {
	i := &entity.MongoInvite{}
	res, err := r.collection.InsertOne(ctx, i.FromInvite(e))
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	e.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return e, nil
}

func (r *repo) GetByUUID(ctx context.Context, uuid string) (*invite.Entity, *i18np.Error) {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return nil, r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		entity.Fields.UUID: id,
	}
	o, exist, error := r.helper.GetFilter(ctx, filter)
	if error != nil {
		return nil, r.factory.Errors.Failed("get by uuid")
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return o.ToInvite(), nil
}

func (r *repo) GetByBusinessUUID(ctx context.Context, businessUUID string) ([]*invite.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.BusinessUUID: businessUUID,
	}
	return r.helper.GetListFilterTransform(ctx, filter, func(o *entity.MongoInvite) *invite.Entity {
		return o.ToInvite()
	})
}

func (r *repo) GetByEmail(ctx context.Context, email string) ([]*invite.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.Email: email,
	}
	return r.helper.GetListFilterTransform(ctx, filter, func(o *entity.MongoInvite) *invite.Entity {
		return o.ToInvite()
	})
}

func (r *repo) Use(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		entity.Fields.UUID: id,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsUsed:    true,
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, uuid string) *i18np.Error {
	id, err := mongo.TransformId(uuid)
	if err != nil {
		return r.factory.Errors.InvalidUUID()
	}
	filter := bson.M{
		entity.Fields.UUID: id,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsDeleted: true,
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}
