package service

import (
	"github.com/andyron/panel/helper"
	"github.com/andyron/panel/models"
	"github.com/robfig/cron/v3"
	"log"
)

func Cron(exit chan struct{}) {
	c := cron.New(cron.WithSeconds())
	list := make([]*models.Task, 0)
	err := models.DB.Find(&list).Error
	if err != nil {
		log.Fatalln("[DB ERROR] : " + err.Error())
	}
	for _, v := range list {
		_, err := c.AddFunc(v.Spec, func() {
			helper.RunShell(v.ShellPath, v.LogPath)
		})
		if err != nil {
			log.Fatalln("[CRON ERROR] : " + err.Error())
		}
	}
	c.Start()
	defer c.Stop()
	select {
	case <-exit:
		log.Println("[CRON] : exit")
	}
}
