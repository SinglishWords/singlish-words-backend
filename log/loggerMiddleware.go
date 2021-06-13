package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"singlishwords/config"
	"time"
)

func RouteLogger() gin.HandlerFunc {
	//gin.DefaultWriter = os.Stdout
	if config.Log.Route != "stdout" {
		f, err := os.Create(config.Log.Route)
		if err != nil {
			Logger.Fatalf("Error when create route log file: %v", err)
		}
		gin.DefaultWriter = io.MultiWriter(f)
	}

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
			param.Latency,
			param.Request.UserAgent(),
			param.ErrorMessage,
		)
	})
}
