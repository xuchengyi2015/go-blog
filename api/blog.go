package api

import (
	"github.com/gin-gonic/gin"
	"go-blog/service/blog"
)

func BlogList(c *gin.Context) {
	res := blog.List(c.Query("keyword"), c.Query("category"))
	c.JSON(200, res)
}

func BlogShow(c *gin.Context) {
	res := blog.Show(c.Param("id"))
	c.JSON(200, res)
}

func BlogDelete(c *gin.Context) {
	res := blog.Delete(c.Param("id"))
	c.JSON(200, res)
}
