package serializer

import (
	"go-blog/model"
	"go-blog/util"
)

type Blog struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Author    string `json:"author"`
	Category  string `json:"category""`
	Content   string `json:"content"`
	Comments  string `json:"comments"`
	CreatedAt string `json:"created_at"`
}

func BuildBlog(blog model.Blog) Blog {
	return Blog{
		ID:        blog.ID,
		Title:     blog.Title,
		Author:    blog.Author,
		Category:  blog.Category,
		Content:   blog.Content,
		Comments:  blog.Comments,
		CreatedAt: blog.CreatedAt.Format(util.STANDARD_TIME_FORMAT),
	}
}

func BuildBlogs(items []model.Blog) (blogs []Blog) {
	for _, v := range items {
		blog := BuildBlog(v)
		blog.Content=util.SubString(blog.Content,0,200)
		blogs = append(blogs, blog)
	}
	return blogs
}
