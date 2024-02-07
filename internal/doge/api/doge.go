package api

import (
	"my-go/internal/doge/model"
	"my-go/internal/doge/store"

	"github.com/google/uuid"
)

type Doge interface {
	CreateDoge(req store.CreateDogeParams) (*model.Doge, error)
	GetDoge(id uuid.UUID) (*model.Doge, error)
	ListDoge() ([]*model.Doge, error)
}

type doge struct {
	dogeStore store.DogeStore
}

func NewDoge(store store.Store) Doge {
	return &doge{
		dogeStore: store,
	}
}

func (doge *doge) CreateDoge(req store.CreateDogeParams) (*model.Doge, error) {
	return doge.dogeStore.CreateDoge(req)
}

func (doge *doge) GetDoge(id uuid.UUID) (*model.Doge, error) {
	return doge.dogeStore.GetDoge(id)
}

func (doge *doge) ListDoge() ([]*model.Doge, error) {
	return doge.dogeStore.ListDoge()
}
