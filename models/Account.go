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

// 查询全部账号
func (a *Account) GetAccountAll() (list []Account, err error) {
	result := conn.DB.Table("account").Find(&list)
	err = result.Error
	if err != nil {
		return
	}
	return
}

// 删除账号
func (a *Account) DeleteAccount() (row int64,err error) {
	result := conn.DB.Table("account").Delete(&a, a.Id)
	err = result.Error
	if err != nil {
		return
	}
	row = result.RowsAffected
	return
}
