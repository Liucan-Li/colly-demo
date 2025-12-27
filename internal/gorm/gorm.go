package gorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormDb struct{}

var GormDB GormDb

var db *gorm.DB

func (d *GormDb) Init() {
	dsn := "liliucan:Llc_2026@tcp(rm-bp1yznltgl11382i05o.mysql.rds.aliyuncs.com)/tcg-card?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("数据库连接失败")
	}

}

func (d *GormDb) GetClient() *gorm.DB {
	return db
}
