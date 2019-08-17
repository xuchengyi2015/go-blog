package blog

import (
	"go-blog/model"
	"go-blog/serializer"
	"go-blog/util"
	"strings"
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

// QueryField: Archive,Category,Tags
type QueryBlogService struct {
	Limit      uint   `json:"limit"`
	Offset     uint   `json:"offset"`
	QueryField string `json:"query_field"`
	QueryValue string `json:"query_value"`
	Category   string `json:"category"`
	Tag string `json:"tag"`
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
		Data: serializer.BuildBlog(blog),
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

func (service *QueryBlogService) List() *serializer.Response {
	var blogs []model.Blog
	var err error
	if service.QueryField != "" {
		err = model.DB.Where(service.QueryField+" = ?", service.QueryValue).Order("id desc").Limit(service.Limit).Offset(service.Offset).Find(&blogs).Error
	} else {
		err = model.DB.Limit(service.Limit).Order("id desc").Offset(service.Offset).Find(&blogs).Error
	}

	if err != nil {
		return &serializer.Response{
			Status: 50001,
			Msg:    "error",
			Error:  err.Error(),
		}
	}
	return &serializer.Response{
		Data: serializer.BuildBlogs(blogs),
	}
}

func Tags() *serializer.Response {
	var blogs []model.Blog
	if err := model.DB.Select([]string{"tags"}).Find(&blogs).Error; err != nil {
		return &serializer.Response{
			Status: 50001,
			Msg:    err.Error(),
		}
	}
	var tags []string
	for _, v := range blogs {
		for _, t := range strings.Split(v.Tags, "|") {
			if util.StringsContains(tags, t) == -1 {
				tags = append(tags, t)
			}
		}
	}
	return &serializer.Response{
		Data: tags,
	}
}
