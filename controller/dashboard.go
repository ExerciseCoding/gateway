package controller

import (
	"gateway/dao"
	"gateway/dto"
	"gateway/middleware"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
)

type DashboardController struct {}

func DashboardRegister(group *gin.RouterGroup) {
	service := &DashboardController{}
	group.GET("/panel_group_data", service.PanelGroupData)

}



// PanelGroupData godoc
// @Summary 指标统计
// @Description 指标统计
// @Tags 首页大盘
// @ID /dashboard/panelGroupData
// @Accept json
// @Product json
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /dashboard/panelGroupData [get]
func (dashboard *DashboardController) PanelGroupData(c *gin.Context){
	tx, err := lib.GetGormPool("default")
	if err != nil{
		middleware.ResponseError(c,2001,err)
		return
	}
	serviceInfo := &dao.ServiceInfo{}
	_, serviceNum, err  := serviceInfo.PageList(c,tx,&dto.ServiceListInput{PageNo: 1,PageSize: 1})
	if err != nil{
		middleware.ResponseError(c,2002,err)
		return
	}
	out := &dto.PanelGroupDataOutput{
		ServiceNum: serviceNum,
	}

	middleware.ResponseSuccess(c,out)

}