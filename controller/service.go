package controller

import (
	"errors"
	"fmt"
	"gateway/dao"
	"gateway/dto"
	"gateway/middleware"
	"gateway/public"
	"github.com/e421083458/golang_common/lib"
	"github.com/gin-gonic/gin"
	"strings"
	"time"
)

type ServiceController struct{}

func ServiceRegister(group *gin.RouterGroup) {
	service := &ServiceController{}
	group.GET("/service_list", service.ServiceList)
	group.GET("/service_delete", service.ServiceDelete)
	group.GET("/service_detail", service.ServiceDetail)
	group.POST("/service_add_http",service.ServiceAddHTTP)
	group.POST("/service_update_http", service.ServiceUpdateHTTP)
	group.GET("/service_stat", service.ServiceStat)

	group.POST("/service_add_tcp",)
	group.POST("/seervice_update_tcp",)

	group.POST("/service_add_grpc",)
	group.POST("/service_update_grpc",)

}
// ServiceStat godoc
// @Summary 服务统计
// @Description 服务统计
// @Tags 服务管理
// @ID /service/service_stat
// @Accept json
// @Product json
// @Param id query string true "服务ID"
// @Success 200 {object} middleware.Response{data=dto.ServiceStatOutput} "success"
// @Router /service/service_stat [get]
func (service *ServiceController) ServiceStat(c *gin.Context) {
	params := &dto.ServiceDeleteInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	//tx, err := lib.GetGormPool("default")
	//if err != nil {
	//	middleware.ResponseError(c, 2001, err)
	//	return
	//}
	//// 读取基本信息
	//serviceInfo := &dao.ServiceInfo{ID: params.ID}
	//serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	//serviceDetail, err := serviceInfo.ServiceDetail(c,tx,serviceInfo)
	//if err != nil {
	//	middleware.ResponseError(c, 2002, err)
	//	return
	//}
	todayList := []int64{}
	for i := 0; i < time.Now().Hour(); i++ {
		todayList = append(todayList, 0)
	}
	yesterdayList := []int64{}
	for i := 0; i <= 23 ; i++{
		yesterdayList = append(yesterdayList, 0)
	}
	outItem := dto.ServiceStatOutput{
		Today: todayList,
		Yesterday: yesterdayList,
	}

	middleware.ResponseSuccess(c, outItem)
}

// ServiceList godoc
// @Summary 服务列表
// @Description 服务列表
// @Tags 服务管理
// @ID /service/service_list
// @Accept json
// @Product json
// @Param info query string false "关键词"
// @Param page_size query int truw "每页个数"
// @Param page_no query int true "当前页数"
// @Success 200 {object} middleware.Response{data=dto.ServiceListOutput} "success"
// @Router /service/service_list [get]
func (service *ServiceController) ServiceList(c *gin.Context) {
	params := &dto.ServiceListInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 1001, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	// 从db中分页读取基本信息
	serviceInfo := &dao.ServiceInfo{}
	list, total, err := serviceInfo.PageList(c, tx, params)
	if err != nil {
		fmt.Println("出错了")
		middleware.ResponseError(c, 2002, err)
		return
	}
	// 格式化输出信息
	outList := []dto.ServiceListItemOutput{}
	for _, listItem := range list {
		serviceDetail, err := listItem.ServiceDetail(c, tx, &listItem)
		if err != nil {
			middleware.ResponseError(c, 2003, err)
			return
		}
		serviceAddr := "unkown"

		// 1. http后缀接入 clusterIP+clusterPort+path
		// 2. http域名接入domain
		// 3. tcp或者grpc接入 clusterIP+servicePort
		clusterIP := lib.GetStringConf("base.cluster.cluster_ip")
		clusterPort := lib.GetStringConf("base.cluster.cluster_port")
		clusterSSLPort := lib.GetStringConf("base.cluster.cluster_ssl_port")
		// HTTPS
		if serviceDetail.Info.LoadType == public.LoadTypeHTTP && serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL &&
			serviceDetail.HTTPRule.NeedHttps == 1 {
			serviceAddr = fmt.Sprintf("%s:%s%s", clusterIP, clusterSSLPort, serviceDetail.HTTPRule.Rule)
		}
		// HTTP
		if serviceDetail.Info.LoadType == public.LoadTypeHTTP &&
			serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypePrefixURL && serviceDetail.HTTPRule.NeedHttps == 0 {
			serviceAddr = fmt.Sprintf("%s:%s%s", clusterIP, clusterPort, serviceDetail.HTTPRule.Rule)
		}
		// Domain
		if serviceDetail.Info.LoadType == public.LoadTypeHTTP && serviceDetail.HTTPRule.RuleType == public.HTTPRuleTypeDomain {
			serviceAddr = serviceDetail.HTTPRule.Rule
		}
		// TCP
		if serviceDetail.Info.LoadType == public.LoadTypeTCP {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.TCPRule.Port)
		}
		// GRPC
		if serviceDetail.Info.LoadType == public.LoadTypeGRPC {
			serviceAddr = fmt.Sprintf("%s:%d", clusterIP, serviceDetail.GRPCRule.Port)
		}
		ipList := serviceDetail.LoadBalance.GetIPListByModel()
		outItem := dto.ServiceListItemOutput{
			ID:          listItem.ID,
			ServiceName: listItem.ServiceName,
			ServiceDesc: listItem.ServiceDesc,
			ServiceAddr: serviceAddr,
			Qps:         0,
			Qpd:         0,
			TotalNode:   len(ipList),
		}
		outList = append(outList, outItem)
	}
	out := &dto.ServiceListOutput{
		Total: total,
		List:  outList,
	}
	middleware.ResponseSuccess(c, out)
}


// ServiceDelete godoc
// @Summary 服务删除
// @Description 服务删除
// @Tags 服务管理
// @ID /service/service_delete
// @Accept json
// @Product json
// @Param id query string true "服务ID"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_delete [get]
func (service *ServiceController) ServiceDelete(c *gin.Context) {
	params := &dto.ServiceDeleteInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	// 读取基本信息
	serviceInfo := &dao.ServiceInfo{ID: params.ID}
	serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	serviceInfo.IsDelete = 1
	if err := serviceInfo.Save(c,tx); err != nil{
		middleware.ResponseError(c,2003,err)
		return
	}
	middleware.ResponseSuccess(c, "")
}


// ServiceAddHTTP godoc
// @Summary 添加HTTP服务
// @Description 添加HTTP服务
// @Tags 服务管理
// @ID /service/service_add_http
// @Accept json
// @Product json
// @Param body body dto.ServiceAddHTTPInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_add_http [post]
func (service *ServiceController) ServiceAddHTTP(c *gin.Context) {
	params := &dto.ServiceAddHTTPInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	tx,err := lib.GetGormPool("default")
	if err != nil{
		middleware.ResponseError(c,2001,err)
		return
	}
	// 开启事务
	tx.Begin()
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}

	if _,err = serviceInfo.Find(c,tx,serviceInfo);err == nil{
		tx.Rollback()
		middleware.ResponseError(c,2002, errors.New("服务已存在"))
		return
	}
	// 校验路径
	httpUrl := &dao.HttpRule{RuleType: params.RuleType, Rule: params.Rule}
	if _ , err = httpUrl.Find(c,tx,httpUrl); err == nil{
		tx.Rollback()
		middleware.ResponseError(c,2003,errors.New("域名接入前缀或者域名已存在"))
		return
	}
	// 校验IP和权重
	 if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")){
	 	tx.Rollback()
	 	middleware.ResponseError(c, 2004, errors.New("IP列表和权重列表不一致"))
	 	return
	 }
	 serviceModel := &dao.ServiceInfo{
	 	ServiceName: params.ServiceName,
	 	ServiceDesc: params.ServiceDesc,
	 }

	if err = serviceModel.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2005, err)
		return
	}
	httpRule := &dao.HttpRule{
		ServiceID: serviceModel.ID,
		RuleType: params.RuleType,
		Rule: params.Rule,
		NeedHttps: params.NeedHttps,
		NeedWebsocket: params.NeedWebsocket,
		UrlRewrite: params.UrlRewrite,
		HeaderTransfor: params.HeaderTransfor,
	}
	if err = httpRule.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2006,err)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID: serviceModel.ID,
		OpenAuth: params.OpenAuth,
		BlockList: params.BlackList ,
		WhiteList: params.WhiteList ,
		ClientIPFlowLimit: params.ClientipFlowLimit,
		ServiceFlowLimit: params.ServiceFlowLimit,

	}
	if err = accessControl.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2007,err)
		return
	}

	loadBalance := &dao.LoadBalance{
		ServiceID: serviceModel.ID,
		RoundType: params.RoundType,
		IpList: params.IpList,
		WeightList: params.WeightList,
		UpstreamConnectTimeout: params.UpstreamConnectTimeout,
		UpstreamHeaderTimeout: params.UpstreamHeaderTimeout,
		UpstreamIdleTimeout: params.UpstreamIdleTimeout,
		UpstreamMaxIdle: params.UpstreamMaxIdle,
	}
	if err = loadBalance.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2008,err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")
}

// ServiceUpdateHTTP godoc
// @Summary 修改HTTP服务
// @Description 修改HTTP服务
// @Tags 服务管理
// @ID /service/service_update_http
// @Accept json
// @Product json
// @Param body body dto.ServiceUpdateHTTPInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_update_http [post]
func (service *ServiceController) ServiceUpdateHTTP(c *gin.Context) {
	params := &dto.ServiceUpdateHTTPInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}
	// 校验IP和权重
	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")){
		middleware.ResponseError(c, 2001, errors.New("IP列表和权重列表不一致"))
		return
	}
	tx,err := lib.GetGormPool("default")
	if err != nil{
		middleware.ResponseError(c,2002,err)
		return
	}

	//// 校验路径
	//httpUrl := &dao.HttpRule{RuleType: params.RuleType, Rule: params.Rule}
	//if _ , err = httpUrl.Find(c,tx,httpUrl); err == nil{
	//	tx.Rollback()
	//	middleware.ResponseError(c,2003,errors.New("域名接入前缀或者域名已存在"))
	//	return
	//}

	// 开启事务
	tx.Begin()
	serviceInfo := &dao.ServiceInfo{ServiceName: params.ServiceName}
	serviceInfo, err = serviceInfo.Find(c,tx,serviceInfo)
	if err != nil{
		tx.Rollback()
		middleware.ResponseError(c, 2003, errors.New("服务不存在"))
	}
	serviceDetail,err := serviceInfo.ServiceDetail(c,tx,serviceInfo)
	if err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2004, errors.New("服务不存在"))
		return
	}
	info := serviceDetail.Info
	info.ServiceDesc = params.ServiceDesc
	if err := info.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2005,err)
		return
	}

	// HTTP服务修改
	httpRule := serviceDetail.HTTPRule
	httpRule.NeedHttps = params.NeedHttps
	httpRule.NeedWebsocket = params.NeedWebsocket
	httpRule.UrlRewrite = params.UrlRewrite
	httpRule.HeaderTransfor = params.HeaderTransfor
	if err = httpRule.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2005,err)
		return
	}

	accessControl := serviceDetail.AccessControl
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlockList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.ClientIPFlowLimit = params.ClientipFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err = accessControl.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2006,err)
		return
	}
	loadBalance := serviceDetail.LoadBalance
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.UpstreamConnectTimeout = params.UpstreamConnectTimeout
	loadBalance.UpstreamIdleTimeout = params.UpstreamIdleTimeout
	loadBalance.UpstreamMaxIdle = params.UpstreamMaxIdle
	if err = loadBalance.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2008,err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")
}



// ServiceDetail godoc
// @Summary 服务详情
// @Description 服务详情
// @Tags 服务管理
// @ID /service/service_detail
// @Accept json
// @Product json
// @Param id query string true "服务ID"
// @Success 200 {object} middleware.Response{data=dao.ServiceDetail} "success"
// @Router /service/service_detail [get]
func (service *ServiceController) ServiceDetail(c *gin.Context) {
	params := &dto.ServiceDeleteInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2000, err)
		return
	}

	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	// 读取基本信息
	serviceInfo := &dao.ServiceInfo{ID: params.ID}
	serviceInfo, err = serviceInfo.Find(c, tx, serviceInfo)
	serviceDetail, err := serviceInfo.ServiceDetail(c,tx,serviceInfo)
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return
	}
	middleware.ResponseSuccess(c, serviceDetail)
}




// ServiceAddTCP godoc
// @Summary 添加TCP服务
// @Description 添加TCP服务
// @Tags 服务管理
// @ID /service/service_add_tcp
// @Accept json
// @Product json
// @Param body body dto.ServiceAddTCPInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_add_tcp [post]
func (service *ServiceController) ServiceAddTCP(c *gin.Context) {
	params := &dto.ServiceAddTCPInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	tx,err := lib.GetGormPool("default")
	if err != nil{
		middleware.ResponseError(c,2001,err)
		return
	}
	serviceInfo := &dao.ServiceInfo{
		ServiceName: params.ServiceName,
		IsDelete: 0,
	}
	if _,err = serviceInfo.Find(c,tx,serviceInfo);err == nil{
		middleware.ResponseError(c,2002, errors.New("服务名已存在，请重新输入"))
		return
	}
	tcpRule := &dao.TcpRule{
		Port: params.Port,
	}
	if _, err  = tcpRule.Find(c,tx,tcpRule); err == nil{
		middleware.ResponseError(c,2003,errors.New("服务端口被占用，请重新输入"))
		return
	}
	grpcRule := &dao.GrpcRule{
		Port: params.Port,
	}
	if _, err := grpcRule.Find(c,tx,grpcRule); err == nil{
		middleware.ResponseError(c,2004, errors.New("服务端口被占用，请重新输入"))
		return
	}
	// 校验IP和权重数量一致
	if len(strings.Split(params.IpList,",")) != len(strings.Split(params.WeightList, "")){
		middleware.ResponseError(c,2005, errors.New("IP列表与权重设置不匹配"))
		return
	}
	// 开启事务
	tx.Begin()
	serviceModel := &dao.ServiceInfo{
		LoadType: public.LoadTypeTCP,
		ServiceName: params.ServiceName,
		ServiceDesc: params.ServiceDesc,
	}

	if err = serviceModel.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2006, err)
		return
	}
	newTcpTule := &dao.TcpRule{
		ServiceID: serviceInfo.ID,
		Port: params.Port,
	}
	if err = newTcpTule.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2007,err)
		return
	}
	loadBalance := &dao.LoadBalance{
		ServiceID: serviceInfo.ID,
		RoundType: params.RoundType,
		IpList: params.IpList,
		WeightList: params.WeightList,
		ForbidList: params.ForbidList,
	}
	if err = loadBalance.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2008,err)
		return
	}

	accessControl := &dao.AccessControl{
		ServiceID: serviceModel.ID,
		OpenAuth: params.OpenAuth,
		BlockList: params.BlackList ,
		WhiteList: params.WhiteList ,
		ClientIPFlowLimit: params.ClientIPFlowLimit,
		ServiceFlowLimit: params.ServiceFlowLimit,
	}
	if err = accessControl.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2009,err)
		return
	}

	tx.Commit()
	middleware.ResponseSuccess(c, "")
}

// ServiceUpdateTCP godoc
// @Summary 修改TCP服务
// @Description 修改TCP服务
// @Tags 服务管理
// @ID /service/service_update_tcp
// @Accept json
// @Product json
// @Param body body dto.ServiceUpdateTCPInput true "body"
// @Success 200 {object} middleware.Response{data=string} "success"
// @Router /service/service_update_tcp [post]
func (service *ServiceController) ServiceUpdateTCP(c *gin.Context) {
	params := &dto.ServiceUpdateTCPInput{}
	if err := params.BindValidParam(c); err != nil {
		middleware.ResponseError(c, 2001, err)
		return
	}
	// 校验IP和权重
	if len(strings.Split(params.IpList, "\n")) != len(strings.Split(params.WeightList, "\n")){
		middleware.ResponseError(c, 2002, errors.New("IP列表和权重列表不一致"))
		return
	}
	tx,err := lib.GetGormPool("default")
	if err != nil{
		middleware.ResponseError(c,2003,err)
		return
	}

	// 开启事务
	tx.Begin()
	serviceInfo := &dao.ServiceInfo{
		ID: params.ID,
	}
	serviceDetail,err := serviceInfo.ServiceDetail(c,tx, serviceInfo)
	if err != nil{
		middleware.ResponseError(c,2004, err)
		return
	}
	info := serviceDetail.Info
	info.ServiceDesc = params.ServiceDesc

	if err = info.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2005,err)
		return
	}

	newTcpTule := &dao.TcpRule{}
	if serviceDetail.TCPRule != nil{
		newTcpTule = serviceDetail.TCPRule
	}
	newTcpTule.ServiceID = info.ID
	newTcpTule.Port = params.Port
	if err = newTcpTule.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2007,err)
		return
	}

	loadBalance := &dao.LoadBalance{}
	if serviceDetail.LoadBalance != nil{
		loadBalance = serviceDetail.LoadBalance
	}
	loadBalance.ServiceID = info.ID
	loadBalance.RoundType = params.RoundType
	loadBalance.IpList = params.IpList
	loadBalance.WeightList = params.WeightList
	loadBalance.ForbidList = params.ForbidList
	if err = loadBalance.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2008,err)
		return
	}

	accessControl := &dao.AccessControl{}
	if accessControl != nil{
		accessControl = serviceDetail.AccessControl
	}
	accessControl.ServiceID = info.ID
	accessControl.OpenAuth = params.OpenAuth
	accessControl.BlockList = params.BlackList
	accessControl.WhiteList = params.WhiteList
	accessControl.ClientIPFlowLimit = params.ClientIPFlowLimit
	accessControl.ServiceFlowLimit = params.ServiceFlowLimit
	if err = accessControl.Save(c,tx); err != nil{
		tx.Rollback()
		middleware.ResponseError(c,2009,err)
		return
	}
	tx.Commit()
	middleware.ResponseSuccess(c, "")
}