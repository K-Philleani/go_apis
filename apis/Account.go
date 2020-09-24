package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_apis/models"
	"log"
)

type Info struct {
	UserAccountAxios 	string 	`json:"userAccount"`
	UserPwdOAxios 		string 	`json:"userPwd"`
	UserPhoneOAxios 	string 	`json:"userPhone"`
}

// 新建账号
func CreateAccount(c *gin.Context) {
	var info Info
	err := c.Bind(&info)
	if err != nil {
		log.Println(err)
		return
	}

	account := models.Account{
		UserAccount: info.UserAccountAxios,
		UserPwd: info.UserPwdOAxios,
		UserPhone: info.UserPhoneOAxios,
	}
	id, _, err := account.GetAccount()
	fmt.Println(id)
	if err != nil {
		log.Println(err)
		return
	}
	if id != 0 {
		c.JSON(200, gin.H{
			"code": -1,
			"message": "账号已存在",
			"id": id,
		})
		return
	}

	id, row, err := account.CreateAccount()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "新建账号失败",
			"err": err,
		})
		panic("新建账户失败， err:")
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message":"新建账号成功",
		"新建账户ID": id,
		"新建数据条数": row,
	})

}

// 查询账号是否存在
func GetAccount(c *gin.Context) {
	userAccount := c.Query("user_account")
	account := models.Account{
		UserAccount: userAccount,
	}
	id, user, err := account.GetAccount()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "查询账号失败",
			"err": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message": "查询账号成功",
		"id": id,
		"info": user,
	})
}

// 登录
func Login(c *gin.Context) {
	var info Info
	err := c.Bind(&info)
	if err != nil {
		log.Println(err)
		return
	}
	account := models.Account{
		UserAccount: info.UserAccountAxios,
		UserPwd: info.UserPwdOAxios,
	}
	id, user, err := account.GetAccount()
	if err != nil {
		log.Println(err)
		return
	}
	if id == 0 {
		c.JSON(200, gin.H{
			"code": -1,
			"message": "账号不存在",
			"id": id,
		})
		return
	}
	if info.UserPwdOAxios != user.UserPwd {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "密码错误",
			"user_account": user.UserAccount,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message": "登录成功",
		"user_account": user.UserAccount,
	})

}
