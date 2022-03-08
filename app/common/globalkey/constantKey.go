package globalkey

/**
global constant key
*/

const (
	TraceIDStr = "4bf92f3577b34da6a3ce929d0e0e4736"

	ZapConsole = "console"
	ZapJson    = "json"

	EOFError    = "EOF"
	EOFErrorMsg = "没有传入body参数"
)

//软删除
var (
	DelStateNo  int64 = 0 //未删除
	DelStateYes int64 = 1 //已删除
)
