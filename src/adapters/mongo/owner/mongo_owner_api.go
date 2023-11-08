package owner

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.owner/src/adapters/mongo/owner/entity"
	"github.com/turistikrota/service.owner/src/domain/owner"
	"github.com/turistikrota/service.shared/db/mongo"
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

func (r *repo) AddUser(ctx context.Context, ownerUUID string, user *owner.User) *i18np.Error {
	id, err := mongo.TransformId(ownerUUID)
	if err != nil {
		return r.factory.Errors.Failed("add user" + ownerUUID)
	}
	filter := bson.M{
		entity.Fields.UUID: id,
		"$or": []bson.M{
			{entity.UserField(entity.UserFields.Name): bson.M{"$ne": user.Name}},
			{entity.UserField(entity.UserFields.UUID): bson.M{"$ne": user.UUID}},
		},
	}
	setter := bson.M{
		"$addToSet": bson.M{
			entity.Fields.Users: bson.M{
				entity.UserFields.UUID:   user.UUID,
				entity.UserFields.Name:   user.Name,
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
	}
	t := time.Now()
	setter := bson.M{
		"$pull": bson.M{
			entity.Fields.Users: bson.M{
				entity.UserFields.Name: user.Name,
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
			entity.Fields.RejectReason: nil,
			entity.Fields.IsVerified:   true,
			entity.Fields.UpdatedAt:    t,
			entity.Fields.VerifiedAt:   t,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Reject(ctx context.Context, nickName string, reason string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName: nickName,
	}
	t := time.Now()
	setter := bson.M{
		"$set": bson.M{
			entity.Fields.RejectReason: reason,
			entity.Fields.UpdatedAt:    t,
			entity.Fields.IsVerified:   false,
			entity.Fields.VerifiedAt:   nil,
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

func (r *repo) AdminListAll(ctx context.Context, listConfig list.Config) (*list.Result[*owner.AdminListDto], *i18np.Error) {
	filter := bson.M{}
	l, err := r.helper.GetListFilterTransform(ctx, filter, func(o *entity.MongoOwner) *owner.Entity {
		return o.ToOwner()
	})
	if err != nil {
		return nil, err
	}
	filtered, _err := r.helper.GetFilterCount(ctx, bson.M{})
	if _err != nil {
		return nil, _err
	}
	total, err := r.helper.GetFilterCount(ctx, filter)
	if err != nil {
		return nil, err
	}
	li := make([]*owner.AdminListDto, 0)
	for _, o := range l {
		dto := &owner.AdminListDto{
			UUID:       o.UUID,
			NickName:   o.NickName,
			RealName:   o.RealName,
			OwnerType:  string(o.OwnerType),
			IsEnabled:  o.IsEnabled,
			IsVerified: o.IsVerified,
			IsDeleted:  o.IsDeleted,
			CreatedAt:  o.CreatedAt.String(),
			UpdatedAt:  o.UpdatedAt.String(),
		}
		if o.VerifiedAt != nil {
			dto.VerifiedAt = o.VerifiedAt.String()
		}
		li = append(li, dto)
	}
	return &list.Result[*owner.AdminListDto]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          li,
	}, nil
}
