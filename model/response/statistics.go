package response

import (
	"gin-demo/model"
	"time"
)

// StatisticsMeta ...
type StatisticsMeta struct {
	Meta model.StatisticsMeta `json:"meta"`
}

// StatisticsMetaList ...
type StatisticsMetaList struct {
	Meta []model.StatisticsMeta `json:"metas"`
}

// Data ...
type Data struct {
	MetaID   uint      `json:"meta_id"`
	Datetime time.Time `json:"datetime"`
	Val      float32   `json:"val"`
}
