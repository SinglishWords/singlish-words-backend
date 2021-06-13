// Package log
// Ref: https://blog.csdn.net/wohu1104/article/details/107326794
// Wrap a zap logger
package log

import "go.uber.org/zap"

var Logger *zap.SugaredLogger

func init() {
	Logger = initLogger()
}
