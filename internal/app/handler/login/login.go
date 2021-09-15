package login

import (
	"github.com/gin-gonic/gin"
	"luoqiangGin/api"
)

func Login(c *gin.Context) {
	c.JSON(200, &api.AAA{
		Name: "zhangsan333 ",
		Age:  0,
	})
}
