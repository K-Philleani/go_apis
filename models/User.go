package models

import (
	"go_apis/conn"
	"log"
)

// user 表
type User struct {
	Id 			int 	`json:"id"`
	UserName 	string 	`json:"user_name"`
	UserPwd 	string 	`json:"user_pwd"`
}

// 获取全部数据
func (u *User) GetUserInfo() (user []User, err error) {
	result := conn.DB.Find(&user)
	err = result.Error
	if err != nil {
		return
	}
	return
}

// 获取单条数据
func (u *User) GetUserRow() (user User, err error) {
	result := conn.DB.Where("id=?", u.Id).Find(&user)
	err = result.Error
	if err != nil {
		return
	}
	return
}

// 更新数据
func (u *User) UpdateRow() (err error){
	result := conn.DB.Model(&User{}).Where("id=?", u.Id).Update("user_pwd", u.UserPwd)
	err = result.Error
	if err != nil {
		return
	}
	log.Println(u)
	return
}

