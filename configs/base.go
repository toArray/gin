package configs

import (
	"fmt"
	"github.com/spf13/viper"
)

//解析配置

type TomlConfig struct {
	Debug    bool   `json:"debug"`
	HttpPort string `json:"httpport"`
	Mysql    Mysql1
	Redis    Redis
}
type Mysql1 struct {
	Hostname string `json:"hostname"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Database string `json:"database"`
	Password string `json:"password"`
	Charset  string `json:"charset"`
}

type Redis struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
}

var C TomlConfig

func Init() {
	viper.SetConfigName("config")   // 文件名
	viper.SetConfigType("toml")     // 文件类型
	viper.AddConfigPath("./config") // 路径
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("read config failed: %v", err)
	}
	//fmt.Println(viper.Get("app_name"))
	viper.Unmarshal(&C)
	viper.WatchConfig()
}

func GetConfig() TomlConfig {
	return C
}
