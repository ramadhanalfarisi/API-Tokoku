package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TkModifierParent struct {
	gorm.Model
	ModifierParentId   uuid.UUID         `json:"ID"`
	ModifierParentName string            `json:"modifierName"`
	SelectedMin        uint8             `json:"selectedMin"`
	SelectedMax        uint8             `json:"selectedMax"`
	ModifierParentType string            `json:"modifierType"`
	ProductId          uuid.UUID         `json:"productId,omitempty"`
	ModifierChilds     []TkModifierChild `gorm:"many2many:tk_parent_child_modifiers;foreignKey:ModifierParentId;joinForeignKey:ModifierParentId;References:ModifierChildId;JoinReferences:ModifierChildId" json:"modifierChilds,omitempty"`
}

func (modifierParent *TkModifierParent) InsertModifierParent(db *gorm.DB) error {
	modifier_p := db.Create(modifierParent)
	if modifier_p.Error != nil {
		return modifier_p.Error
	}
	return nil
}
