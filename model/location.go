package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TkLocation struct {
	LocationId       uuid.UUID `json:"ID"`
	UserId           uuid.UUID `json:"userId"`
	LocationName     string    `json:"locationName,omitempty" validate:"required"`
	LocationAddress  string    `json:"locationAddress,omitempty" validate:"required"`
	LocationPhone    string    `json:"locationPhone,omitempty" validate:"required"`
	LocationCity     string    `json:"locationCity,omitempty"`
	LocationProvince string    `json:"locationProvince,omitempty"`
	LocationCountry  string    `json:"locationCountry,omitempty"`
}

func (location *TkLocation) InsertLocation(db *gorm.DB) error {
	loc := db.Create(location)
	if loc.Error != nil {
		return loc.Error
	}
	return nil
}

func (location *TkLocation) SelectAllLocation(db *gorm.DB) ([]TkLocation, error) {
	var locations []TkLocation
	loc := db.Find(&locations)
	if loc.Error != nil {
		return nil, loc.Error
	}
	return locations, nil
}

func (location *TkLocation) SelectOneLocation(db *gorm.DB) (TkLocation, error) {
	var loc TkLocation
	res := db.Where("product_id = ?", location.LocationId).Find(&loc)
	if res.Error != nil {
		return loc, res.Error
	}
	return loc, nil
}

func (location *TkLocation) UpdateLocation(db *gorm.DB) (TkLocation, error) {
	var loc TkLocation
	data_loc := map[string]interface{}{
		"location_name":     location.LocationName,
		"location_address":  location.LocationAddress,
		"location_phone":    location.LocationPhone,
		"location_city":     location.LocationCity,
		"location_province": location.LocationProvince,
		"location_country":  location.LocationCountry,
	}
	res := db.Model(&loc).Where("location_id = ?", location.LocationId).Updates(data_loc)
	if res.Error != nil {
		return loc, res.Error
	}
	return location.SelectOneLocation(db)
}

func (location *TkLocation) DeleteLocation(db *gorm.DB) error {
	var loc TkLocation
	del := db.Where("location_id = ?", location.LocationId).Delete(&loc)
	if del.Error != nil {
		return del.Error
	}
	return nil
}
