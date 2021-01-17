package model

import (
	"time"

	"gorm.io/gorm"
)


// BaseModel ...
type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}


// StatisticsMeta 统计 meta_id
type StatisticsMeta struct {
	BaseModel

	Title string `json:"title" gorm:"comment:标题"`
}

// StatisticsData 统计数据
type StatisticsData struct {
	ID        uint `gorm:"primarykey"`
	MetaID    uint `json:"meta_id" gorm:"comment:统计的 meta_id"`
	DateTime time.Time `json:"datetimes" gorm:"comment:统计时间`
	Val       float32 `json:"val" gorm:"comment:值"`
}
