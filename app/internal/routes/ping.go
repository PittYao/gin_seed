package routes

import (
	"github.com/PittYao/gin_seed/app/common/response"
	"github.com/gin-gonic/gin"
)

//Ping godoc
//@Summary 默认的 Ping 接口
//@Description 返回 server 相关信息，可以用于健康检查
//@Tags x
//@Accept json
//@Produce json
//@Success 200 {object} response.Response
//@Security ApiKeyAuth
//@Security BasicAuth
//@Param trace_id header string false "you can set custom trace id in header"
//@Router /ping [get]
func Ping(c *gin.Context) {
	response.OK(c, "ok")
	return
}
