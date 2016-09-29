package cli

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.EqualFold("OPTIONS", c.Request.Method) {
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
			c.Writer.Header().Add("Access-Control-Allow-Headers", "content-type, authorization, accept, x-requested-with")
			c.Writer.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
			c.Writer.Header().Add("Access-Control-Expose-Headers", "x-total-count, x-limit-count, link")
			c.Next()
			return
		}

		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Add("Access-Control-Allow-Headers", "content-type, authorization, accept, x-requested-with")
		c.Writer.Header().Add("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")
		c.Writer.Header().Add("Access-Control-Expose-Headers", "x-total-count, x-limit-count, link")
		c.Next()
	}
}
