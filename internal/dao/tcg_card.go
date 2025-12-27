package dao

import (
	"github.com/Liucan-Li/colly-demo/internal/dao/model"
	"github.com/Liucan-Li/colly-demo/internal/gorm"
)

type TcgCardDao struct{}

var TcgCard = &TcgCardDao{}

func (t *TcgCardDao) Create(v *model.TcgCardOrigin) {
	db := gorm.GormDB.GetClient()
	db.Create(v)
}
