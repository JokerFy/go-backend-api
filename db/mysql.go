package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
	"gobackend/conf"
)

var Eloquent *gorm.DB

func MysqlDial(config *conf.MysqlConf) error {
	var err error
	Eloquent, err = gorm.Open("mysql", config.DSN)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if Eloquent.Error != nil {
		fmt.Printf("database error %v", Eloquent.Error)
	}
	return nil
}
