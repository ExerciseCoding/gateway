package controller

import (
	"encoding/json"
	"fmt"
	"gateway/dto"
	"gateway/middleware"
	"gateway/public"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AdminController struct {}

func AdminRegister(group *gin.RouterGroup){
	admin := &AdminController{}
	group.GET("/admin_info", admin.AdminInfo)
}

// AdminInfo godoc
// @Summary 管理员信息
// @Description 管理员信息
// @Tags 管理员接口
// @ID /admin/admin_info
// @Accept json
// @Product json
// @Success 200 {object} middleware.Response{data=dto.AdminInfoOutput} "success"
// @Router /admin/admin_info [get]
func (admin *AdminController) AdminInfo(c *gin.Context){
	sess := sessions.Default(c)
	sessInfo := sess.Get(public.AdminSessionInfoKey)
	adminSessionInfo := &dto.AdminSessionInfo{}
	if err := json.Unmarshal([]byte(fmt.Sprint(sessInfo)),adminSessionInfo); err != nil{
		middleware.ResponseError(c,2000, err)
		return
	}

	out := &dto.AdminInfoOutput{
		ID: adminSessionInfo.ID,
		Name: adminSessionInfo.UserName,
		LoginTime: adminSessionInfo.LoginTime,
		Avatar: "https://vip.fxxkpython.com/wp-content/uploads/2020/02/cropped-python6-1581516763-1.jpeg",
		Introduction: "I am a super administrator",
		Roles: []string{"admin"},
	}
	middleware.ResponseSuccess(c,out)
}