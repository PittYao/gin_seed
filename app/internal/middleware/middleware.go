package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/PittYao/gin_seed/app/common/ginconfig"
	"github.com/PittYao/gin_seed/app/common/global"
	"github.com/PittYao/gin_seed/app/common/tool"
	"github.com/axiaoxin-com/logging"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"runtime/debug"
	"strings"
	"time"
)

// DefaultGinMiddlewares 默认使用的中间件列表
func DefaultGinMiddlewares() []gin.HandlerFunc {
	m := []gin.HandlerFunc{
		// 记录请求处理日志，最顶层执行
		GinLogger(),
		// 捕获 panic 保存到 context 中由 GinLogger 统一打印， panic 时返回 -1 JSON
		GinRecovery(),
		// 跨域
		Cors(),
	}

	return m
}

// GinBasicAuth ginconfig 的基础认证中间件
// 加到 ginconfig app 的路由中可以对该路由添加 basic auth 登录验证
func GinBasicAuth() gin.HandlerFunc {
	username := global.CONFIG.BasicAuth.UserName
	password := global.CONFIG.BasicAuth.Password
	logging.Debug(nil, "Basic auth username:"+username+" password:"+password)
	return gin.BasicAuth(gin.Accounts{
		username: password,
	})
}

// GinLogger 接收gin框架默认的日志
func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		// read request body
		var requestBody string
		if method == http.MethodPost || method == http.MethodPut {
			requestBody = ReadRequestBody(c)
		}

		blw := &ginconfig.CustomResponseWriter{Body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()

		cost := time.Since(start)

		// read response json body
		responseBody := json.RawMessage(blw.Body.String())

		if method == http.MethodPost || method == http.MethodPut {
			global.LOG.Info(path,
				zap.String("trace_id", tool.GenTraceId()),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("request_body", requestBody),
				zap.Any("response_body", &responseBody),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)

		}

		if method == http.MethodGet || method == http.MethodDelete {
			requestParams := c.Request.URL.RawQuery

			global.LOG.Info(path,
				zap.String("trace_id", tool.GenTraceId()),
				zap.Int("status", c.Writer.Status()),
				zap.String("method", method),
				zap.String("path", path),
				zap.String("request_params", requestParams),
				zap.Any("response_body", &responseBody),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost),
			)
		}

	}
}

//ReadRequestBody 从request中读取body
func ReadRequestBody(c *gin.Context) string {
	var bodyBytes []byte

	bodyBytes, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Errorf("invalid request body")
	}

	// 新建缓冲区并替换原有Request.body
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// 当前函数可以使用body内容
	body := bodyBytes
	str := tool.Bytes2str(body)

	// 去除json中的制表符
	compressStr := tool.CompressStr(str)
	return compressStr
}

// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
func GinRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				global.LOG.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(httpRequest)),
					zap.String("stack", string(debug.Stack())),
				)
			}
		}()
		c.Next()
	}
}

//Cors 跨域
func Cors() gin.HandlerFunc {
	return func(context *gin.Context) {
		method := context.Request.Method
		origin := context.Request.Header.Get("Origin")
		var headerKeys []string
		for k, _ := range context.Request.Header {
			headerKeys = append(headerKeys, k)
		}
		headerStr := strings.Join(headerKeys, ",")
		if headerStr != "" {
			headerStr = fmt.Sprintf("access-control-allow-origin, access-control-allow-headers, %s", headerStr)
		} else {
			headerStr = "access-control-allow-origin, access-control-allow-headers"
		}

		if origin != "" {
			context.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			context.Header("Access-Control-Allow-Origin", "*") // 设置允许访问所有域
			context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			context.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token, Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
			context.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
			context.Header("Access-Control-Max-Age", "172800")
			context.Header("Access-Control-Allow-Credentials", "false")
			context.Set("content-type", "application/json") //// 设置返回格式是json
		}

		if method == "OPTIONS" {
			context.JSON(http.StatusOK, "Options Request!")
		}
		//处理请求
		context.Next()
	}
}
