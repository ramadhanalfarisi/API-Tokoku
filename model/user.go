package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TkUser struct {
	UserId           uuid.UUID `json:"ID"`
	UserEmail        string    `json:"userEmail,omitempty"`
	UserPassword     string    `json:"userPassword,omitempty"`
	UserFirstname    string    `json:"userFirstname,omitempty"`
	UserLastname     string    `json:"userLastname,omitempty"`
	UserRole         string    `json:"userRole,omitempty"`
	UserImageProfile string    `json:"userImageProfile,omitempty"`
	DateVerification time.Time `json:"dateVerification,omitempty"`
	IsActive         string    `json:"isActive,omitempty"`
}

func (user *TkUser) InsertUser(db *gorm.DB) error {
	result := db.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (user *TkUser) GetAllUser(db *gorm.DB) ([]TkUser, error) {
	var resUser []TkUser
	err := db.Select("user_id", "user_email", "user_firstname", "user_lastname", "user_role", "user_image_profile", "date_verification", "is_active").Find(&resUser)
	if err.Error != nil {
		return nil, err.Error
	}
	return resUser, nil
}

func (user *TkUser) GetOneUser(db *gorm.DB) (TkUser, error) {
	var resUser TkUser
	err := db.Select("user_id", "user_email", "user_firstname", "user_lastname", "user_role", "user_image_profile", "date_verification", "is_active").Where("user_id = ?", user.UserId).Find(&resUser)
	if err.Error != nil {
		return resUser, err.Error
	}
	return resUser, nil
}

func (user *TkUser) UpdateUser(db *gorm.DB, dataUser map[string]interface{}) (TkUser, error) {
	var resUser TkUser
	err := db.Model(&resUser).Where("location_id = ?", user.UserId).Updates(dataUser)
	if err.Error != nil {
		return resUser, err.Error
	}
	return resUser, nil
}

func (user *TkUser) DeleteUser(db *gorm.DB) (TkUser, error) {
	var resUser TkUser
	err := db.Where("user_id = ?", user.UserId).Delete(&resUser)
	if err.Error != nil {
		return resUser, err.Error
	}
	return resUser, nil
}

func (user *TkUser) LoginUser(db *gorm.DB) (TkUser, error) {
	var resUser TkUser
	err := db.Select("user_id", "user_email", "user_role").Where("user_email = ? AND user_password = ? AND is_active = ?", user.UserEmail, user.UserPassword, "1").Find(&resUser)
	if err.Error != nil {
		return resUser, err.Error
	}
	return resUser, nil
}
