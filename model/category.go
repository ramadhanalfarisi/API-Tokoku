package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TkCategory struct {
	gorm.Model
	CategoryId   uuid.UUID   `json:"ID"`
	CategoryName string      `json:"categoryName" validate:"required"`
	Products     []TkProduct `gorm:"foreignKey:CategoryId;references:CategoryId" json:"products,omitempty"`
	Isactive     string      `json:"isActive,omitempty" validate:"required"`
	UserId       uuid.UUID   `json:"userId,omitempty" validate:"required"`
	CreatedAt    string      `json:"omitempty"`
	UpdatedAt    string      `json:"omitempty"`
	DeletedAt    string      `json:"omitempty"`
}

type Categories []TkCategory

func (category *TkCategory) InsertCategory(db *gorm.DB) error {
	cat := db.Create(category)
	if cat.Error != nil {
		return cat.Error
	}
	return nil
}

func (category *TkCategory) SelectAllCategory(db *gorm.DB) (Categories, error) {
	var categories []TkCategory
	cat := db.Select("category_id", "category_name", "isactive").Where("user_id = ?", category.UserId).Find(&categories)
	if cat.Error != nil {
		return nil, cat.Error
	}
	return categories, nil
}

func (category *TkCategory) SelectOneCategory(db *gorm.DB) (TkCategory, error) {
	var cat TkCategory
	res := db.Select("category_id", "category_name", "isactive").Where("category_id = ? AND user_id = ?", category.CategoryId, category.UserId).Find(&cat)
	if res.Error != nil {
		return cat, res.Error
	}
	return cat, nil
}

func (category *TkCategory) SelectAllMenu(db *gorm.DB) (Categories, error) {
	var categories []TkCategory
	cat := db.Preload("Products.Modifiers.ModifierChilds").Preload(clause.Associations).Find(&categories)
	if cat.Error != nil {
		return nil, cat.Error
	}
	return categories, nil
}

func (category *TkCategory) UpdateCategory(db *gorm.DB) (TkCategory, error) {
	var cat TkCategory
	res := db.Model(&cat).Where("category_id = ?", category.CategoryId).Update("category_name", category.CategoryName)
	if res.Error != nil {
		return cat, res.Error
	}
	return category.SelectOneCategory(db)
}

func (category *TkCategory) DeleteCategory(db *gorm.DB) error {
	var cat TkCategory
	del := db.Where("category_id = ?", category.CategoryId).Delete(&cat)
	if del.Error != nil {
		return del.Error
	}
	return nil
}
