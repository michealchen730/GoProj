package databases

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var Db *gorm.DB

//这里的方法名为init，并不是随意命名的
//init()函数会在每个包完成初始化后自动执行，并且执行优先级比main函数高。
func init() {
	var err error
	dsn := "root:35013501@tcp(192.168.0.2:3307)/test"
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("err:", err.Error())
	}
}
