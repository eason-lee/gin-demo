package response

import (
	"gin-demo/model"
)

// StatisticsMeta ...
type StatisticsMeta struct {
	Meta model.StatisticsMeta `json:"meta"`
}

// StatisticsMetaList ...
type StatisticsMetaList struct {
	Meta []model.StatisticsMeta `json:"metas"`
}

// IntervalData ...
type IntervalData struct {
	MetaID uint    `json:"meta_id"`
	Total  float32 `json:"val"`
}

// IntervalUnitData ...
type IntervalUnitData struct {
	MetaID uint    `json:"meta_id"`
	Time   string  `json:"d_t"`
	Total  float32 `json:"val"`
}
