package models

import (
	"gorm.io/gorm"
)

type BlacklistedUser struct {
	gorm.Model
	ReservationID int
	UserID        string
}

func init() {
	AllModels = append(AllModels, BlacklistedUser{})
}
