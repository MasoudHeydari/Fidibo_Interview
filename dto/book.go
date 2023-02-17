package dto

import "fidibo_interview/entity"

type SearchBookRequest struct {
	Keyword string `json:"keyword"`
}

type SearchBookResponse struct {
	Result []entity.Book `json:"result"`
}
