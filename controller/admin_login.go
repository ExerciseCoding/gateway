package controller

import (
	"encoding/json"
	"gateway/dao"
	"gateway/dto"
	"gateway/middleware"
	"gateway/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"time"
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
	tx, err := lib.GetGormPool("default")
	if err != nil{
		middleware.ResponseError(c,2001, err)
		return
	}

	admin := &dao.Admin{}
	admin,err = admin.LoginCheck(c,tx,params)
	if err != nil{
		middleware.ResponseError(c,2002, err)
		return
	}
	sesInfo := &dto.AdminSessionInfo{
		ID: admin.Id,
		UserName: admin.UserName,
		LoginTime: time.Now(),
	}
	//设置session
	sesBts, err := json.Marshal(sesInfo)
	if err != nil{
		middleware.ResponseError(c, 2003, err)
		return
	}
	ses := sessions.Default(c)
	ses.Set(public.AdminSessionInfoKey,string(sesBts) )
	// session存入redis
	ses.Save()
	out := &dto.AdminLoginOutput{
		Token: admin.UserName,
	}
	middleware.ResponseSuccess(c,out)
}