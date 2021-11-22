package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TkUser struct {
	UserId              uuid.UUID `json:"ID"`
	UserEmail           string    `json:"userEmail,omitempty"`
	UserPassword        string    `json:"userPassword,omitempty"`
	ConfirmUserPassword string    `json:"confirmUserPassword,omitempty"`
	UserFirstname       string    `json:"userFirstname,omitempty"`
	UserLastname        string    `json:"userLastname,omitempty"`
	UserRole            string    `json:"userRole,omitempty"`
	UserImageProfile    string    `json:"userImageProfile,omitempty"`
	DateVerification    string    `json:"dateVerification,omitempty"`
	IsActive            string    `json:"isActive,omitempty"`
}

type TkUserRegister struct {
	UserId              uuid.UUID `json:"ID"`
	UserEmail           string    `json:"userEmail,omitempty" validate:"required,email"`
	UserPassword        string    `json:"userPassword,omitempty" validate:"required"`
	ConfirmUserPassword string    `json:"confirmUserPassword,omitempty" validate:"required,eqfield=UserPassword"`
	UserFirstname       string    `json:"userFirstname,omitempty" validate:"required,alpha"`
	UserLastname        string    `json:"userLastname,omitempty" validate:"required,alpha"`
	UserRole            string    `json:"userRole,omitempty" validate:"required,alpha"`
	UserImageProfile    interface{}    `json:"userImageProfile,omitempty"`
	DateVerification    string    `json:"dateVerification,omitempty"`
	IsActive            string    `json:"isActive,omitempty"`
}

type TkUserLogin struct {
	UserEmail           string    `json:"userEmail,omitempty" validate:"required,email"`
	UserPassword        string    `json:"userPassword,omitempty" validate:"required"`
}


func (user *TkUserRegister) InsertUser(db *gorm.DB) error {
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

func (user *TkUserLogin) LoginUser(db *gorm.DB) (TkUser, error) {
	var resUser TkUser
	err := db.Select("user_id", "user_email", "user_role").Where("user_email = ? AND user_password = ? AND is_active = ?", user.UserEmail, user.UserPassword, "1").Find(&resUser)
	if err.Error != nil {
		return resUser, err.Error
	}
	return resUser, nil
}
