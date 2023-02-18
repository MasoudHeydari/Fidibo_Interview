package contract

import (
	"context"
	"fidibo_interview/dto"
)

type UserInteractor interface {
	CreateUser(context.Context, dto.CreateUserRequest) (dto.CreateUserResponse, error)
	Login(context.Context, dto.LoginUserRequest) (dto.LoginUserResponse, error)
}
