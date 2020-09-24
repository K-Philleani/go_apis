package models

import (
	"go_apis/conn"
)

type Account struct {
	Id 			 int 		`json:"id" gorm:"primaryKey"`
	UserAccount  string 	`json:"user_account" gorm:"unique"`
	UserPwd 	 string 	`json:"user_pwd"`
	UserPhone 	 string 	`json:"user_phone"`
}

// 新建账号
func (a *Account) CreateAccount() (id int,row int64, err error){
	result := conn.DB.Table("account").Create(&a)
	err = result.Error
	if err != nil {
		return
	}
	id = a.Id
	row = result.RowsAffected
	return
}

// 查询单条账号
func (a *Account) GetAccount() (id int, user Account, err error) {
	result := conn.DB.Table("account").Where("user_account=?", a.UserAccount).Find(&user)
	err = result.Error
	if err != nil {
		return
	}
	id = user.Id
	return
}

