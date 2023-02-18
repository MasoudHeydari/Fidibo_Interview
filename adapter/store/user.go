package store

import (
	"context"
	"github.com/MasoudHeydari/Fidibo_Interview/entity"
)

func (m MySQLStore) CreateUser(ctx context.Context, user entity.User) (entity.User, error) {
	u := mapFromUserEntity(user)

	if err := m.db.WithContext(ctx).Create(&u).Error; err != nil {
		return entity.User{}, err
	}

	return mapToUserEntity(u), nil
}
