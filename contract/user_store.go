package contract

import (
	"context"
	"github.com/MasoudHeydari/Fidibo_Interview/entity"
)

type UserStore interface {
	ValidatorStore
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
}
