package contract

import "github.com/MasoudHeydari/Fidibo_Interview/entity"

type BookCache interface {
	Set(key string, value *[]entity.Book) error
	Get(key string) (*[]entity.Book, error)
}
