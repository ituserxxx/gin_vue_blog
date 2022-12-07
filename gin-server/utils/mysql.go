package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/frame/g"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var db *gorm.DB

// ConnectDB connect to database
func ConnectDB() {
	c := g.Config()
	//databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
	databaseURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?loc=Local",
		c.GetString("database.user"),
		c.GetString("database.password"),
		c.GetString("database.host"),
		c.GetString("database.port"),
		c.GetString("database.dbName"),
	)
	fmt.Printf("111111111111111111111111111111%#v", databaseURL)

	var err error
	mysqlConfig := gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.GetString("database.prefix"), // table name prefix, table for `User` would be `t_users`
			SingularTable: true,                           // use singular table name, table for `User` would be `user` with this option enabled
		},
	}
	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN:                       databaseURL, // data source name
		DefaultStringSize:         255,         // default size for string fields
		DisableDatetimePrecision:  true,        // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex:    true,        // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn:   true,        // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false,       // auto configure based on currently MySQL version
	}), &mysqlConfig)
	if err != nil {
		panic(err)
	}
	mysqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	//设置与数据库建立连接的最大数目
	mysqlDB.SetMaxOpenConns(c.GetInt("database.max_open_conns"))
	//设置连接池中的最大闲置连接数
	mysqlDB.SetMaxIdleConns(c.GetInt("database.max_idle_conns"))
	// 设置了连接可复用的最大时间
	mysqlDB.SetConnMaxLifetime(time.Hour)

}

// DB inject with ctx for log
func DB(ctx *gin.Context) *gorm.DB {
	return db.WithContext(ctx)
}

//DisconnectDB disconnect database
func DisconnectDB() {
	mysqlDB, _ := db.DB()
	if err := mysqlDB.Close(); err != nil {
		panic(err)
	}
}
