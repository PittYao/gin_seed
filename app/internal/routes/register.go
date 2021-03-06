// @contact.name API Support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

package routes

import (
	"github.com/PittYao/gin_seed/app/common/global"
	"github.com/PittYao/gin_seed/app/internal/middleware"
	"net/http"

	"github.com/PittYao/gin_seed/app/internal/routes/docs"
	// docs is generated by Swag CLI, you have to import it.

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

const (
	// DisableGinSwaggerEnvKey 设置该环境变量时关闭 swagger 文档
	DisableGinSwaggerEnvKey = "DISABLE_GIN_SWAGGER_dd"
)

// Register 注册swagger路由
func Register(httpHandler http.Handler) {
	app, ok := httpHandler.(*gin.Engine)
	if !ok {
		panic("HTTP handler must be *ginconfig.Engine")
	}

	// swagger 文档变量设置
	docs.SwaggerInfo.Title = global.CONFIG.Swagger.Title
	docs.SwaggerInfo.Description = global.CONFIG.Swagger.Desc
	docs.SwaggerInfo.Host = global.CONFIG.Swagger.Host
	docs.SwaggerInfo.BasePath = global.CONFIG.Swagger.BasePath
	docs.SwaggerInfo.Schemes = global.CONFIG.Swagger.Schemes

	// 访问swagger默认账号密码配置
	x := app.Group("", middleware.GinBasicAuth())
	{
		// Swagger 生成的在线 API 文档路由
		x.GET(global.CONFIG.Swagger.Url+"/*any", ginSwagger.DisablingWrapHandler(swaggerFiles.Handler, DisableGinSwaggerEnvKey))
		x.GET("/ping", Ping)
	}

	// 注册其他路由
	Routes(app)
}
