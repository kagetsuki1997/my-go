package store

import (
	"my-go/internal/doge/model"

	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

type DogeStore interface {
	CreateDoge(arg CreateDogeParams) (*model.Doge, error)
	GetDoge(id uuid.UUID) (*model.Doge, error)
	ListDoge() ([]*model.Doge, error)
}

type CreateDogeParams struct {
	Name        string
	MagicNumber uint32
	Type        model.DogeType
}

func (store *store) CreateDoge(arg CreateDogeParams) (*model.Doge, error) {
	doge := &model.Doge{Name: arg.Name, MagicNumber: arg.MagicNumber, Type: arg.Type}

	if err := store.database.Clauses(clause.Returning{}).Create(doge).Error; err != nil {
		return nil, err
	}

	return doge, nil
}

func (store *store) GetDoge(id uuid.UUID) (*model.Doge, error) {
	doge := &model.Doge{}
	if err := store.database.First(&doge, id).Error; err != nil {
		return nil, err
	}

	return doge, nil
}

func (store *store) ListDoge() ([]*model.Doge, error) {
	var doges []*model.Doge
	if err := store.database.Find(&doges).Order("createTimestamp DESC").Error; err != nil {
		return nil, err
	}

	return doges, nil
}
