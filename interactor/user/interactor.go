package user

import (
	"context"
	"fmt"
	"github.com/MasoudHeydari/Fidibo_Interview/auth"
	"github.com/MasoudHeydari/Fidibo_Interview/config"
	"github.com/MasoudHeydari/Fidibo_Interview/contract"
	"github.com/MasoudHeydari/Fidibo_Interview/dto"
	"github.com/MasoudHeydari/Fidibo_Interview/entity"
	"time"
)

type Interactor struct {
	store contract.UserStore
}

func New(store contract.UserStore) Interactor {
	return Interactor{store: store}
}

func (i Interactor) CreateUser(ctx context.Context, req dto.CreateUserRequest) (dto.CreateUserResponse, error) {
	user := entity.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	userExists, err := i.store.DoesUserExist(ctx, user.Email)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	if userExists {
		return dto.CreateUserResponse{}, fmt.Errorf("user already exists")
	}

	createdUser, err := i.store.CreateUser(ctx, user)
	if err != nil {
		return dto.CreateUserResponse{}, err
	}

	return dto.CreateUserResponse{User: createdUser}, nil
}

func (i Interactor) Login(ctx context.Context, req dto.LoginUserRequest) (dto.LoginUserResponse, error) {
	exists, err := i.store.DoesUserExist(ctx, req.Email)
	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	if exists {
		tokenMaker, err := auth.NewJWtMaker(config.GetSecretKey())
		if err != nil {
			return dto.LoginUserResponse{}, err
		}

		payload, err := tokenMaker.CreateTokenPayload(req.Email, time.Hour)
		if err != nil {
			return dto.LoginUserResponse{}, err
		}

		return dto.LoginUserResponse{
			AccessToken: payload,
		}, nil
	}

	return dto.LoginUserResponse{}, fmt.Errorf("failed to create access token, user does not exist")
}
