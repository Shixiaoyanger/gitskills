package controllers

import (
	"email/models"
	"email/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MailPack struct {
	Message    models.Message    `json:"message" binding:"required"`
	Smtpclient models.SmtpClient `json:"smtpclient" binding:"required"`
}

func Regist(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err == nil { //绑定成功
		if len(user.Username) != 0 && len(user.Password) != 0 {
			_, _, err := models.Getinfo(user.Username)
			if err == true {
				c.JSON(http.StatusOK, gin.H{
					"msg":    "用户已存在",
					"status": -1,
				})
			} else {
				models.Insert(user.Username, models.Getkey(user.Password))
				c.JSON(http.StatusOK, gin.H{
					"msg":    "注册成功",
					"status": 1,
				})
			}

		} else {

			c.JSON(http.StatusOK, gin.H{
				"msg":    "请输入用户名和密码",
				"status": 1,
			})
		}
	} else {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"msg":    "服务器获取用户注册信息错误",
			"status": 1,
		})

	}
}
func Login(c *gin.Context) {
	var user models.User
	if c.BindJSON(&user) == nil {
		if flag := models.CheckAuth(user); flag == true {
			token, _ := utils.GenerateToken(c, user.Username, user.Password)
			c.JSON(http.StatusOK, gin.H{
				"msg":    "登陆成功",
				"status": -1,
				"data":   token,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "mima错误，认证失败",
				"status": -1,
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "token解析失败",
			"status": -1,
		})

	}

}

type EmaiInfo struct {
	Addr string `json:"addr" binding:"required"` // 账号/邮箱地址
	Pass string `json:"pass"  binding:"required"`
	Host string `json:"host" binding:"required"` // smtp服务器地址

	Subject  string `json:"subject" binding:"required"`
	Body     string `json:"body" binding:"required"`
	To       string `json:"to" binding:"required"`
	Nickname string `json:"nickname" binding:"required"`
}

func Send(c *gin.Context) {
	var e EmaiInfo

	if err := c.BindJSON(&e); err == nil { //绑定成功
		var msg = models.Message{
			Subject:  e.Subject,
			Body:     e.Body,
			To:       e.To,
			Nickname: e.Nickname,
		}

		client := models.GenerateAuth(e.Addr, e.Pass, e.Host)

		ok := client.AsyncSend(msg, false, func(err error) {
			if err == nil {
			}
		})
		if ok == nil {
			c.JSON(http.StatusOK, gin.H{
				"msg":    "发送成功",
				"status": "1",
			})
		}
	}

}
func Receive(c *gin.Context) {

}
