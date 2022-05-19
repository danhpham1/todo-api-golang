package models

import (
	"fmt"
	beego "github.com/beego/beego/v2/server/web"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dns, err := beego.AppConfig.String("mysqlDNS")
	if err != nil {
		fmt.Println("Có lỗi")
	}
	db, _ = gorm.Open(mysql.Open(dns), &gorm.Config{})
}
