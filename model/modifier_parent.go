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

func (modifierParent *TkModifierParent) SelectAllModifier(db *gorm.DB) ([]TkModifierParent, error) {
	var modifiers []TkModifierParent
	res := db.Preload("ModifierChilds").Find(&modifiers)
	if res.Error != nil {
		return nil, res.Error
	}
	return modifiers, nil
}

func (modifierParent *TkModifierParent) SelectOneModifier(db *gorm.DB) (TkModifierParent, error) {
	var modifier TkModifierParent
	res := db.Where("modifier_parent_id = ?",modifierParent.ModifierParentId).Preload("ModifierChilds").Find(&modifier)
	if res.Error != nil {
		return modifier, res.Error
	}
	return modifier, nil
}

func (modifierParent *TkModifierParent) UpdateModifier(db *gorm.DB) (TkModifierParent, error) {
	var mod TkModifierParent

	data_mod := map[string]interface{}{
		"modifier_parent_name":  modifierParent.ModifierParentName,
		"selected_min":  modifierParent.SelectedMin,
		"selected_max": modifierParent.SelectedMax,
		"modifier_parent_type": modifierParent.ModifierParentType,
	}
	res := db.Model(&mod).Where("modifier_parent_id = ?", modifierParent.ModifierParentId).Updates(data_mod)
	if res.Error != nil {
		return mod, res.Error
	}
	return modifierParent.SelectOneModifier(db)
}

func (modifierParent *TkModifierParent) DeleteModifier(db *gorm.DB) error {
	res := db.Where("modifier_parent_id = ?",modifierParent.ModifierParentId).Find(modifierParent)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
