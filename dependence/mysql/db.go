package mysql

import (
	"github.com/GeneralProjectApi/toolkit/settings"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" //这个一定要引入哦！！
)

var DB *gorm.DB

func Setup() {
	dbType := settings.MysqlSettings.Type
	dbUser := settings.MysqlSettings.User
	dbPwd := settings.MysqlSettings.Password
	dbHost := settings.MysqlSettings.Host
	dbName := settings.MysqlSettings.Name
	dns := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPwd, dbHost, dbName)
	db, err := gorm.Open(dbType, dns)
}
