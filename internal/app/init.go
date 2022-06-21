package app

import (
	"fmt"
	"luoqiangGin/internal/app/router"
)

func Init() {
	//init router
	err := router.InitRouter()
	if err != nil {
		fmt.Println(err)
		return
	}
}
