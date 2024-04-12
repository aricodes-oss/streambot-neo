package models

import (
	"gorm.io/gorm"
)

type YoutubeSubscription struct {
	gorm.Model
	ReservationID int
	ChannelID     string
}

func init() {
	AllModels = append(AllModels, YoutubeSubscription{})
}
