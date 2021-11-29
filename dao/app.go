package dao

import (
	"gateway/dto"
	"gateway/public"
	"github.com/e421083458/gorm"
	"github.com/gin-gonic/gin"
	"time"
)

type App struct {
	ID        int64     `json:"id" gorm:"primary_key"`
	AppID     string    `json:"app_id" gorm:"column:app_id" description:"租户id	"`
	Name      string    `json:"name" gorm:"column:name" description:"租户名称	"`
	Secret    string    `json:"secret" gorm:"column:secret" description:"密钥"`
	WhiteIPS  string    `json:"white_ips" gorm:"column:white_ips" description:"ip白名单，支持前缀匹配"`
	Qpd       int64     `json:"qpd" gorm:"column:qpd" description:"日请求量限制"`
	Qps       int64     `json:"qps" gorm:"column:qps" description:"每秒请求量限制"`
	CreatedAt time.Time `json:"create_at" gorm:"column:create_at" description:"添加时间	"`
	UpdatedAt time.Time `json:"update_at" gorm:"column:update_at" description:"更新时间"`
	IsDelete  int8      `json:"is_delete" gorm:"column:is_delete" description:"是否已删除；0：否；1：是"`
}

func(app *App) TableName()string{
	return "gateway_app"
}

func(app *App) Find(c *gin.Context, tx *gorm.DB, search *App)(*App, error){
	appModel := &App{}
	err := tx.SetCtx(public.GetTraceContext(c)).Where(search).Find(appModel).Error
	return appModel, err
}

func(app *App) Save(c *gin.Context, tx *gorm.DB) error {
	if err := tx.SetCtx(public.GetTraceContext(c)).Save(app).Error; err != nil{
		return err
	}
	return nil
}

func(app *App) AppList(c *gin.Context, tx *gorm.DB, params *dto.APPListInput)([]App , int64, error){
	var appList []App
	var count int64
	pageNo := params.PageNo
	pageSize := params.PageSize

	offset := (pageNo - 1) * pageSize
	query := tx.SetCtx(public.GetTraceContext(c))
	query = query.Table(app.TableName()).Select("*")
	query = query.Where("is_delete=?", 0)
	if params.Info != ""{
		query = query.Where("(name like ? or app_id like ?)", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	err := query.Limit(pageSize).Offset(offset).Order("id desc").Find(&appList).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, 0, err
	}

	errCount := query.Count(&count).Error
	if errCount != nil{
		return nil, 0, err
	}

	return appList, count, nil
}