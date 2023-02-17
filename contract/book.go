package contract

import (
	"context"
	"fidibo_interview/dto"
)

type BookInteractor interface {
	SearchBook(context.Context, dto.SearchBookRequest) (dto.SearchBookResponse, error)
}
