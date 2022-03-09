package httpclient

import (
	"fmt"
	"github.com/PittYao/gin_seed/app/common/response"
	"github.com/guonaihong/gout"
	"time"
)

func ExampleHttpClient() {
	rsp := response.Response{}

	//code := 0
	err := gout.

		// POST请求
		POST("127.0.0.1:8080").

		// 打开debug模式
		Debug(true).

		// 设置查询字符串
		SetQuery(gout.H{"page": 10, "size": 10}).

		// 设置http header
		SetHeader(gout.H{"X-IP": "127.0.0.1", "sid": fmt.Sprintf("%x", time.Now().UnixNano())}).

		// SetJSON设置http body为json
		// 同类函数有SetBody, SetYAML, SetXML, SetForm, SetWWWForm
		SetJSON(gout.H{"text": "gout"}).

		// BindJSON解析返回的body内容
		// 同类函数有BindBody, BindYAML, BindXML
		BindJSON(&rsp).

		// http code
		// Code(&code).

		// 结束函数
		Do()

	// 判断错误
	if err != nil {
		fmt.Printf("send fail:%s\n", err)
	}
}
