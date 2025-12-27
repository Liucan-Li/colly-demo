package main

import (
	"github.com/Liucan-Li/colly-demo/internal/dao"
	"github.com/Liucan-Li/colly-demo/internal/dao/model"
	"github.com/Liucan-Li/colly-demo/internal/gorm"
	"github.com/Liucan-Li/colly-demo/internal/spider"
)

func main() {
	gorm.GormDB.Init()

	v := &model.TcgCardOrigin{
		OriginContent: `{"test": "测试1"}`,
	}

	dao.TcgCard.Create(v)
	spider.Spider.DoScrawling()
}
