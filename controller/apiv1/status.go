// Package status
// import "singlishwords/controller/api/v1/status"
// Definition of API v1 status codes
// https://golang.org/src/net/http/status.go

package apiv1

type HttpStatus struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var StatusOK = HttpStatus{200, "Ok."}
var StatusCreated = HttpStatus{201, "Successful add this object to database."}
var StatusFailure = HttpStatus{500, "Err."}

var StatusPostParamError = HttpStatus{500, "Cannot extract parameters in form correctly."}
var StatusQueryParamError = HttpStatus{500, "Cannot extract parameters in query string correctly."}

func StatusFail(msg string) HttpStatus {
	return HttpStatus{500, msg}
}
