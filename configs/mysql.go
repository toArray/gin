package configs

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"luoqiangGin/internal/app/constants"
)

//MysqlCfgData mysql配置
type MysqlCfgData struct {
	Address  string //连接地址
	Port     int32  //端口
	User     string //用户名
	Password string //密码
	DB       string //数据库

	MaxOpen int    //最大连接数
	MaxIdle int    //最大空闲数
	Charset string //字符集
}

/*
initMysqlConn
@Desc	初始化mysql连接
*/
func initMysqlConn() error {
	//get cfg
	config := GetConfig()
	mysqlData := config.GetMysqlCfg()
	dataSourceName := getMysqlDataSourceName()

	//sql connect
	db, err := sql.Open(constants.DEFAULT_MYSQL_DRIVER_NAME, dataSourceName)
	if err != nil {
		return err
	}

	//set args
	db.SetMaxOpenConns(mysqlData.MaxOpen)
	db.SetMaxIdleConns(mysqlData.MaxIdle)

	err = db.Ping()
	if err != nil {
		return err
	}

	fmt.Printf("mysql connect sucess. listen port:%v ------------\n", mysqlData.Port)
	return err
}

/*
getMysqlDataSourceName
@Desc 	获得mysql连接信息
*/
func getMysqlDataSourceName() string {
	config := GetConfig()
	mysqlData := config.GetMysqlCfg()
	charset := constants.DEFAULT_MYSQL_CHARSET
	if mysqlData.Charset != "" {
		charset = mysqlData.Charset
	}

	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", mysqlData.User, mysqlData.Password,
		mysqlData.Address, mysqlData.Port, mysqlData.DB, charset)

	return str
}
