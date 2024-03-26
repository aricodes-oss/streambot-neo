package models

import (
	"gorm.io/gorm"
)

type Stream struct {
	gorm.Model
}

func init() {
	AllModels = append(AllModels, Stream{})
}
