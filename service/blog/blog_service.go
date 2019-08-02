package blog

import (
	"go-blog/model"
	"go-blog/serializer"
)

type CreateBlogService struct {
	Title    string `json:"title"`
	Author   string `json:"author"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

func (service *CreateBlogService) Create() serializer.Response {
	return serializer.Response{}
}

func List(keyword string, category string) serializer.Response {
	var blogs []model.Blog
	var err error
	if category == "all" {
		err = model.DB.Find(&blogs).Error
	} else {
		err = model.DB.Where("category = ?", category).Find(&blogs).Error
	}
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "error",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildBlogs(blogs),
	}
}

func Show(id string) serializer.Response {
	var blog model.Blog

	err := model.DB.First(&blog).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "error",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildBlog(blog),
	}
}
