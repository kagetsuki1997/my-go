package api

import "my-go/internal/doge/store"

type Api interface {
	Doge
}

type api struct {
	Doge
}

func New(store store.Store) Api {
	return &api{Doge: NewDoge(store)}
}
