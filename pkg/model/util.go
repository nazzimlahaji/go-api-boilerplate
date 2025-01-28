package model

import "gorm.io/gorm"

type DB struct {
	*gorm.DB
}
