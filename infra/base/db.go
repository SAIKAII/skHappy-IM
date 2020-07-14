package base

import (
	"fmt"
	"github.com/SAIKAII/skHappy-IM/cmd/config"
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
	user := config.GetString("mysql.user")
	pwd := config.GetString("mysql.password")
	host := config.GetString("mysql.host")
	db := config.GetString("mysql.database")
	port := config.GetInt("mysql.port")
	database, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8mb4&parseTime=true&loc=Local", user, pwd, host, port, db))
	if err != nil {
		panic(err)
	}
}
