package request

import "time"

// CreateMeta ...
type CreateMeta struct {
	Title string `json:"title" binding:"required"`
}

// GetMeta ...
type GetMeta struct {
	Title string `json:"title" form:"title" binding:"required"`
}

// CreateData ...
type CreateData struct {
	MetaID   uint      `json:"meta_id" binding:"required"`
	Datetime time.Time `json:"datetime" binding:"required" time_format:"2006-01-02"`
	Val      float32   `json:"val" binding:"required"`
}
