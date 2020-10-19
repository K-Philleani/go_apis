package apis

import (
	"github.com/gin-gonic/gin"
	"go_apis/models"
	"log"
)

type Request struct {
	Id 			int 	`json:"id"`
	UserName 	string 	`json:"username"`
	Pwd 		string  `json:"pwd"`
}

// 获取全部数据
func GetUserInfo(c *gin.Context) {
	var user models.User
	list, err := user.GetUserInfo()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
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

// 获取单条数据
func GetUserRow(c *gin.Context) {
	var req Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("params failed, err:%v", err)
		return
	}
	user := models.User{
		Id: req.Id,
	}
	info, err := user.GetUserRow()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "查询失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message": "查询成功",
		"user": info,
	})

}

// 更新数据
func UpdateRow(c *gin.Context) {
	var req Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("params failed, err:%v", err)
		return
	}
	user := models.User{
		Id: req.Id,
		UserPwd: req.Pwd,
	}
	err = user.UpdateRow()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "更新失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message": "更新成功",
	})

}

// 删除数据
func DeleteRow(c *gin.Context) {
	var req Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("params failed, err:%v", err)
		return
	}
	user := models.User{
		Id: req.Id,
	}
	err = user.DeleteRow()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message": "删除成功",
	})
}

// 创建数据
func CreateRow(c *gin.Context) {
	var req Request
	err := c.ShouldBindJSON(&req)
	if err != nil {
		log.Printf("params failed, err:%v", err)
		return
	}
	user := models.User{
		UserName: req.UserName,
		UserPwd: req.Pwd,
	}
	n, err := user.CreateRow()
	if err != nil {
		c.JSON(200, gin.H{
			"code": 0,
			"message": "创建失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 1,
		"message": "创建成功",
		"num": n,
	})

}