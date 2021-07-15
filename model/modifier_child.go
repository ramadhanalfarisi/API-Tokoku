package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TkModifierChild struct {
	gorm.Model
	ModifierChildId     uuid.UUID `json:"ID"`
	ModifierChildName   string    `json:"modifierName"`
	ModifierChildPrice  float64   `json:"modifierPrice"`
	ModifierChildDesc   string    `json:"modifierDesc"`
}

func (modifierChild *TkModifierChild) InsertModifierChild(db *gorm.DB) error {
	modifier_c := db.Create(modifierChild)
	if modifier_c.Error != nil {
		return modifier_c.Error
	}
	return nil
}
