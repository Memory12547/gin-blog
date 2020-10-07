/*
 * @Author: your name
 * @Date: 1970-01-01 08:00:00
 * @LastEditTime: 2020-10-07 22:08:23
 * @LastEditors: Matt Meng
 * @Description: In User Settings Edit
 * @FilePath: /go/src/gin-blog/models/models.go
 */
package models

import(
	"log"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gin-blog/pkg/setting"
)

var db *gorm.DB

type Model struct{
	ID int `gorm:"primary_key" json:"id"`
	CreatedOn int `json:"created_on"`
    ModifiedOn int `json:"modified_on"`
}

func init(){
	var (
        err error
        dbType, dbName, user, password, host, tablePrefix string
    )

    sec, err := setting.Cfg.GetSection("database")
    if err != nil {
        log.Fatal(2, "Fail to get section 'database': %v", err)
    }

    dbType = sec.Key("TYPE").String()
    dbName = sec.Key("NAME").String()
    user = sec.Key("USER").String()
    password = sec.Key("PASSWORD").String()
    host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	//建立数据库连接
	db,err = gorm.Open(dbType,fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",user,password,host,dbName))
	if err!=nil{
		log.Println(err)
	}

	//默认表名处理函数，gorm默认的函数是直接返回入参defaultTableName，这里表名添加了前缀tablePrefix（为app.ini中设置的"blog"）
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string)string{
		return tablePrefix+defaultTableName
	}

	//设置默认表名不使用对象名复数
	db.SingularTable(true)
	//设置数据库最大空闲连接为10
	db.DB().SetMaxIdleConns(10)
	//设置数据库最大连接为100
	db.DB().SetMaxOpenConns(100)
}

//关闭数据库连接
func closeDB(){
	defer db.Close()
}