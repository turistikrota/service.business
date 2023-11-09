package owner

import (
	"context"

	"github.com/mixarchitecture/i18np"
	"github.com/mixarchitecture/microp/types/list"
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
	GetWithUser(ctx context.Context, nickName string, user UserDetail) (*EntityWithUser, *i18np.Error)
	GetWithUserName(ctx context.Context, nickName string, userName string) (*EntityWithUser, *i18np.Error)
	ProfileView(ctx context.Context, nickName string) (*Entity, *i18np.Error)
	ListByUserUUID(ctx context.Context, user UserDetail) ([]*Entity, *i18np.Error)
	ListOwnershipUsers(ctx context.Context, nickName string, user UserDetail) ([]User, *i18np.Error)
	AddUser(ctx context.Context, ownerUUID string, user *User) *i18np.Error
	RemoveUser(ctx context.Context, nickName string, user UserDetail) *i18np.Error
	RemoveUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error
	AddUserPermission(ctx context.Context, nickName string, user UserDetail, permission string) *i18np.Error
	Enable(ctx context.Context, nickName string) *i18np.Error
	Verify(ctx context.Context, nickName string) *i18np.Error
	Reject(ctx context.Context, nickName string, reason string) *i18np.Error
	Disable(ctx context.Context, nickName string) *i18np.Error
	Delete(ctx context.Context, nickName string) *i18np.Error
	Recover(ctx context.Context, nickName string) *i18np.Error
	AdminListAll(ctx context.Context, listConfig list.Config) (*list.Result[*AdminListDto], *i18np.Error)
}
