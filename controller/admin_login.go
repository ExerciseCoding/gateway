package controller

import (
	"gateway/dto"
	"gateway/middleware"
	"github.com/gin-gonic/gin"
)

// AdminLoginController handle login entry
type AdminLoginController struct {}


func AdminLoginRegister(group *gin.RouterGroup){
	adminLogin := &AdminLoginController{}
	group.POST("/login",adminLogin.AdminLogin)
}
// AdminLogin godoc
// @Summary 管理员登录
// @Description 管理员登录
// @Tags 管理员接口
// @ID /admin_login/login
// @Accept json
// @Product json
// @Param body body dto.AdminLoginInput true "body"
// @Success 200 {object} middleware.Response{data=dto.AdminLoginOutput} "success"
// @Router /admin_login/login [post]
func (adminLogin *AdminLoginController) AdminLogin(c *gin.Context){
	params := &dto.AdminLoginInput{}
	if err := params.BindValidParam(c); err != nil{
		middleware.ResponseError(c,1001,err)
		return
	}
	username := params.Username

	out := &dto.AdminLoginOutput{
		Token: params.Username,
	}
	middleware.ResponseSuccess(c,out)
}