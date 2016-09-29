package cli

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func CorsMiddleware(cfg *Cors) gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.EqualFold("OPTIONS", c.Request.Method) {
			c.Writer.WriteHeader(200)
			c.Writer.Header().Add("Access-Control-Allow-Origin", cfg.AllowOrigin)
			c.Writer.Header().Add("Access-Control-Allow-Headers", cfg.AllowHeaders)
			c.Writer.Header().Add("Access-Control-Allow-Methods", cfg.AllowMethods)
			c.Writer.Header().Add("Access-Control-Expose-Headers", cfg.ExposeHeaders)
			c.Next()
			return
		}

		c.Writer.Header().Add("Access-Control-Allow-Origin", cfg.AllowOrigin)
		c.Writer.Header().Add("Access-Control-Allow-Headers", cfg.AllowHeaders)
		c.Writer.Header().Add("Access-Control-Allow-Methods", cfg.AllowMethods)
		c.Writer.Header().Add("Access-Control-Expose-Headers", cfg.ExposeHeaders)
		c.Next()
	}
}
