package configs

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

type ConfigData struct {
	SwitchMysql bool //是否需要连接Mysql
	SwitchRedis bool //是否需要连接Redis
	SwitchMongo bool //是否需要连接Mongo

	Mysql MysqlCfgData `toml:"mysql"` //Mysql基础配置
}

var Config *ConfigData

func InitConfig() {
	//decode cfg
	filePath := "./config/config.toml"
	if _, err := toml.DecodeFile(filePath, &Config); err != nil {
		panic(err)
	}

	fmt.Println("decode config success ----------")
}

func InitConn() (err error) {
	config := GetConfig()
	if config.SwitchMysql {
		err = initMysqlConn()
	}

	return
}

func GetConfig() *ConfigData {
	return Config
}

/*
IsNeedConnMysql
@Desc	是否需要连接mysql
*/
func (b *ConfigData) IsNeedConnMysql() bool {
	return b.SwitchMysql
}

/*
IsNeedConnRedis
@Desc	是否需要连接redis
*/
func (b *ConfigData) IsNeedConnRedis() bool {
	return b.SwitchMysql
}

/*
IsNeedConnMongo
@Desc	是否需要连接mongo
*/
func (b *ConfigData) IsNeedConnMongo() bool {
	return b.SwitchMysql
}

/*
GetMysqlCfg
@Desc	获得mysql配置
*/
func (b *ConfigData) GetMysqlCfg() MysqlCfgData {
	return b.Mysql
}
