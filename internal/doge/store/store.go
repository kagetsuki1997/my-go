package store

import "gorm.io/gorm"

type Store interface {
	DogeStore
}

type store struct {
	database *gorm.DB
}

func New(database *gorm.DB) Store {
	s := &store{database}

	return s
}
