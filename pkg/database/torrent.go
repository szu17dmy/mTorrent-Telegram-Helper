package database

import (
	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TorrentRepository interface {
	FindNotPushed() ([]*model.Torrent, error)
	SaveAll(torrents []*model.Torrent) error
	UpdatePushed(remoteId string, pushState int, messageId int) error
}

type GormTorrentRepository struct {
	db *gorm.DB
}

func NewGormTorrentRepository(db *gorm.DB) *GormTorrentRepository {
	return &GormTorrentRepository{db: db}
}

func (repo *GormTorrentRepository) FindNotPushed() ([]*model.Torrent, error) {
	var torrents []*model.Torrent
	r := repo.db.Where("push_state = ?", 0).Find(&torrents)
	if r.Error != nil {
		return nil, r.Error
	}
	return torrents, nil
}

func (repo *GormTorrentRepository) SaveAll(torrents []*model.Torrent) error {
	r := repo.db.Clauses(clause.OnConflict{DoNothing: true}).
		CreateInBatches(torrents, 100)
	return r.Error
}

func (repo *GormTorrentRepository) UpdatePushed(remoteId string, pushState int, messageId int) error {
	r := repo.db.Model(&model.Torrent{}).
		Where("remote_id = ?", remoteId).
		Updates(&model.Torrent{
			PushState: pushState,
			MessageId: messageId,
		})
	return r.Error
}
