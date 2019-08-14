package api

import (
	"github.com/gin-gonic/gin"
	"go-blog/serializer"
	"go-blog/service/blog"
)

func BlogList(c *gin.Context) {
	var service blog.QueryBlogService
	if err := c.ShouldBind(&service); err != nil {
		c.JSON(200, err)
	} else {
		res := service.List()
		c.JSON(200, res)
	}
}

func BlogShow(c *gin.Context) {
	res := blog.Show(c.Param("id"))
	c.JSON(200, res)
}

func BlogDelete(c *gin.Context) {
	res := blog.Delete(c.Param("id"))
	c.JSON(200, res)
}

func BlogSave(c *gin.Context) {
	var service blog.CreateBlogService
	if err := c.ShouldBind(&service); err == nil {
		if blog, err := service.Save(); err != nil {
			c.JSON(200, err)
		} else {
			res := serializer.BuildBlog(blog)
			c.JSON(200, res)
		}
	} else {
		c.JSON(200, err)
	}
}

func BlogTags(c *gin.Context) {
	res := blog.Tags()
	c.JSON(200, res)
}
