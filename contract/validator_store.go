package contract

import "context"

type ValidatorStore interface {
	DoesUserExist(ctx context.Context, userEmail string) (bool, error)
}
