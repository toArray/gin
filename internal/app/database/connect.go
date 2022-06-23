package database

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"luoqiangGin/configs"
	"time"
)

var db *gorm.DB

var RedisClient *redis.Pool

//相关db连接
func InitDb() *gorm.DB {
	// 读取配置文件
	driverName := "mysql"
	host := configs.GetConfig().Mysql.Hostname
	port := configs.GetConfig().Mysql.Port
	database := configs.GetConfig().Mysql.Database
	username := configs.GetConfig().Mysql.Username
	password := configs.GetConfig().Mysql.Password
	charset := configs.GetConfig().Mysql.Charset
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	dbInit, err := gorm.Open(driverName, args)
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}
	//自动创建数据表
	// db.AutoMigrate(&User{})
	return dbInit
}

// SetPool 设置连接池
func SetPool() {
	sqlDB := GetDB().DB()
	// SetMaxIdleConns 设置空闲连接池的最大连接数。
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns 设置到数据库的最大打开连接数。
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime 设置连接可以重用的最长时间。
	sqlDB.SetConnMaxLifetime(time.Hour)

	//sqlDB.Close()
}

// OpenDB 开启数据库
func OpenDB() {
	db = InitDb()
}

// GetDB 连接数据库
func GetDB() *gorm.DB {
	return db
}

// CloseDB 关闭数据库
func CloseDB() {
	db.Close()
}

// redis
func GetRedis() *redis.Pool {

	RedisClient = &redis.Pool{
		//连接方法
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", configs.GetConfig().Redis.Addr)
			if err != nil {
				return nil, err
			}
			// 密码验证
			_, err = c.Do("auth", configs.GetConfig().Redis.Password)
			if err != nil {
				return nil, err
			}
			c.Do("SELECT", 0)
			return c, nil
		},
		//DialContext:     nil,
		//TestOnBorrow:    nil,
		//最大的空闲连接数，表示即使没有redis连接时依然可以保持N个空闲的连接，而不被清除，随时处于待命状态。
		MaxIdle: 1,
		//最大的激活连接数，表示同时最多有N个连接
		MaxActive: 10,
		//最大的空闲连接等待时间，超过此时间后，空闲连接将被关闭
		IdleTimeout: 180 * time.Second,
		//Wait:            false,
		//MaxConnLifetime: 0,

	}
	return RedisClient
}
