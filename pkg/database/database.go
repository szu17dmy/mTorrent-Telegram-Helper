package database

import (
	"log"

	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Storage struct {
	TorrentRepository TorrentRepository
}

func New() *Storage {
	db, err := gorm.Open(sqlite.Open("./sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database, err: %v", err)
	}
	if err := db.AutoMigrate(&model.Torrent{}); err != nil {
		log.Fatalf("failed to migrate, err: %v", err)
	}
	return &Storage{
		TorrentRepository: NewGormTorrentRepository(db),
	}
}
