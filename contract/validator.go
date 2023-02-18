package contract

import "github.com/MasoudHeydari/Fidibo_Interview/dto"

type (
	ValidateCreateUser func(req dto.CreateUserRequest) error
	ValidateLoginUser  func(req dto.LoginUserRequest) error
)
