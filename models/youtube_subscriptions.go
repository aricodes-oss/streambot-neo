package models

import (
	"gorm.io/gorm"
)

type YoutubeSubscriptions struct {
	gorm.Model
	ReservationID int
	ChannelID     string
}

func init() {
	AllModels = append(AllModels, YoutubeSubscriptions{})
}
