package mtorrent

import (
	"log"
	"strings"

	"github.com/szu17dmy/mtorrent-telegram-helper/internal/config"
	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/database"
	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/ds"
	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/model"
	mt "github.com/szu17dmy/mtorrent-telegram-helper/pkg/mtorrent"
)

type Dependency interface {
	Configs() *config.Configs
	Storage() *database.Storage
}

const (
	ParamDefaultPageNumber = 1
	ParamDefaultPageSize   = 100
	// ParamDefaultVisible 仅活跃的种子=1
	ParamDefaultVisible    = 1
	ParamModeNormal        = "normal"
	ParamModeNSFW          = "adult"
	ParamDiscountFree      = "FREE"
	ParamPinNormalLevel    = 0
	abstractActivityPrefix = "*活動置頂"
)

func FindNotPushedTorrents(app Dependency) []*model.Torrent {
	ts, err := app.Storage().TorrentRepository.FindNotPushed()
	if err != nil {
		log.Printf("failed to find not pushed torrents, err: %v", err)
		return nil
	}
	return ts
}

func FindExpiredTorrents() []*mt.Torrent {
	// @TODO
	return nil
}

func UpdateTorrentPushed(app Dependency, id string, messageId int) error {
	return app.Storage().TorrentRepository.UpdatePushed(id, 1, messageId)
}

func SaveTorrents(app Dependency, torrents []*mt.Torrent, category string) {
	tf := func(t *mt.Torrent) *model.Torrent {
		dexp, _ := ds.DateParse(t.Status.DiscountExpirationDate)
		pexp, _ := ds.DateParse(t.Status.PinExpirationDate)
		return &model.Torrent{
			RemoteId:               t.Id,
			Name:                   t.Name,
			Category:               category,
			Abstract:               t.Abstract,
			Size:                   t.Size,
			Discount:               t.Status.Discount,
			DiscountExpirationDate: dexp,
			Pin:                    t.Status.Pin,
			PinExpirationDate:      pexp,
			PushState:              0,
		}
	}
	err := app.Storage().TorrentRepository.SaveAll(ds.TorrentMap(torrents, tf))
	if err != nil {
		log.Printf("failed to save torrents, err: %v", err)
	}
}

func GetFreeAndLargeNormalTorrents() ([]*mt.Torrent, error) {
	resp, err := mt.TorrentSearch(&mt.TorrentSearchRequest{
		PageNumber: ParamDefaultPageNumber,
		PageSize:   ParamDefaultPageSize,
		Keyword:    "",
		Visible:    ParamDefaultVisible,
		Mode:       ParamModeNormal,
	})
	if err != nil {
		return nil, err
	}
	return filter(resp.Data.Data), nil
}

func GetFreeAndLargeNSFWTorrents() ([]*mt.Torrent, error) {
	resp, err := mt.TorrentSearch(&mt.TorrentSearchRequest{
		PageNumber: ParamDefaultPageNumber,
		PageSize:   ParamDefaultPageSize,
		Keyword:    "",
		Visible:    ParamDefaultVisible,
		Mode:       ParamModeNSFW,
	})
	if err != nil {
		return nil, err
	}
	return filter(resp.Data.Data), nil
}

func filter(torrents []*mt.Torrent) []*mt.Torrent {
	return ds.TorrentFilter(torrents, func(t *mt.Torrent) bool {
		return t.Status.Discount == ParamDiscountFree &&
			t.Status.Pin != ParamPinNormalLevel &&
			dateAfterNow(t.Status.DiscountExpirationDate) &&
			dateAfterNow(t.Status.PinExpirationDate) &&
			strings.HasPrefix(t.Abstract, abstractActivityPrefix)
	})
}

func dateAfterNow(date string) bool {
	if date == "" {
		return true
	}
	b, err := ds.DateAfterNow(date)
	if err != nil {
		log.Printf("failed to parse date, date: %s", date)
		return false
	}
	return b
}
