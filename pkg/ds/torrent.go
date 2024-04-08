package ds

import (
	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/model"
	mt "github.com/szu17dmy/mtorrent-telegram-helper/pkg/mtorrent"
)

func TorrentFilter(torrents []*mt.Torrent, condition func(*mt.Torrent) bool) []*mt.Torrent {
	var result []*mt.Torrent
	for _, t := range torrents {
		if condition(t) {
			result = append(result, t)
		}
	}
	return result
}

func TorrentMap(torrents []*mt.Torrent, transform func(torrent *mt.Torrent) *model.Torrent) []*model.Torrent {
	var result []*model.Torrent
	for _, t := range torrents {
		result = append(result, transform(t))
	}
	return result
}
