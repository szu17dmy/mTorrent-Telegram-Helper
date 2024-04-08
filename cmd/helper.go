package main

import (
	"github.com/szu17dmy/mtorrent-telegram-helper/internal/app"

	"github.com/robfig/cron/v3"
)

func main() {
	a := app.NewApp()
	c := cron.New()
	a.SetupCronJob(c)
	c.Start()
	select {}
}
