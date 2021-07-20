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

func (product *TkProduct) SelectAllProduct(db *gorm.DB) ([]TkProduct, error) {
	var prods []TkProduct
	res := db.Preload("CategoryProduct").Preload("Modifiers.ModifierChilds").Find(&prods)
	if res.Error != nil {
		return prods, res.Error
	}
	return prods, nil
}

func (product *TkProduct) SelectOneProduct(db *gorm.DB) (TkProduct, error) {
	var prod TkProduct
	res := db.Where("product_id = ?",product.ProductId).Preload("CategoryProduct").Preload("Modifiers.ModifierChilds").Find(&prod)
	if res.Error != nil {
		return prod, res.Error
	}
	return prod, nil
}

func (product *TkProduct) UpdateProduct(db *gorm.DB) (TkProduct, error) {
	var prod TkProduct

	data_prod := map[string]interface{}{
		"product_name" : product.ProductName,
		"product_desc" : product.ProductDesc,
		"product_price" : product.ProductPrice,
		"product_image" : product.ProductImage,
		"category_id" : product.CategoryId,
	}
	res := db.Model(&prod).Where("product_id = ?",product.ProductId).Updates(data_prod)
	if res.Error != nil {
		return prod, res.Error
	}
	return product.SelectOneProduct(db)
}

func (product *TkProduct) DeleteProduct(db *gorm.DB) error {
	res := db.Where("product_id = ?",product.ProductId).Delete(product)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
