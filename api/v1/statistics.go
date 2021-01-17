package v1

import (
	"gin-demo/global"
	"gin-demo/model"
	"gin-demo/model/request"
	"gin-demo/model/response"
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
		global.LOG.Error("注册失败", zap.Any("err", err))
		response.FailWithDetailed(response.StatisticsMeta{Meta: meta}, "添加失败", c)
	} else {
		response.OkWithDetailed(response.StatisticsMeta{Meta: meta}, "添加成功", c)
	}

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
	err := global.DB.Where("title LIKE ?", "%" + m.Title + "%").Find(&metas).Error

	if err != nil {
		global.LOG.Error("获取 meta 失败", zap.Any("err", err))
		response.FailWithDetailed(response.StatisticsMetaList{Meta: metas}, "获取 meta 失败", c)
	} else {
		response.OkWithDetailed(response.StatisticsMetaList{Meta: metas}, "获取 meta 成功", c)
	}

}
