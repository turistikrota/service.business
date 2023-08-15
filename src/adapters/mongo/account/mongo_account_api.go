package account

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.owner/src/adapters/mongo/account/entity"
	"github.com/turistikrota/service.owner/src/domain/account"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) updateOne(ctx context.Context, filter bson.M, setter bson.M, opts ...*options.UpdateOptions) *i18np.Error {
	res, err := r.collection.UpdateOne(ctx, filter, setter, opts...)
	if err != nil {
		return r.factory.Errors.Failed("update")
	}
	if res.MatchedCount == 0 {
		return r.factory.Errors.NotFound()
	}
	return nil
}

func (r *repo) Create(ctx context.Context, account *account.Entity) *i18np.Error {
	n := &entity.MongoAccount{}
	_, err := r.collection.InsertOne(ctx, n.FromEntity(account))
	if err != nil {
		return r.factory.Errors.Failed("create")
	}
	return nil
}

func (r *repo) Update(ctx context.Context, u account.UserUnique, account *account.Entity) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UserUUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.UserCode: u.Code,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.UserName:  account.UserName,
			entity.Fields.UserCode:  account.UserCode,
			entity.Fields.FullName:  account.FullName,
			entity.Fields.Avatar:    account.AvatarURL,
			entity.Fields.BirthDate: account.BirthDate,
		},
	}
	return r.updateOne(ctx, filter, setter)
}

func (r *repo) Disable(ctx context.Context, u account.UserUnique) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UserUUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.UserCode: u.Code,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsActive: false,
		},
	}
	return r.updateOne(ctx, filter, setter)
}

func (r *repo) Enable(ctx context.Context, u account.UserUnique) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UserUUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.UserCode: u.Code,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsActive: true,
		},
	}
	return r.updateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, u account.UserUnique) *i18np.Error {
	filter := bson.M{
		entity.Fields.UserUUID: u.UserUUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.UserCode: u.Code,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsDeleted: true,
		},
	}
	return r.updateOne(ctx, filter, setter)
}

func (r *repo) Get(ctx context.Context, u account.UserUnique) (*account.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.UserName: u.Name,
		entity.Fields.UserCode: u.Code,
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
		entity.Fields.IsActive: true,
	}
	return r.handleErrAndReturn(r.helper.GetFilter(ctx, filter))
}

func (r *repo) GetByUserUUID(ctx context.Context, u account.UserUnique) (*account.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.UserUUID: u.UserUUID,
		entity.Fields.UserName: u.Name,
		entity.Fields.UserCode: u.Code,
		entity.Fields.IsDeleted: bson.M{
			"$ne": true,
		},
		entity.Fields.IsActive: true,
	}
	return r.handleErrAndReturn(r.helper.GetFilter(ctx, filter))
}

func (r *repo) handleErrAndReturn(res *entity.MongoAccount, exist bool, err *i18np.Error) (*account.Entity, *i18np.Error) {
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return res.ToEntity(), nil
}
