package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	"strings"
)

// Cors 跨域配置
func Cors() gin.HandlerFunc {
	godotenv.Load()
	corsDomain := os.Getenv("CORS_DOMIAN")

	config := cors.DefaultConfig()
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Cookie"}
	config.AllowOrigins = strings.Split(corsDomain, "|")
	config.AllowCredentials = true
	return cors.New(config)
}
