package telegram

import (
	"bytes"
	"embed"
	_ "embed"
	"html/template"

	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/fs"
	mt "github.com/szu17dmy/mtorrent-telegram-helper/pkg/mtorrent"
	tg "github.com/szu17dmy/mtorrent-telegram-helper/pkg/telegram"

	"gopkg.in/telebot.v3"
)

var (
	//go:embed torrent_message_template.html
	torrentMessageTemplate embed.FS
	//go:embed nsfw_torrent_message_template.html
	nsfwTorrentMessageTemplate embed.FS
)

type Torrent struct {
	Id         string
	Title      string
	Abstract   string
	Size       string
	Expiration string
}

func SendTorrentMessage(torrent *mt.Torrent) (*telebot.Message, error) {
	tpl, err := template.ParseFS(torrentMessageTemplate, "torrent_message_template.html")
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, parseTorrent(torrent))
	if err != nil {
		return nil, err
	}
	return tg.SendHtml(buf.String())
}

func SendNSFWTorrentMessage(torrent *mt.Torrent) (*telebot.Message, error) {
	tpl, err := template.ParseFS(nsfwTorrentMessageTemplate, "nsfw_torrent_message_template.html")
	if err != nil {
		return nil, err
	}
	var buf bytes.Buffer
	err = tpl.Execute(&buf, parseTorrent(torrent))
	if err != nil {
		return nil, err
	}
	return tg.SendHtml(buf.String())
}

func parseTorrent(torrent *mt.Torrent) *Torrent {
	return &Torrent{
		Id:         torrent.Id,
		Title:      torrent.Name,
		Abstract:   torrent.Abstract,
		Size:       fs.Parse(torrent.Size).String(),
		Expiration: torrent.Status.PinExpirationDate,
	}
}
