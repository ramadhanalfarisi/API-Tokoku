package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TkProduct struct {
	gorm.Model
	ProductId       uuid.UUID          `json:"ID"`
	ProductName     string             `json:"productName"`
	ProductDesc     string             `json:"productDesc"`
	ProductPrice    string             `json:"productPrice"`
	ProductImage    string             `json:"productImage"`
	CategoryId      uuid.UUID          `json:"categoryId,omitempty"`
	LocationId      uuid.UUID          `json:"locationId,omitempty"`
	CategoryProduct *TkCategory         `gorm:"foreignKey:CategoryId;references:CategoryId" json:"categoryProduct,omitempty"`
	Modifiers       []TkModifierParent `gorm:"foreignKey:ProductId;references:ProductId" json:"productModifier,omitempty"`
}

func (product *TkProduct) InsertProduct(db *gorm.DB) error {
	prod := db.Create(product)
	if prod.Error != nil {
		return prod.Error
	}
	return nil
}
