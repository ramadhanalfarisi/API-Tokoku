package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TkProductModifier struct {
	gorm.Model
	ID             interface{}       `json:",omitempty"`
	ModifierId     uuid.UUID         `json:"modifierId,omitempty"`
	ProductId      uuid.UUID         `json:"productId,omitempty"`
}

func (productModifier *TkProductModifier) InsertProductModifier(db *gorm.DB, productModifiers []TkProductModifier) error {
	modifier := db.Create(&productModifiers)
	if modifier.Error != nil {
		return modifier.Error
	}
	return nil
}

func (productModifier *TkProductModifier) DeleteProductModifier(db *gorm.DB) error {
	delete := db.Where("product_id = ?",productModifier.ProductId).Delete(productModifier)
	if delete.Error != nil {
		return delete.Error
	}
	return nil
}
