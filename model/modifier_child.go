package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TkModifierChild struct {
	gorm.Model
	ModifierChildId     uuid.UUID `json:"ID"`
	ModifierChildName   string    `json:"modifierName" validate:"required"`
	ModifierChildPrice  float64   `json:"modifierPrice" validate:"required, numeric"`
	ModifierChildDesc   string    `json:"modifierDesc"`
}

func InsertModifierChild(db *gorm.DB, modifierChilds []TkModifierChild) error {
	modifier_c := db.Create(modifierChilds)
	if modifier_c.Error != nil {
		return modifier_c.Error
	}
	return nil
}

func (modifierChild *TkModifierChild) SelectOneModifierChild(db *gorm.DB) (TkModifierChild, error) {
	var modifier TkModifierChild
	res := db.Where("modifier_child_id = ?",modifierChild.ModifierChildId).Preload("ModifierChilds").Find(&modifier)
	if res.Error != nil {
		return modifier, res.Error
	}
	return modifier, nil
}

func (modifierChild *TkModifierChild) UpdateModifierChild(db *gorm.DB) error {
	var mod TkModifierChild

	data_mod := map[string]interface{}{
		"modifier_child_name":  modifierChild.ModifierChildName,
		"modifier_child_price":  modifierChild.ModifierChildPrice,
		"modifier_child_desc": modifierChild.ModifierChildDesc,
	}
	res := db.Model(&mod).Where("modifier_parent_id = ?", modifierChild.ModifierChildId).Updates(data_mod)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
