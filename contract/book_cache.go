package contract

import "fidibo_interview/entity"

type BookCache interface {
	Set(key string, value *[]entity.Book) error
	Get(key string) (*[]entity.Book, error)
}
