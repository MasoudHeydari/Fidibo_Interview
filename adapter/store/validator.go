package store

import (
	"context"
	"gorm.io/gorm"
)

func (m MySQLStore) DoesUserExist(ctx context.Context, userEmail string) (bool, error) {
	if err := m.db.WithContext(ctx).Where("email = ?", userEmail).First(&User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}

		return false, err
	}

	return true, nil
}
