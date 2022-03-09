package example_service

import (
	"github.com/PittYao/gin_seed/app/common/response"
	"github.com/PittYao/gin_seed/app/internal/dao"
	"github.com/gin-gonic/gin"
)

//RecordOne godoc
//@Summary 录制单个流
//@Tags stream one
//@Accept json
//@Produce json
//@Success 200 {object} response.Response
//@Param recordOneReq body dao.RecordOneReq true " "
//@Router /record/one [post]
func RecordOne(c *gin.Context) {
	var recordOneReq dao.RecordOneReq
	if err := c.ShouldBindJSON(&recordOneReq); err != nil {
		response.JsonBindError(c, err)
		return
	}

	recordOne := dao.GetOneById(2)

	response.OK(c, recordOne)
}

// RecordOnePage godoc
//@Summary 分页查询
//@Tags stream one
//@Accept json
//@Produce json
//@Success 200 {object} response.Response
//@Param recordOneReq body dao.RecordOnePageReq true " "
//@Router /record/one/page [post]
func RecordOnePage(c *gin.Context) {
	var recordOnePageReq dao.RecordOnePageReq
	if err := c.ShouldBindJSON(&recordOnePageReq); err != nil {
		response.Err(c, err.Error())
		return
	}

	pageRecordOne, err := dao.Page(recordOnePageReq)
	if err != nil {
		response.Err(c, err.Error())
		return
	}

	response.OK(c, pageRecordOne)
}
