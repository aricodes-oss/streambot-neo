package models

import (
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	GuildID   string
	ChannelID string
	GameID    string
}

func init() {
	AllModels = append(AllModels, Reservation{})
}
