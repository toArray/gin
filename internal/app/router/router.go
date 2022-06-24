package router

import (
	"github.com/gin-gonic/gin"
	"luoqiangGin/configs"
	"luoqiangGin/internal/app/handler/algorithm"
	"luoqiangGin/internal/app/handler/login"
	"luoqiangGin/internal/app/handler/user"
)

func InitRouter() (err error) {
	router := gin.Default()
	router.Use(middlePrintInfo())
	groupUser := router.Group("/user")
	groupUser.GET("/login3445", login.Login)
	router.GET("/login2", login.Login)
	router.POST("/sing_in", user.SingIn)
	router.POST("/user_sign", user.UserSign)
	router.POST("/user_sign_list", user.UserSignList)
	router.POST("/sorting", user.Sorting)
	router.POST("/algorithm", algorithm.Index)

	port := ":" + configs.GetConfig().HttpPort
	err = router.Run(port)
	return
}
