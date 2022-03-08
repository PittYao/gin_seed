package main

import (
	"flag"
	"github.com/PittYao/gin_seed/app/common/db"
	"github.com/PittYao/gin_seed/app/common/ginconfig"
	"github.com/PittYao/gin_seed/app/common/global"
	"github.com/PittYao/gin_seed/app/common/zap"
	"github.com/PittYao/gin_seed/app/internal/middleware"
	"github.com/PittYao/gin_seed/app/internal/routes"
	"github.com/zeromicro/go-zero/core/conf"
)

// configFile path
var configFile = flag.String("c", "app/etc/config.yaml", "the config file")

func main() {
	initApp()

	// new app
	app := ginconfig.NewGinEngine(middleware.DefaultGinMiddlewares()...)

	routes.Register(app)

	ginconfig.Run(app)
}

//initApp
func initApp() {
	// load configFile
	flag.Parse()
	conf.MustLoad(*configFile, &global.CONFIG)

	// init zap logger
	global.LOG = zap.Zap()

	// init mysql
	global.DB = db.GormMySQL()

}
