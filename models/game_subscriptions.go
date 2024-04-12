package models

import (
	"gorm.io/gorm"
)

type GameSubscription struct {
	gorm.Model
	ReservationID int
	GameID        string
	Name          string
}

func init() {
	AllModels = append(AllModels, GameSubscription{})
}
