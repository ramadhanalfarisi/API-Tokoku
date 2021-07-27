package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TkParentChildModifiers struct {
	gorm.Model
	ID               interface{}       `json:",omitempty"`
	ModifierParentId uuid.UUID         `json:"parentID"`
	ModifierChildId  uuid.UUID         `json:"childID"`
}

func InsertModifierChildParent(db *gorm.DB, modifierChildParents []TkParentChildModifiers) error {
	modifier_cp := db.Create(modifierChildParents)
	if modifier_cp.Error != nil {
		return modifier_cp.Error
	}
	return nil
}

func DeleteModifierChildParent(db *gorm.DB, modifierChildIds []uuid.UUID, modifierParentId uuid.UUID) error {
	var modChildParent TkParentChildModifiers
	res := db.Where("modifier_parent_id = ? AND NOT modifier_child_id IN ?",modifierParentId, modifierChildIds).Find(&modChildParent)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
