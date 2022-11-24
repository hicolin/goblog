package models

import (
	"goblog/pkg/types"
	"time"
)

type BaseModel struct {
	ID uint64 `gorm:"colum:id;primaryKey;autoIncrement;not null"`

	CreatedAt time.Time `gorm:"colum:created_at;index"`
	UpdatedAt time.Time `gorm:"colum:updated_at;index"`
}

func (a BaseModel) GetStringID() string {
	return types.Uint64ToString(a.ID)
}
