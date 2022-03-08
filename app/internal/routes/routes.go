// 在这个文件中注册 URL handler

package routes

import (
	. "github.com/PittYao/gin_seed/app/internal/service/example_service"
	"github.com/gin-gonic/gin"
)

// Routes 注册 API URL 路由
func Routes(app *gin.Engine) {
	// TODO: 在这里注册你的 ginconfig API，如： app.GET("/", HandlerFunc)

	recordGroup := app.Group("/record")
	{
		recordOneGroup := recordGroup.Group("/one")
		{
			recordOneGroup.POST("", RecordOne)
			recordOneGroup.POST("/page", RecordOnePage)
		}
	}

}
