package v1

import (
	"gin-demo/global"
	"gin-demo/model"
	"gin-demo/model/request"
	"gin-demo/model/response"
	"gin-demo/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// CreateMetas ...
// @Tags CreateMetas
// @Summary 创建 meta
// @accept application/json
// @Produce application/json
// @Param data body model.StatisticsMeta true "标题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/statistics/meta [post]
func CreateMetas(c *gin.Context) {
	var m request.CreateMeta
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"msg": err.Error(),
			},
		)
	}

	meta := model.StatisticsMeta{Title: m.Title}
	err := global.DB.Create(&meta).Error
	if err != nil {
		global.LOG.Error("创建 meta 失败", zap.Any("err", err))
		response.FailWithDetailed(response.StatisticsMeta{Meta: meta}, "创建 meta 失败", c)
		return
	}
	response.OkWithDetailed(response.StatisticsMeta{Meta: meta}, "创建 meta 成功", c)

}

// GetMetas ...
// @Tags GetMetas
// @Summary 获取 meta 列表
// @accept application/json
// @Produce application/json
// @Param data body model.StatisticsMeta true "标题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/statistics/meta [get]
func GetMetas(c *gin.Context) {
	var m request.GetMeta
	if err := c.ShouldBindQuery(&m); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	metas := []model.StatisticsMeta{}
	err := global.DB.Where("title LIKE ?", "%"+m.Title+"%").Find(&metas).Error

	if err != nil {
		global.LOG.Error("获取 meta 失败", zap.Any("err", err))
		response.FailWithDetailed(response.StatisticsMetaList{Meta: metas}, "获取 meta 失败", c)
		return
	}
	response.OkWithDetailed(response.StatisticsMetaList{Meta: metas}, "获取 meta 成功", c)

}

// CreateStatisticsData ...
// @Tags CreateStatisticsData
// @Summary 创建统计数据
// @accept application/json
// @Produce application/json
// @Param data body model.StatisticsData true "标题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"创建成功"}"
// @Router /v1/statistics/data [post]
func CreateStatisticsData(c *gin.Context) {
	var m request.CreateData
	if err := c.ShouldBindJSON(&m); err != nil {
		c.JSON(
			http.StatusOK,
			gin.H{
				"msg": err.Error(),
			},
		)
	}

	data := model.StatisticsData{MetaID: m.MetaID, Dt: m.Dt, Val: m.Val}
	err := global.DB.Create(&data).Error
	if err != nil {
		global.LOG.Error("创建 statistics data失败", zap.Any("err", err))
		response.FailWithDetailed(data, "创建 statistics data 失败", c)
		return
	}
	response.OkWithDetailed(data, "创建 statistics data 成功", c)

}

// GetIntervalData ...
// @Tags GetIntervalData
// @Summary 获取某时间段内的统计数据和
// @accept application/json
// @Produce application/json
// @Param data body model.StatisticsMeta true "标题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/statistics/interval [get]
func GetIntervalData(c *gin.Context) {
	var m request.GetIntervalData
	if err := c.ShouldBindQuery(&m); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data := service.GetIntervalData(&m)
	response.OkWithDetailed(data, "获取 interval data 成功", c)

}

// GetIntervalUnitData ...
// @Tags GetIntervalUnitData
// @Summary 获取某时间段内根据某时间单位分组的统计数据
// @accept application/json
// @Produce application/json
// @Param data body model.StatisticsMeta true "标题"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /v1/statistics/interval/unit [get]
func GetIntervalUnitData(c *gin.Context) {
	var m request.GetIntervalUnitData
	if err := c.ShouldBindQuery(&m); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	data, err := service.GetIntervalUnitData(&m)
	if err != nil {
		response.FailWithDetailed(data, "获取 interval unit 失败", c)
		return
	}
	response.OkWithDetailed(data, "获取 interval unit data 成功", c)

}
