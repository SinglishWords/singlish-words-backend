package util

import (
	"fmt"
	"runtime"
	"strconv"
)

type TraceBackError struct {
	msg string
	Err error
}

func NewTraceBackError(previousMsg string) *TraceBackError {
	_, srcName, line, _ := runtime.Caller(1) // (1)
	prefix := "[" + srcName + ":" + strconv.Itoa(line) + "] "
	msg := fmt.Sprintf(prefix + previousMsg)

	return &TraceBackError{msg: msg}
}

func (e *TraceBackError) Error() string {
	return e.msg + ": " + e.Err.Error()
}

// func NewBizErrorf(format string, args ...interface{}) *TraceBack {
// 	_, srcName, line, _ := runtime.Caller(1) // (1)
// 	// [源文件名:行号]
// 	prefix := "[" + srcName + ":" + strconv.Itoa(line) + "] "
// 	msg := fmt.Sprintf(prefix+format, args...)

// 	return &TraceBack{msg: msg}
// }
