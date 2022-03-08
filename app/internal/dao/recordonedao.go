package dao

import (
	"github.com/PittYao/gin_seed/app/common/global"
	"github.com/PittYao/gin_seed/app/common/request"
	. "github.com/PittYao/gin_seed/app/internal/model/recordone"
)

type PageRecordOne struct {
	*request.PagerSummary
	Data []*RecordOne `json:"data"`
}

type RecordOneReq struct {
	RtspUrl string ` json:"rtspUrl"  binding:"required"`
}

type RecordOnePageReq struct {
	request.Pager
}

func GetOneById(id int) *RecordOne {
	var recordOne RecordOne
	global.DB.Where("id = ?", id).Find(&recordOne)
	return &recordOne
}

func Page(pageReq RecordOnePageReq) (pageRecordOne PageRecordOne, err error) {
	var recordOnes []*RecordOne
	// query page
	if err := global.DB.Scopes(pageReq.Pager.Scope()).Find(&recordOnes).Error; err != nil {
		return pageRecordOne, err
	}

	// query total count
	global.DB.Model(&RecordOne{}).Count(&pageReq.Pager.TotalCount)
	pagerSummary := pageReq.Pager.Summary()

	// concat result
	pageRecordOne.PagerSummary = pagerSummary
	pageRecordOne.Data = recordOnes

	return pageRecordOne, nil
}
