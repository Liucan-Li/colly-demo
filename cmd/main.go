package main

import (
	"github.com/Liucan-Li/colly-demo/internal/gorm"
	"github.com/Liucan-Li/colly-demo/internal/spider"
)

func main() {
	gorm.GormDB.Init()

	spider.Spider.DoBatchScrawling()

}
