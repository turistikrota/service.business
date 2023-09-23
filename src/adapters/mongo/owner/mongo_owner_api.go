package owner

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/turistikrota/service.owner/src/adapters/mongo/owner/entity"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Create(ctx context.Context, owner *owner.Entity) (*owner.Entity, *i18np.Error) {
	n := &entity.MongoOwner{}
	res, err := r.collection.InsertOne(ctx, n.FromOwner(owner))
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	owner.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return owner, nil
}

func (r *repo) GetByNickName(ctx context.Context, nickName string) (*owner.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return o.ToOwner(), nil
}

func (r *repo) GetByIndividual(ctx context.Context, individual owner.Individual) (*owner.Entity, bool, *i18np.Error) {
	filter := bson.M{
		entity.IndividualField(entity.IndividualFields.FirstName):   individual.FirstName,
		entity.IndividualField(entity.IndividualFields.LastName):    individual.LastName,
		entity.IndividualField(entity.IndividualFields.DateOfBirth): individual.DateOfBirth,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, false, err
	}
	if !exist {
		return nil, true, nil
	}
	return o.ToOwner(), false, nil
}

func (r *repo) GetByCorporation(ctx context.Context, corporation owner.Corporation) (*owner.Entity, bool, *i18np.Error) {
	filter := bson.M{
		entity.CorporationField(entity.CorporationFields.TaxOffice): corporation.TaxOffice,
		entity.CorporationField(entity.CorporationFields.Title):     corporation.Title,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, false, err
	}
	if !exist {
		return nil, true, nil
	}
	return o.ToOwner(), false, nil
}

func (r *repo) CheckNickName(ctx context.Context, nickName string) (bool, *i18np.Error) {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return true, err
	}
	return o != nil && exist, nil
}

func (r *repo) GetWithUser(ctx context.Context, nickName string, user owner.UserDetail) (*owner.EntityWithUser, *i18np.Error) {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
		entity.UserField(entity.UserFields.Code): user.Code,
		entity.UserField(entity.UserFields.UUID): user.UUID,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return o.ToOwnerWithUser(user), nil
}

func (r *repo) ProfileView(ctx context.Context, nickName string) (*owner.Entity, *i18np.Error) {
	filter := bson.M{
		entity.Fields.NickName:  nickName,
		entity.Fields.IsEnabled: true,
		entity.Fields.IsDeleted: false,
	}
	opts := options.FindOne().SetProjection(bson.M{
		entity.Fields.UUID:        0,
		entity.Fields.Users:       0,
		entity.Fields.IsEnabled:   0,
		entity.Fields.Corporation: 0,
		entity.Fields.Individual:  0,
		entity.Fields.ActivatedAt: 0,
		entity.Fields.DisabledAt:  0,
		entity.Fields.VerifiedAt:  0,
		entity.Fields.UpdatedAt:   0,
	})
	o, exist, err := r.helper.GetFilter(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return o.ToOwner(), nil
}

func (r *repo) ListByUserUUID(ctx context.Context, user owner.UserDetail) ([]*owner.Entity, *i18np.Error) {
	filter := bson.M{
		entity.UserField(entity.UserFields.Name): user.Name,
		entity.UserField(entity.UserFields.Code): user.Code,
		entity.UserField(entity.UserFields.UUID): user.UUID,
	}
	opts := options.Find().SetProjection(bson.M{
		entity.Fields.UUID:        0,
		entity.Fields.Users:       0,
		entity.Fields.Corporation: 0,
		entity.Fields.Individual:  0,
		entity.Fields.ActivatedAt: 0,
		entity.Fields.DisabledAt:  0,
		entity.Fields.VerifiedAt:  0,
	})
	return r.helper.GetListFilterTransform(ctx, filter, func(o *entity.MongoOwner) *owner.Entity {
		return o.ToOwner()
	}, opts)
}

func (r *repo) AddUser(ctx context.Context, nickName string, user *owner.User) *i18np.Error {
	// filter with nickName and users array not contains user
	filter := bson.M{
		entity.Fields.NickName: nickName,
		"$or": []bson.M{
			{entity.UserField(entity.UserFields.Name): bson.M{"$ne": user.Name}},
			{entity.UserField(entity.UserFields.Code): bson.M{"$ne": user.Code}},
		},
	}
	setter := bson.M{
		"$addToSet": bson.M{
			entity.Fields.Users: bson.M{
				entity.UserFields.UUID:   user.UUID,
				entity.UserFields.Name:   user.Name,
				entity.UserFields.Code:   user.Code,
				entity.UserFields.Roles:  user.Roles,
				entity.UserFields.JoinAt: user.JoinAt,
			},
		},
		"$set": bson.M{
			entity.Fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) RemoveUser(ctx context.Context, nickName string, user owner.UserDetail) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
		entity.UserField(entity.UserFields.Code): user.Code,
	}
	t := time.Now()
	setter := bson.M{
		"$pull": bson.M{
			entity.Fields.Users: bson.M{
				entity.UserFields.Name: user.Name,
				entity.UserFields.Code: user.Code,
			},
		},
		"$set": bson.M{
			entity.Fields.UpdatedAt: t,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) RemoveUserPermission(ctx context.Context, nickName string, user owner.UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
		entity.UserField(entity.UserFields.Code): user.Code,
	}
	t := time.Now()
	setter := bson.M{
		"$pull": bson.M{
			entity.UserArrayFieldInArray(entity.UserFields.Roles): permission,
		},
		"$set": bson.M{
			entity.Fields.UpdatedAt: t,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) AddUserPermission(ctx context.Context, nickName string, user owner.UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
		entity.UserField(entity.UserFields.Code): user.Code,
		entity.UserField(entity.UserFields.Roles): bson.M{
			"$ne": permission,
		},
	}
	t := time.Now()
	setter := bson.M{
		"$push": bson.M{
			entity.UserArrayFieldInArray(entity.UserFields.Roles): permission,
		},
		"$set": bson.M{
			entity.Fields.UpdatedAt: t,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Enable(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsEnabled:  true,
			entity.Fields.IsVerified: false,
			entity.Fields.VerifiedAt: nil,
			entity.Fields.UpdatedAt:  time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Verify(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	t := time.Now()
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsVerified: true,
			entity.Fields.UpdatedAt:  t,
			entity.Fields.VerifiedAt: t,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Disable(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsEnabled:  false,
			entity.Fields.IsVerified: false,
			entity.Fields.VerifiedAt: nil,
			entity.Fields.UpdatedAt:  time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsDeleted:  true,
			entity.Fields.IsVerified: false,
			entity.Fields.VerifiedAt: nil,
			entity.Fields.UpdatedAt:  time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Recover(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.IsDeleted:  false,
			entity.Fields.IsVerified: false,
			entity.Fields.VerifiedAt: nil,
			entity.Fields.UpdatedAt:  time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) ListOwnershipUsers(ctx context.Context, nickName string, user owner.UserDetail) ([]owner.User, *i18np.Error) {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
		entity.UserField(entity.UserFields.Code): user.Code,
	}
	opts := options.FindOne().SetProjection(bson.M{
		entity.Fields.Users: 1,
	})
	o, exist, err := r.helper.GetFilter(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return o.ToOwnerUsers(), nil
}
