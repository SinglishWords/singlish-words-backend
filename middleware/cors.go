package middleware

import (
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		//method := c.Request.Method

		c.Header("Access-Control-Allow-Origin", "localhost:3000")
		c.Header("Access-Control-Allow-Methods", "POST, GET, PATCH, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		//c.Header("Access-Control-Allow-Credentials", "true")

		c.Next()
	}
}