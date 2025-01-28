package model

import (
	"main/pkg/entity"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}

func (db *DB) FetchUserIdentity(email string) (*entity.UserIdentity, error) {
	userIdentity := new(entity.UserIdentity)

	if err := db.Raw(`
		SELECT id, name, email
		FROM users
		WHERE email = ?
		LIMIT 1
	`, email).Scan(&userIdentity).Error; err != nil {
		return nil, err
	}

	return userIdentity, nil
}
