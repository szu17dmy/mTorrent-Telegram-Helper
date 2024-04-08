package model

import (
	"time"

	"gorm.io/gorm"
)

type Torrent struct {
	gorm.Model
	RemoteId               string `gorm:"uniqueIndex"`
	Name                   string
	Category               string
	Abstract               string
	Size                   string
	Discount               string
	DiscountExpirationDate time.Time
	Pin                    int
	PinExpirationDate      time.Time
	PushState              int
	MessageId              int
}
