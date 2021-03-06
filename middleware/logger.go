package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"singlishwords/config"
	"singlishwords/log"
	"time"
)

func RouteLogger() gin.HandlerFunc {
	//gin.DefaultWriter = os.Stdout
	if config.Log.Route != "stdout" {
		f, err := os.OpenFile(config.Log.Route,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Logger.Fatalf("Error when create route log file: %v", err)
		}
		gin.DefaultWriter = io.MultiWriter(f)
	}

	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		// your custom format
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC3339),
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
