package middleware

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:     []string{"https://foo.com"},
			AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
			AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			AllowOriginFunc: func(origin string) bool {
				return origin == "https://github.com"
			},
			MaxAge: 12 * time.Hour,
		})
		c.Next()
	}
}
