package contract

import (
	"context"
	"fidibo_interview/entity"
)

//go:generate mockgen -destination=../mocks/store/user_store_mock.go -package=store_mock  -self_package=github.com/gocastsian/adamak/contract github.com/gocastsian/adamak/contract UserStore

type UserStore interface {
	ValidatorStore
	CreateUser(ctx context.Context, user entity.User) (entity.User, error)
}
