package dto

import (
	"gateway/public"
	"github.com/gin-gonic/gin"
)

type AdminLoginInput struct {
	// example 默认值
	// Comment	数据库备注 对应数据库中的列Comment值
	// json: 输出时转为json form:转换为结构体
	Username string `json:"username" form:"username" comment:"姓名" example:"admin" validate:"required"` // 管理员用户名
	Password string `json:"password" form:"password" comment:"密码" example:"123456" validate:"required"` // 密码
}

// BindValidParam 校验用户信息，绑定到AdminLoginInput
//

func (param *AdminLoginInput) BindValidParam(c *gin.Context) error{
	return public.DefaultGetValidParams(c, param)
}

type AdminLoginOutput struct {
	Token string `json:"token" form:"token" comment:"token" example:"token" validate:""` //token

}