package mtorrent

import (
	"log"

	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/ds"
	mt "github.com/szu17dmy/mtorrent-telegram-helper/pkg/mtorrent"
)

const (
	ParamDefaultPageNumber = 1
	ParamDefaultPageSize   = 25
	// ParamDefaultVisible 仅活跃的种子=1
	ParamDefaultVisible = 1
	ParamModeNormal     = "normal"
	ParamModeNSFW       = "adult"
	ParamDiscountFree   = "FREE"
	ParamPinNormalLevel = 0
)

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
	return ds.TorrentFilter(torrents, func(torrent *mt.Torrent) bool {
		return torrent.Status.Discount == ParamDiscountFree &&
			torrent.Status.Pin != ParamPinNormalLevel &&
			dateAfterNow(torrent.Status.DiscountExpirationDate) &&
			dateAfterNow(torrent.Status.PinExpirationDate)
	})
}

func dateAfterNow(date string) bool {
	b, err := ds.DateAfterNow(date)
	if err != nil {
		log.Printf("failed to parse date, date: %s", date)
		return false
	}
	return b
}
