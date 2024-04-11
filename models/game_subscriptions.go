package models

import (
	"gorm.io/gorm"
)

type GameSubscriptions struct {
	gorm.Model
	ReservationID int
	GameID        string
	Name          string
}

func init() {
	AllModels = append(AllModels, GameSubscriptions{})
}
