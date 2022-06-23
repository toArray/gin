package user

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"luoqiangGin/api"
	"luoqiangGin/internal/app/database"
	"time"
)

type UserSingIn struct {
	Day     int    `json:"day"`
	OldTime string `json:"old_time"`
	NewTime string `json:"new_time"`
}

// SingIn 连续签到七天 （不管周）
func SingIn(c *gin.Context) {
	userId := c.Query("user_id")
	key := "SingIn" + userId
	var prompt string
	var userSingIn UserSingIn
	// redis 操作
	RedisClient := database.GetRedis()
	c1 := RedisClient.Get()
	defer c1.Close()
	r, err := c1.Do("hGet", "signIn", key)
	if err != nil {
		fmt.Println("redis：", err)
	}
	// 获取当前时间
	currentTime := time.Now()
	oldTime := currentTime.AddDate(0, 0, -1).Format("2006-01-02")
	newTime := currentTime.Format("2006-01-02")
	// 判断redis 里是否有这个用户的值
	if r != nil {
		json.Unmarshal(r.([]byte), &userSingIn)
		// 判断是否累计完成7天签到
		if userSingIn.Day == 7 {
			prompt = "你已累计完成七天任务签到"
		} else {
			// 判断今天是否已经签到
			if userSingIn.NewTime == newTime {
				prompt = "你已签到请明天再来！"
			} else {
				// 判断前一天是否签到
				if userSingIn.NewTime == oldTime {
					// 加天数
					userSingIn.Day += 1
					prompt = SingInData(userSingIn, oldTime, newTime, c1, key, prompt, userSingIn.Day)
				} else {
					prompt = SingInData(userSingIn, oldTime, newTime, c1, key, prompt, 1)
				}
			}
		}
	} else {
		// 如果没有则直接添加
		prompt = SingInData(userSingIn, oldTime, newTime, c1, key, prompt, 1)
	}

	c.JSON(200, &api.ReturnData{
		Code: c.Writer.Status(),
		Msg:  prompt,
	})
}

func SingInData(userSingIn UserSingIn, oldTime string, newTime string, c1 redis.Conn, key string, prompt string, day int) string {
	userSingIn.Day = day
	userSingIn.OldTime = oldTime
	userSingIn.NewTime = newTime
	dataJson, _ := json.Marshal(userSingIn)
	_, redisErr := c1.Do("hSet", "signIn", key, dataJson)
	if redisErr != nil {
		fmt.Println("redis：", redisErr)
		prompt = "签到失败"
	} else {
		if userSingIn.Day == 7 {
			prompt = fmt.Sprintf("签到成功！恭喜您累计完成七天签到")
		} else if userSingIn.Day == 3 {
			prompt = fmt.Sprintf("签到成功！你已累计签到%v天,3日礼包已发送你的账户请查收", userSingIn.Day)
		} else {
			prompt = fmt.Sprintf("签到成功！你已累计签到%v天", userSingIn.Day)
		}
	}
	return prompt
}

func UserSign(c *gin.Context) {
	//userId := c.Query("user_id")
	//key := fmt.Sprintf("SingIn%v", userId)
	//// 获取今天是周几
	//t := time.Now()
	//whatDay := int(t.Weekday())
	//var prompt string
	//var userSingIn UserSingIn
	//// redis 操作
	//RedisClient := database.GetRedis()
	//c1 := RedisClient.Get()
	//defer c1.Close()
	//r, err := c1.Do("hGet", "signIn", key)
	//if err != nil {
	//	fmt.Println("redis：", err)
	//}
	//if r != nil {
	//
	//} else {
	//
	//}

}
