package service

import (
	"errors"
	"gin-demo/global"
	"gin-demo/model/request"
	"gin-demo/model/response"
	"strings"
)

// GetIntervalData ...
func GetIntervalData(d *request.GetIntervalData) []response.IntervalData {
	metaIds := strings.Split(d.MetaIDs, ",")

	results := []response.IntervalData{}
	global.DB.Table("statistics_data").Select("meta_id, SUM(val) AS total").Where(
		"dt >= ? AND dt <= ? AND meta_id IN ? ",
		d.Start,
		d.End,
		metaIds,
	).Group("meta_id").Scan(&results)

	return results
}

// GetIntervalUnitData ...
func GetIntervalUnitData(d *request.GetIntervalUnitData) (results []response.IntervalUnitData, err error) {
	metaIds := strings.Split(d.MetaIDs, ",")

	units := map[string]string{
		"day":   "%Y-%m-%d",
		"week":  "%x-%v",
		"month": "%Y-%m",
		"year":  "%Y",
	}
	formatter, ok := units[d.Unit]
	if !ok {
		err = errors.New("unit 不支持")
		return
	}

	results = []response.IntervalUnitData{}
	global.DB.Table("statistics_data").Select(
		"meta_id, DATE_FORMAT(dt, ? ) AS time, SUM(val) AS total",
		formatter,
	).Where(
		"dt >= ? AND dt <= ? AND meta_id IN ? ",
		d.Start,
		d.End,
		metaIds,
	).Group("meta_id,time").Scan(&results)

	return

}
