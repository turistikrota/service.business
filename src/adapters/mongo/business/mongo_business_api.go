package business

import (
	"context"
	"time"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/types/list"
	"github.com/turistikrota/service.business/src/adapters/mongo/business/entity"
	"github.com/turistikrota/service.business/src/domain/business"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *repo) Create(ctx context.Context, business *business.Entity) (*business.Entity, *i18np.Error) {
	n := &entity.MongoBusiness{}
	res, err := r.collection.InsertOne(ctx, n.FromBusiness(business))
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	business.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return business, nil
}

func (r *repo) GetByNickName(ctx context.Context, nickName string) (*business.Entity, *i18np.Error) {
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
	return o.ToBusiness(), nil
}

func (r *repo) GetByIndividual(ctx context.Context, individual business.Individual) (*business.Entity, bool, *i18np.Error) {
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
	return o.ToBusiness(), false, nil
}

func (r *repo) GetByCorporation(ctx context.Context, corporation business.Corporation) (*business.Entity, bool, *i18np.Error) {
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
	return o.ToBusiness(), false, nil
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

func (r *repo) GetWithUser(ctx context.Context, nickName string, user business.UserDetail) (*business.EntityWithUser, *i18np.Error) {
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
	return o.ToBusinessWithUser(user), nil
}

func (r *repo) GetWithUserName(ctx context.Context, nickName string, userName string) (*business.EntityWithUser, *i18np.Error) {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): userName,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return o.ToBusinessWithUser(business.UserDetail{
		Name: userName,
	}), nil
}

func (r *repo) ProfileView(ctx context.Context, nickName string) (*business.Entity, *i18np.Error) {
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
	return o.ToBusiness(), nil
}

func (r *repo) ListByUserUUID(ctx context.Context, user business.UserDetail) ([]*business.Entity, *i18np.Error) {
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
	return r.helper.GetListFilterTransform(ctx, filter, func(o *entity.MongoBusiness) *business.Entity {
		return o.ToBusiness()
	}, opts)
}

func (r *repo) AddUser(ctx context.Context, businessName string, user *business.User) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   businessName,
		entity.UserField(entity.UserFields.Name): bson.M{"$ne": user.Name},
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

func (r *repo) RemoveUser(ctx context.Context, nickName string, user business.UserDetail) *i18np.Error {
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

func (r *repo) RemoveUserPermission(ctx context.Context, nickName string, user business.UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                   nickName,
		entity.UserField(entity.UserFields.Name): user.Name,
	}
	t := time.Now()
	setter := bson.M{
		"$pull": bson.M{
			entity.UserFieldInArray(entity.UserFields.Roles): permission,
		},
		"$set": bson.M{
			entity.Fields.UpdatedAt: t,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) AddUserPermission(ctx context.Context, nickName string, user business.UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		entity.Fields.NickName:                           nickName,
		entity.UserField(entity.UserFields.Name):         user.Name,
		entity.UserFieldInArray(entity.UserFields.Roles): bson.M{"$ne": permission},
	}
	t := time.Now()
	setter := bson.M{
		"$push": bson.M{
			entity.UserFieldInArray(entity.UserFields.Roles): permission,
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

func (r *repo) ListBusinessUsers(ctx context.Context, nickName string, user business.UserDetail) ([]business.User, *i18np.Error) {
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
	return o.ToBusinessUsers(), nil
}

func (r *repo) AdminListAll(ctx context.Context, listConfig list.Config) (*list.Result[*business.AdminListDto], *i18np.Error) {
	filter := bson.M{}
	l, err := r.helper.GetListFilterTransform(ctx, filter, func(o *entity.MongoBusiness) *business.Entity {
		return o.ToBusiness()
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
	li := make([]*business.AdminListDto, 0)
	for _, o := range l {
		dto := &business.AdminListDto{
			UUID:         o.UUID,
			NickName:     o.NickName,
			RealName:     o.RealName,
			BusinessType: string(o.BusinessType),
			IsEnabled:    o.IsEnabled,
			IsVerified:   o.IsVerified,
			IsDeleted:    o.IsDeleted,
			CreatedAt:    o.CreatedAt.String(),
			UpdatedAt:    o.UpdatedAt.String(),
		}
		if o.VerifiedAt != nil {
			dto.VerifiedAt = o.VerifiedAt.String()
		}
		li = append(li, dto)
	}
	return &list.Result[*business.AdminListDto]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          li,
	}, nil
}
