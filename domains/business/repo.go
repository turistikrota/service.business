package business

import (
	"context"
	"time"

	"github.com/cilloparch/cillop/i18np"
	"github.com/cilloparch/cillop/types/list"
	mongo2 "github.com/turistikrota/service.shared/db/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserDetail struct {
	Name string
	UUID string
}

type Repository interface {
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	GetByIndividual(ctx context.Context, individual Individual) (*Entity, bool, *i18np.Error)
	GetByCorporation(ctx context.Context, corporation Corporation) (*Entity, bool, *i18np.Error)
	GetByNickName(ctx context.Context, nickName string) (*Entity, *i18np.Error)
	CheckNickName(ctx context.Context, nickName string) (bool, *i18np.Error)
	GetWithUser(ctx context.Context, nickName string, user UserDetail) (*EntityWithUserDto, *i18np.Error)
	GetWithUserName(ctx context.Context, nickName string, userName string) (*EntityWithUserDto, *i18np.Error)
	ProfileView(ctx context.Context, nickName string) (*Entity, *i18np.Error)
	ListByUserUUID(ctx context.Context, user UserDetail) ([]BusinessListDto, *i18np.Error)
	ListBusinessUsers(ctx context.Context, nickName string, user UserDetail) ([]User, *i18np.Error)
	ListAsClaim(ctx context.Context, userUUID string) ([]*Entity, *i18np.Error)
	AddUser(ctx context.Context, businessName string, user *User) *i18np.Error
	RemoveUser(ctx context.Context, nickName string, user UserDetail) *i18np.Error
	RemoveUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error
	AddUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error
	SetPreferredLocale(ctx context.Context, nickName string, locale string) *i18np.Error
	Enable(ctx context.Context, nickName string) *i18np.Error
	Verify(ctx context.Context, nickName string) *i18np.Error
	Reject(ctx context.Context, nickName string, reason string) *i18np.Error
	Disable(ctx context.Context, nickName string) *i18np.Error
	Delete(ctx context.Context, nickName string) *i18np.Error
	Recover(ctx context.Context, nickName string) *i18np.Error
	AdminListAll(ctx context.Context, listConfig list.Config) (*list.Result[*AdminListDto], *i18np.Error)
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

func (r *repo) Create(ctx context.Context, business *Entity) (*Entity, *i18np.Error) {
	res, err := r.collection.InsertOne(ctx, business)
	if err != nil {
		return nil, r.factory.Errors.Failed("create")
	}
	business.UUID = res.InsertedID.(primitive.ObjectID).Hex()
	return business, nil
}

func (r *repo) GetByNickName(ctx context.Context, nickName string) (*Entity, *i18np.Error) {
	filter := bson.M{
		fields.NickName: nickName,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *o, nil
}

func (r *repo) SetPreferredLocale(ctx context.Context, nickName string, locale string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.PreferredLocale: locale,
			fields.UpdatedAt:       time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)

}

func (r *repo) GetByIndividual(ctx context.Context, individual Individual) (*Entity, bool, *i18np.Error) {
	filter := bson.M{
		individualField(individualFields.FirstName):   individual.FirstName,
		individualField(individualFields.LastName):    individual.LastName,
		individualField(individualFields.DateOfBirth): individual.DateOfBirth,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, false, err
	}
	if !exist {
		return nil, true, nil
	}
	return *o, false, nil
}

func (r *repo) GetByCorporation(ctx context.Context, corporation Corporation) (*Entity, bool, *i18np.Error) {
	filter := bson.M{
		corporationField(corporationFields.TaxOffice): corporation.TaxOffice,
		corporationField(corporationFields.Title):     corporation.Title,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, false, err
	}
	if !exist {
		return nil, true, nil
	}
	return *o, false, nil
}

func (r *repo) CheckNickName(ctx context.Context, nickName string) (bool, *i18np.Error) {
	filter := bson.M{
		fields.NickName: nickName,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return true, err
	}
	return o != nil && exist, nil
}

func (r *repo) GetWithUser(ctx context.Context, nickName string, user UserDetail) (*EntityWithUserDto, *i18np.Error) {
	filter := bson.M{
		fields.NickName:            nickName,
		userField(userFields.Name): user.Name,
		userField(userFields.UUID): user.UUID,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return (*o).ToEntityWithUserDto(user), nil
}

func (r *repo) GetWithUserName(ctx context.Context, nickName string, userName string) (*EntityWithUserDto, *i18np.Error) {
	filter := bson.M{
		fields.NickName:            nickName,
		userField(userFields.Name): userName,
	}
	o, exist, err := r.helper.GetFilter(ctx, filter)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return (*o).ToEntityWithUserDto(UserDetail{
		Name: userName,
	}), nil
}

func (r *repo) ProfileView(ctx context.Context, nickName string) (*Entity, *i18np.Error) {
	filter := bson.M{
		fields.NickName:  nickName,
		fields.IsEnabled: true,
		fields.IsDeleted: false,
	}
	opts := options.FindOne().SetProjection(bson.M{
		fields.UUID:        0,
		fields.Users:       0,
		fields.IsEnabled:   0,
		fields.Corporation: 0,
		fields.Individual:  0,
		fields.ActivatedAt: 0,
		fields.DisabledAt:  0,
		fields.VerifiedAt:  0,
		fields.UpdatedAt:   0,
	})
	o, exist, err := r.helper.GetFilter(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return *o, nil
}

func (r *repo) ListByUserUUID(ctx context.Context, user UserDetail) ([]BusinessListDto, *i18np.Error) {
	filter := bson.M{
		userField(userFields.Name): user.Name,
		userField(userFields.UUID): user.UUID,
	}
	opts := options.Find().SetProjection(bson.M{
		fields.UUID:        0,
		fields.Users:       0,
		fields.Corporation: 0,
		fields.Individual:  0,
		fields.ActivatedAt: 0,
		fields.DisabledAt:  0,
		fields.VerifiedAt:  0,
	})
	res, err := r.helper.GetListFilter(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	li := make([]BusinessListDto, 0)
	for _, o := range res {
		li = append(li, BusinessListDto{
			NickName:     o.NickName,
			RealName:     o.RealName,
			BusinessType: string(o.BusinessType),
			RejectReason: o.RejectReason,
			IsVerified:   o.IsVerified,
			IsEnabled:    o.IsEnabled,
			IsDeleted:    o.IsDeleted,
			UpdatedAt:    o.UpdatedAt,
		})
	}
	return li, nil
}

func (r *repo) ListAsClaim(ctx context.Context, userUUID string) ([]*Entity, *i18np.Error) {
	filter := bson.M{
		userField(userFields.UUID): userUUID,
	}
	opts := options.Find()
	return r.helper.GetListFilter(ctx, filter, opts)
}

func (r *repo) AddUser(ctx context.Context, businessName string, user *User) *i18np.Error {
	filter := bson.M{
		fields.NickName:            businessName,
		userField(userFields.Name): bson.M{"$ne": user.Name},
	}
	setter := bson.M{
		"$addToSet": bson.M{
			fields.Users: bson.M{
				userFields.UUID:   user.UUID,
				userFields.Name:   user.Name,
				userFields.Roles:  user.Roles,
				userFields.JoinAt: user.JoinAt,
			},
		},
		"$set": bson.M{
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) RemoveUser(ctx context.Context, nickName string, user UserDetail) *i18np.Error {
	filter := bson.M{
		fields.NickName:            nickName,
		userField(userFields.Name): user.Name,
	}
	t := time.Now()
	setter := bson.M{
		"$pull": bson.M{
			fields.Users: bson.M{
				userFields.Name: user.Name,
			},
		},
		"$set": bson.M{
			fields.UpdatedAt: t,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) RemoveUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		fields.NickName:            nickName,
		userField(userFields.Name): user.Name,
	}
	t := time.Now()
	setter := bson.M{
		"$pull": bson.M{
			userFieldInArray(userFields.Roles): permission,
		},
		"$set": bson.M{
			fields.UpdatedAt: t,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) AddUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error {
	filter := bson.M{
		fields.NickName:                    nickName,
		userField(userFields.Name):         user.Name,
		userFieldInArray(userFields.Roles): bson.M{"$ne": permission},
	}
	t := time.Now()
	setter := bson.M{
		"$push": bson.M{
			userFieldInArray(userFields.Roles): permission,
		},
		"$set": bson.M{
			fields.UpdatedAt: t,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Enable(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsEnabled: true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Verify(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	t := time.Now()
	setter := bson.M{
		"$set": bson.M{
			fields.RejectReason: nil,
			fields.IsVerified:   true,
			fields.UpdatedAt:    t,
			fields.VerifiedAt:   t,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Reject(ctx context.Context, nickName string, reason string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	t := time.Now()
	setter := bson.M{
		"$set": bson.M{
			fields.RejectReason: reason,
			fields.UpdatedAt:    t,
			fields.IsVerified:   false,
			fields.VerifiedAt:   nil,
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Disable(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsEnabled: false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Delete(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsDeleted: true,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) Recover(ctx context.Context, nickName string) *i18np.Error {
	filter := bson.M{
		fields.NickName: nickName,
	}
	setter := bson.M{
		"$set": bson.M{
			fields.IsDeleted: false,
			fields.UpdatedAt: time.Now(),
		},
	}
	return r.helper.UpdateOne(ctx, filter, setter)
}

func (r *repo) ListBusinessUsers(ctx context.Context, nickName string, user UserDetail) ([]User, *i18np.Error) {
	filter := bson.M{
		fields.NickName:            nickName,
		userField(userFields.Name): user.Name,
	}
	opts := options.FindOne().SetProjection(bson.M{
		fields.Users: 1,
	})
	o, exist, err := r.helper.GetFilter(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, r.factory.Errors.NotFound()
	}
	return (*o).Users, nil
}

func (r *repo) AdminListAll(ctx context.Context, listConfig list.Config) (*list.Result[*AdminListDto], *i18np.Error) {
	filter := bson.M{}
	l, err := r.helper.GetListFilter(ctx, filter)
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
	li := make([]*AdminListDto, 0)
	for _, o := range l {
		dto := &AdminListDto{
			UUID:            o.UUID,
			NickName:        o.NickName,
			RealName:        o.RealName,
			BusinessType:    string(o.BusinessType),
			IsEnabled:       o.IsEnabled,
			IsVerified:      o.IsVerified,
			IsDeleted:       o.IsDeleted,
			Application:     string(o.Application),
			PreferredLocale: o.PreferredLocale,
			CreatedAt:       o.CreatedAt.String(),
			UpdatedAt:       o.UpdatedAt.String(),
		}
		if o.VerifiedAt != nil {
			dto.VerifiedAt = o.VerifiedAt.String()
		}
		li = append(li, dto)
	}
	return &list.Result[*AdminListDto]{
		IsNext:        filtered > listConfig.Offset+listConfig.Limit,
		IsPrev:        listConfig.Offset > 0,
		FilteredTotal: filtered,
		Total:         total,
		Page:          listConfig.Offset/listConfig.Limit + 1,
		List:          li,
	}, nil
}
