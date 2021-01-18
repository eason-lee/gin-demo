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
	MetaID uint      `json:"meta_id" binding:"required"`
	Dt     time.Time `json:"d_t" binding:"required" time_format:"2006-01-02"`
	Val    float32   `json:"val" binding:"required"`
}

// GetIntervalData ...
type GetIntervalData struct {
	MetaIDs string    `json:"meta_ids" form:"meta_ids" binding:"required"`
	Start   time.Time `json:"start" form:"start" binding:"required" time_format:"2006-01-02"`
	End     time.Time `json:"end" form:"end" binding:"required" time_format:"2006-01-02"`
}

// GetIntervalUnitData ...
type GetIntervalUnitData struct {
	MetaIDs string    `json:"meta_ids" form:"meta_ids" binding:"required"`
	Start   time.Time `json:"start" form:"start" binding:"required" time_format:"2006-01-02"`
	End     time.Time `json:"end" form:"end" binding:"required" time_format:"2006-01-02"`
	Unit    string    `json:"unit" form:"unit" binding:"required"`
}
