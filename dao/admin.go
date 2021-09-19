package dao

import "time"

type Admin struct {
	Id int `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName string `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt string `json:"user_name" gorm:"column:user_name" description:"盐"`
	Password string `json:"password" gorm:"column:password" description:"密码"`
	UpdatedAt time.Time `json:"updated_at" gorm:"update_at" description:"更新时间"`
	CreateAt time.Time `json:"create_at" gorm:"create_at" description:"创建时间"`
	IsDelete int `json:"is_delete" gorm:"column:is_delete" description:"删除标志"`



}