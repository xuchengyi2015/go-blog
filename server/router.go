package server

import (
	"github.com/gin-gonic/gin"
	"go-blog/api"
	"go-blog/middleware"
	"os"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	r.LoadHTMLGlob("dist/*.html")                    // 添加入口index.html
	r.LoadHTMLFiles("dist/*/*")                      // 添加资源路径
	r.Static("/css", "./dist/css")                   // 添加资源路径
	r.Static("/js", "./dist/js")                     // 添加资源路径
	r.Static("/fonts", "./dist/fonts")               // 添加资源路径
	r.StaticFile("/favicon.ico", "dist/favicon.ico") // 添加资源路径
	r.StaticFile("/", "dist/index.html")             //前端接口

	// 中间件, 顺序不能改
	r.Use(middleware.Session(os.Getenv("SESSION_SECRET")))
	r.Use(middleware.Cors())
	r.Use(middleware.CurrentUser())
	//r.Use(static.Serve("/", static.LocalFile("/dist", false)))
	r.NoRoute(func(c *gin.Context) {
		c.File("dist/index.html")
	})

	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户登录
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", api.UserLogin)

		v1.GET("blog/:id", api.BlogShow)
		v1.POST("blogs", api.BlogList)
		v1.POST("blog", api.BlogSave)
		v1.DELETE("blog/:id", api.BlogDelete)
		v1.GET("tags", api.BlogTags)
		v1.GET("categories", api.BlogCategories)

		// 需要登录保护的
		auth := r.Group("")
		auth.Use(middleware.AuthRequired())
		{
			// User Routing
			auth.GET("user/me", api.UserMe)
			auth.DELETE("user/logout", api.UserLogout)
		}
	}
	return r
}
