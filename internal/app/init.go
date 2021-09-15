package app

import (
	"fmt"
	"luoqiangGin/configs"
	"luoqiangGin/internal/app/router"
)

func Init() {
	//init configs
	configs.InitConfig()

	//init Mysql
	err := configs.InitConn()
	if err != nil {
		fmt.Printf("init conn is failed. err:%v", err)
		return
	}

	//init router
	err = router.InitRouter()
	if err != nil {
		fmt.Println(err)
		return
	}
}
