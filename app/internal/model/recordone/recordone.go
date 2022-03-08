package recordone

import (
	"time"
)

type RecordOne struct {
	ID                   int
	RtspUrl              string
	Host                 string
	SavePath             string
	FfmpegTransformState int
	FfmpegTransformCmd   string
	FfmpegSaveCmd        string
	FfmpegSaveState      int
	CreateTime           time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	UpdateTime           time.Time
}

func (RecordOne) TableName() string {
	return "record_one"
}
