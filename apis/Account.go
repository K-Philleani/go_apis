package apis

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_apis/JWT"
	"go_apis/models"
	"log"
	"net/http"
	"strconv"
)

type Info struct {
	UserIdAxios    		int 	`json:"userId"`
	UserAccountAxios 	string 	`json:"userAccount"`
	UserPwdOAxios 		string 	`json:"userPwd"`
	UserPhoneOAxios 	string 	`json:"userPhone"`
}

// 新建账号
func CreateAccount(c *gin.Context) {
	userAccount := c.PostForm("userAccount")
	userPwd := c.PostForm("userPwd")
	userPhone := c.PostForm("userPhone")
	account := models.Account{
		UserAccount: userAccount,
		UserPwd: userPwd,
		UserPhone: userPhone,
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
	userAccount := c.PostForm("userAccount")
	userPwd := c.PostForm("userPwd")
	account := models.Account{
		UserAccount: userAccount,
		UserPwd: userPwd,
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
	if userPwd != user.UserPwd {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "密码错误",
			"user_account": user.UserAccount,
		})
		return
	}
	token, err := JWT.ReleaseToken(account)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": -1,
			"message": "系统异常",
		})
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message": "登录成功",
		"user_account": user.UserAccount,
		"token": token,

	})

}

// 查询全部账号
func GetAccountAll(c *gin.Context) {
	var account models.Account
	list, err := account.GetAccountAll()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 1,
			"message": "查询失败",
			"err": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message": "查询成功",
		"userList": list,
	})
}

// 删除账号
func DeleteAccount(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println("获取参数失败")
		return
	}

	account := models.Account{
		Id: id,
	}
	row, err := account.DeleteAccount()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "删除失败",
			"err": err,
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message": "删除成功",
		"row": row,
	})
}

// 更新账号数据
func UpdateAccount(c *gin.Context) {
	userId, err := strconv.Atoi(c.PostForm("userId"))
	if err != nil {
		log.Println(err)
		return
	}
	userPwd := c.PostForm("userPwd")
	userPhone := c.PostForm("userPhone")
	account := models.Account{
		Id: userId,
		UserPwd: userPwd,
		UserPhone: userPhone,
	}
	row, err := account.UpdateAccount()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "更新失败",
			"err": err,
		})
		return
	}

	c.JSON(200, gin.H{
		"code": 1,
		"message": "更新成功",
		"row": row,
	})
}
