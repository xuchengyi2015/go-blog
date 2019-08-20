package blog

import (
	"go-blog/model"
	"go-blog/serializer"
)

type CreateBlogService struct {
	ID         uint   `json:"id"`
	Archive    string `json:"archive"`
	Title      string `json:"title"`
	Author     string `json:"author"`
	Content    string `json:"content"`
	Category   string `json:"category"`
	Tags       string `json:"tags"`
	Brief      string `json:"brief"`
	ThemeImage string `json:"theme_image"`
}

func (service *CreateBlogService) Save() (model.Blog, *serializer.Response) {
	if service.ID > 0 {
		return service.update()
	} else {
		return service.create()
	}
}

// 添加
func (service *CreateBlogService) create() (model.Blog, *serializer.Response) {
	blog := model.Blog{
		Title:      service.Title,
		Author:     service.Author,
		Category:   service.Category,
		Content:    service.Content,
		Tags:       service.Tags,
		Brief:      service.Brief,
		ThemeImage: service.ThemeImage,
	}
	if err := model.DB.Create(&blog).Error; err != nil {
		return blog, &serializer.Response{
			Status: 50001,
			Error:  err.Error(),
		}
	}

	return blog, &serializer.Response{
		Msg: "添加成功",
	}
}

// 修改
func (service *CreateBlogService) update() (model.Blog, *serializer.Response) {
	blog := model.Blog{
		Title:      service.Title,
		Author:     service.Author,
		Category:   service.Category,
		Content:    service.Content,
		Tags:       service.Tags,
		Brief:      service.Brief,
		ThemeImage: service.ThemeImage,
	}
	blog.ID = service.ID
	if err := model.DB.Save(&blog).Error; err != nil {
		return blog, &serializer.Response{
			Status: 50001,
			Error:  err.Error(),
		}
	}

	return blog, &serializer.Response{
		Msg: "修改成功",
	}
}

func Show(id string) serializer.Response {
	var blog model.Blog

	err := model.DB.First(&blog, id).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "error",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Data: serializer.BuildBlog(blog, false),
	}
}

func Delete(id string) serializer.Response {
	var blog model.Blog

	err := model.DB.First(&blog, id).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "error",
			Error:  err.Error(),
		}
	}

	err = model.DB.Delete(&blog).Error
	if err != nil {
		return serializer.Response{
			Status: 50001,
			Msg:    "error",
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Msg: "It's success to delete a blog.",
	}
}
