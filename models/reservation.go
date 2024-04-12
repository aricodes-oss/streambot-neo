package models

import (
	"gorm.io/gorm"
)

type Reservation struct {
	gorm.Model
	GuildID              string
	ChannelID            string
	GameSubscriptions    []GameSubscription
	YoutubeSubscriptions []YoutubeSubscription
	BlacklistedUsers     []BlacklistedUser
}

func init() {
	AllModels = append(AllModels, Reservation{})
}
