package contract

import (
	"context"
	"github.com/MasoudHeydari/Fidibo_Interview/dto"
)

type UserInteractor interface {
	CreateUser(context.Context, dto.CreateUserRequest) (dto.CreateUserResponse, error)
	Login(context.Context, dto.LoginUserRequest) (dto.LoginUserResponse, error)
}
