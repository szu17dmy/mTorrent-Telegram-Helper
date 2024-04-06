package ds

import (
	mt "github.com/szu17dmy/mtorrent-telegram-helper/pkg/mtorrent"
)

func TorrentFilter(torrents []*mt.Torrent, condition func(*mt.Torrent) bool) []*mt.Torrent {
	var result []*mt.Torrent
	for _, v := range torrents {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}
