package mysql

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"go-microservice/common/basic/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
	"time"
)

func initMysql() {

	var err error

	logx.Info("初始化Mysql...")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.GetMysqlConfig().GetUser(),
		config.GetMysqlConfig().GetPassword(),
		config.GetMysqlConfig().GetURL(),
		config.GetMysqlConfig().GetPort(),
		config.GetMysqlConfig().GetDbname(),
	)

	// 创建连接
	mysqlDB, err = gorm.Open(
		mysql.New(
			mysql.Config{
				DSN:                       dsn,   // DSN data source name
				DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
				SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
			}),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				//TablePrefix:   "",   // 表名前缀
				SingularTable: true, // 使用单数表名，启用该选项后，`User` 表将是`user`
			},
		})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	dbConn, _ := mysqlDB.DB()

	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	dbConn.SetMaxOpenConns(config.GetMysqlConfig().GetMaxOpenConnection())

	//SetMaxIdleConns 设置空闲连接池中连接的最大数量
	dbConn.SetMaxIdleConns(config.GetMysqlConfig().GetMaxIdleConnection())

	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	dbConn.SetConnMaxLifetime(time.Hour)

	logx.Info("初始化Mysql，检测连接...")
	// 激活链接
	if err = dbConn.Ping(); err != nil {
		log.Fatal(err)
	}
	logx.Info("初始化Mysql，检测连接Ping. YES")
}
