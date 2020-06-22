package base

import (
	"github.com/SAIKAII/skHappy-IM/infra"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var database *gorm.DB

func Database() *gorm.DB {
	return database
}

type DatabaseStarter struct {
	infra.BaseStarter
}

func (s *DatabaseStarter) Setup(ctx infra.StarterContext) {
	var err error
	database, err = gorm.Open("mysql", "root:123456@(localhost:3306)/happy_im?charset=utf8mb4&parseTime=true&loc=Local")
	if err != nil {
		panic(err)
	}
}
