package login

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"luoqiangGin/api"
	"luoqiangGin/internal/app/database"
	"time"
)

type User struct {
	Id       uint   `gorm:"primary_key" json:"id"`
	Account  string `gorm:"type:varchar(50); not null" json:"account"`
	Password string `gorm:"type:varchar(50); not null" json:"password"`
}

func Login(c *gin.Context) {
	// 数据库操作
	db := database.GetDB()
	var user User
	db.Table("user").Find(&user)
	fmt.Println(user)
	// redis 操作
	RedisClient := database.GetRedis()

	c1 := RedisClient.Get()
	defer c1.Close()
	data := make(map[string]interface{})
	data["day"] = 1
	data["time"] = time.Now().Unix()
	//dataJson, _ := json.Marshal(data)
	//r, err := c1.Do("hSet", "signIn", "singIn3", dataJson)
	r, err := c1.Do("hGet", "signIn", "singIn3")
	if err != nil {
		fmt.Println("redis：", err)
	}
	data1 := make(map[string]interface{})
	json.Unmarshal(r.([]byte), &data1)
	fmt.Println(data1["day"])
	c.JSON(200, &api.AAA{
		Name: "zhangsan333 ",
		Age:  0,
	})
}
