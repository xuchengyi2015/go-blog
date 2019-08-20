package blog

import (
	"go-blog/model"
	"go-blog/serializer"
	"go-blog/util"
	"strings"
)

type QueryBlogService struct {
	Limit          uint   `json:"limit"`
	Offset         uint   `json:"offset"`
	Category       string `json:"category"`
	Tag            string `json:"tag"`
	Classification string `json:"classification"`
}

func (service *QueryBlogService) List() *serializer.Response {
	var blogs []model.Blog
	var err error

	cond := ""
	if !util.IsEmpty(service.Tag) {
		cond = " tags like '%" + service.Tag + "%' AND"
	}
	if !util.IsEmpty(service.Category) {
		cond += " category = '" + service.Category + "' AND"
	}
	if !util.IsEmpty(service.Classification) {
		cond += " classification = '" + service.Classification + "' AND"
	}
	cond = strings.Trim(cond, "AND")

	err = model.DB.Where(cond).Limit(service.Limit).Order("id desc").Offset(service.Offset).Find(&blogs).Error
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

type CategoryTemp struct {
	Category string `json:"category"`
	Num      int    `json:"num"`
}

func Categories() *serializer.Response {
	var categoryTemps []CategoryTemp

	sql := "SELECT category ,COUNT(1) AS num FROM blogs GROUP BY category ORDER BY category	;"
	rows, err := model.DB.Raw(sql).Rows()
	defer rows.Close()
	if err != nil {
		return &serializer.Response{
			Status: 50001,
			Error:  err.Error(),
		}
	}

	for rows.Next() {
		categoryTemp := CategoryTemp{}
		model.DB.ScanRows(rows, &categoryTemp)
		categoryTemps = append(categoryTemps, categoryTemp)
	}

	return &serializer.Response{
		Data: categoryTemps,
	}
}
