package ginconfig

import (
	"context"
	"github.com/PittYao/gin_seed/app/common/global"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/PittYao/gin_seed/app/statics"
	"github.com/axiaoxin-com/goutils"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func init() {
	// 替换 ginconfig 默认的 validator，更加友好的错误信息
	binding.Validator = &goutils.GinStructValidator{}
	// causes the json binding Decoder to unmarshal a number into an interface{} as a Number instead of as a float64.
	binding.EnableDecoderUseNumber = true
}

// NewGinEngine 根据参数创建 ginconfig 的 router engine
// middlewares 需要使用到的中间件列表，默认不为 engine 添加任何中间件
func NewGinEngine(middlewares ...gin.HandlerFunc) *gin.Engine {
	// set ginconfig mode
	gin.SetMode(global.CONFIG.Mode)

	engine := gin.New()

	// ///a///b -> /a/b
	engine.RemoveExtraSlash = true

	// use middlewares
	for _, middleware := range middlewares {
		engine.Use(middleware)
	}

	// register statics
	staticsURL := global.CONFIG.Statics.Url
	if staticsURL != "" {
		engine.StaticFS(staticsURL, http.FS(&statics.Files))
	}

	return engine
}

func Run(app http.Handler) {
	addr := global.CONFIG.ListenOn
	srv := &http.Server{
		Addr:         addr,
		Handler:      app,
		ReadTimeout:  5 * time.Minute,
		WriteTimeout: 10 * time.Minute,
	}

	// 启动 http server
	go func() {
		var ln net.Listener
		var err error
		if strings.ToLower(strings.Split(addr, ":")[0]) == "unix" {
			ln, err = net.Listen("unix", strings.Split(addr, ":")[1])
			if err != nil {
				panic(err)
			}
		} else {
			ln, err = net.Listen("tcp", addr)
			if err != nil {
				panic(err)
			}
		}
		if err := srv.Serve(ln); err != nil {
			global.LOG.Sugar().Errorf("Server runing error: %s", err.Error())
		}
	}()
	global.LOG.Sugar().Infof("Server is running on %s", srv.Addr)

	// 监听中断信号， WriteTimeout 时间后优雅关闭服务
	// syscall.SIGTERM 不带参数的 kill 命令
	// syscall.SIGINT ctrl-c kill -2
	// syscall.SIGKILL 是 kill -9 无法捕获这个信号
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	global.LOG.Sugar().Info("Server is shutting down.")

	// 创建一个 context 用于通知 server 3 秒后结束当前正在处理的请求
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.LOG.Sugar().Error("Server shutdown with error: " + err.Error())
	}
	global.LOG.Sugar().Info("Server exit.")
}
