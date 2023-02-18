package contract

import (
	"context"
	"github.com/MasoudHeydari/Fidibo_Interview/dto"
)

type BookInteractor interface {
	SearchBook(context.Context, dto.SearchBookRequest) (dto.SearchBookResponse, error)
}
