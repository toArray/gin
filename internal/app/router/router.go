package router

import (
	"github.com/gin-gonic/gin"
	"luoqiangGin/internal/app/handler/login"
)

func InitRouter() (err error) {
	router := gin.Default()

	router.Use(middlePrintInfo())
	groupUser := router.Group("/user")
	groupUser.GET("/login3445", login.Login)
	router.GET("/login2", login.Login)

	err = router.Run()
	return
}
