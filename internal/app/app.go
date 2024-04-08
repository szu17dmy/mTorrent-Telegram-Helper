package app

import (
	"log"
	"sync"
	"time"

	"github.com/szu17dmy/mtorrent-telegram-helper/internal/config"
	mt "github.com/szu17dmy/mtorrent-telegram-helper/internal/mtorrent"
	tg "github.com/szu17dmy/mtorrent-telegram-helper/internal/telegram"
	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/database"
	"github.com/szu17dmy/mtorrent-telegram-helper/pkg/ds"

	"github.com/robfig/cron/v3"
	"gopkg.in/telebot.v3"
)

const (
	defaultDelaySeconds = 30
)

var (
	pushMutex sync.Mutex
)

type App struct {
	configs *config.Configs
	storage *database.Storage
}

func (a *App) Configs() *config.Configs {
	return a.configs
}

func (a *App) Storage() *database.Storage {
	return a.storage
}

func NewApp() *App {
	app := &App{
		storage: database.New(),
	}
	app.configs = &config.Configs{
		Jobs: []*config.Job{
			{
				Name: "Fetch",
				Spec: "*/15 * * * *",
				Func: app.FetchAndSaveTorrents,
			},
			{
				Name: "Push",
				Spec: "* * * * *",
				Func: app.PushMessage,
			},
			{
				Name: "Destruct",
				Spec: "* * * * *",
				Func: app.DestructMessage,
			},
		},
	}
	return app
}

func (a *App) SetupCronJob(c *cron.Cron) {
	for _, j := range a.Configs().Jobs {
		_, err := c.AddFunc(j.Spec, j.Func)
		if err != nil {
			log.Fatalf("failed to add cron job %s, err: %v", j.Name, err)
		}
		log.Printf("added cron job %s, spec: %s", j.Name, j.Spec)
	}
}

func (a *App) FetchAndSaveTorrents() {
	time.Sleep(time.Second * time.Duration(ds.RandInt(defaultDelaySeconds)))
	ts, err := mt.GetFreeAndLargeNormalTorrents()
	if err != nil {
		log.Printf("failed to fetch normal torrents, err: %v", err)
		return
	}
	log.Printf("normal torrents fetched, amount: %d", len(ts))
	mt.SaveTorrents(a, ts, mt.ParamModeNormal)
	time.Sleep(time.Second * time.Duration(ds.RandInt(defaultDelaySeconds)))
	ts, err = mt.GetFreeAndLargeNSFWTorrents()
	if err != nil {
		log.Printf("failed to fetch nsfw torrents, err: %v", err)
	}
	mt.SaveTorrents(a, ts, mt.ParamModeNSFW)
	log.Printf("nsfw torrents fetched, amount: %d", len(ts))
}

func (a *App) PushMessage() {
	if !pushMutex.TryLock() {
		log.Printf("failed to lock push message, another job may be running.")
		return
	}
	defer pushMutex.Unlock()
	ts := mt.FindNotPushedTorrents(a)
	if len(ts) == 0 {
		log.Printf("nothing to push")
		return
	}
	var count int
	for _, t := range ts {
		time.Sleep(time.Second * time.Duration(ds.RandInt(defaultDelaySeconds)))
		var msg *telebot.Message
		var err error
		switch t.Category {
		case mt.ParamModeNormal:
			msg, err = tg.SendTorrentMessage(t)
		case mt.ParamModeNSFW:
			msg, err = tg.SendNSFWTorrentMessage(t)
		}
		if err != nil {
			log.Printf("failed to push, msg: %v, err: %v", msg, err)
			continue
		}
		if err := mt.UpdateTorrentPushed(a, t.RemoteId, msg.ID); err != nil {
			log.Printf("failed to mark torrent as pushed, you may receive duplicate message, err: %v", err)
		} else {
			log.Printf("torrent pushed, id: %s, name: %s", t.RemoteId, t.Name)
			count++
		}
	}
	log.Printf("push finished, total: %d, success: %d", len(ts), count)
}

func (a *App) DestructMessage() {
	// @TODO
}
