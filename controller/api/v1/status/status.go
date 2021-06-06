package status

type Code struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var Success = Code{200, "Ok"}
var Failure = Code{500, "Err"}

var ParamIncomplete = Code{500, "Parameter not complete"}

func Fail(msg string) Code {
	return Code{500, msg}
}
