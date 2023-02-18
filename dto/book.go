package dto

import "github.com/MasoudHeydari/Fidibo_Interview/entity"

type SearchBookRequest struct {
	Keyword string `json:"keyword"`
}

type SearchBookResponse struct {
	Result []entity.Book `json:"result"`
}
