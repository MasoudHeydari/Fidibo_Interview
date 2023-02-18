package contract

import "fidibo_interview/dto"

type (
	ValidateCreateUser func(req dto.CreateUserRequest) error
	ValidateLoginUser  func(req dto.LoginUserRequest) error
)
