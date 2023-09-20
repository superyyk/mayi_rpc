package db

import (
	"fmt"
	"log"
	"github.com/superyyk/mayi_rpc/model"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var cn = model.ConnInfo{
	"yunboss",  //本地 root gaoding
	"yyk*2012", //本地 yyk*2012  yyk2012
	"127.0.0.1",
	"yunboss", //kuanku
	3306,
}
var Db *gorm.DB

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err.Error())
	}
}

// 数据库连接池
func init() {
	Db = DbConn(cn.Username, cn.Password, cn.Host, cn.Db, cn.Port)
}

func NewDb() {

}

func DbConn(Username, Password, Host, Db string, Port int) *gorm.DB {

	connArgs := fmt.Sprintf("%s:%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", Username, Password, Host, Port, Db)
	db, err := gorm.Open("mysql", connArgs)
	CheckErr(err)

	db.SingularTable(true)
	//defer db.Close()
	//db.DB().SetMaxIdleConns(1000) //设置数据库连接池最大连接数
	//db.DB().SetMaxIdleConns(1000) //连接池最大允许的空闲连接数，如果没有sql任务需要执行的连接数大于1，超过的连接会被连接池关闭。
	return db
}
