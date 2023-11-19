package invite

import (
	"context"

	"github.com/mixarchitecture/i18np"
)

type Repository interface {
	Create(ctx context.Context, entity *Entity) (*Entity, *i18np.Error)
	GetByUUID(ctx context.Context, uuid string) (*Entity, *i18np.Error)
	GetByBusinessUUID(ctx context.Context, businessUUID string) ([]*Entity, *i18np.Error)
	GetByEmail(ctx context.Context, email string) ([]*Entity, *i18np.Error)
	Use(ctx context.Context, uuid string) *i18np.Error
	Delete(ctx context.Context, uuid string) *i18np.Error
}
