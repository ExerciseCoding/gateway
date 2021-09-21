package dao

import (
	"errors"
	"fmt"
	"gateway/dto"
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"time"
)

type Admin struct {
	Id int `json:"id" gorm:"primary_key" description:"自增主键"`
	UserName string `json:"user_name" gorm:"column:user_name" description:"管理员用户名"`
	Salt string `json:"user_name" gorm:"column:user_name" description:"盐"`
	Password string `json:"password" gorm:"column:password" description:"密码"`
	UpdatedAt time.Time `json:"updated_at" gorm:"update_at" description:"更新时间"`
	CreateAt time.Time `json:"create_at" gorm:"create_at" description:"创建时间"`
	IsDelete int `json:"is_delete" gorm:"column:is_delete" description:"删除标志"`



}

func (t *Admin) TableName() string{
	return "gateway_admin"
}

func (t *Admin) Find(c *gin.Context, tx *gorm.DB, admin *Admin)(*Admin, error){
	out := &Admin{}
	if err := tx.SetCtx(public.GetTraceContext(c)).Where(admin).Find(out).Error; err != nil{
		return nil,err
	}
	return out, nil
}

func (t *Admin) LoginCheck(c *gin.Context, tx *gorm.DB, param *dto.AdminLoginInput)(*Admin, error){
	adminInfo,err := t.Find(c,tx, &Admin{
			UserName:param.Username,
			IsDelete: 0,
		})
	if err != nil{
		return nil,errors.New("用户信息不存在")
	}
	//param.Password
	//adminInfo.Salt
	fmt.Println(param.Password)
	saltPassword := public.GenSaltPassword(adminInfo.Salt, param.Password)
	fmt.Println(adminInfo.Password, saltPassword)
	if adminInfo.Password != saltPassword {
		return nil, errors.New("密码错误，请重新输入")
	}
	return adminInfo,nil
}