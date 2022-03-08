package tool

import (
	"fmt"
	"github.com/PittYao/gin_seed/app/common/globalkey"
	"go.opentelemetry.io/otel/trace"
	"regexp"
	"unsafe"
)

func Bytes2str(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

// CompressStr 利用正则表达式压缩字符串，去除空格或制表符
func CompressStr(str string) string {
	if str == "" {
		return ""
	}
	//匹配一个或多个空白符的正则表达式
	reg := regexp.MustCompile("\\s+")
	return reg.ReplaceAllString(str, "")
}

func GenTraceId() (traceIdStr string) {
	traceId, err := trace.TraceIDFromHex(globalkey.TraceIDStr)
	if err != nil {
		fmt.Sprintln("traceId gen fail")
		return ""
	} else {
		traceIdStr = traceId.String()
		return traceIdStr
	}
}
