package model

import (
	"time"

	"github.com/google/uuid"
)

type Doge struct {
	ID              uuid.UUID `gorm:"type:uuid;default:gen_random_uuid()" json:"id"`
	Name            string    `gorm:"not null" json:"name"`
	MagicNumber     uint32    `gorm:"not null" json:"magicNumber"`
	Type            DogeType  `gorm:"not null;type:doge_type;" json:"type"`
	CreateTimestamp time.Time `gorm:"not null;type:timestamp(0) without time zone;default:now()" json:"createTimestamp"`
}

func (Doge) TableName() string {
	return "doge.doge"
}

type DogeType string

const (
	DogeTypeDoge             DogeType = "Doge"
	DogeTypeDogeCry          DogeType = "DogeCry"
	DogeTypeDogeBuffed       DogeType = "DogeBuffed"
	DogeTypeDogeKachitoritai DogeType = "DogeKachitoritai"
)
