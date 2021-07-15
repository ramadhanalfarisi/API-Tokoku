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
