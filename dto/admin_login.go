package dto

import (
	"gateway/public"
	"github.com/gin-gonic/gin"
)

type AdminLoginInput struct {
	// example 默认值
	// Comment	数据库备注 对应数据库中的列Comment值
	Username string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required"` // 用户名, json: 输出时转为json form:转换为结构体
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"`
}

// BindValidParam 校验用户信息，绑定到AdminLoginInput
//

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error{
	return public.DefaultGetValidParams(c, param)
}