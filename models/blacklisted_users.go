package models

import (
	"gorm.io/gorm"
)

type BlacklistedUsers struct {
	gorm.Model
	ReservationID int
	UserID        string
}

func init() {
	AllModels = append(AllModels, BlacklistedUsers{})
}
