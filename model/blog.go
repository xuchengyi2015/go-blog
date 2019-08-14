package model

import "github.com/jinzhu/gorm"

type Blog struct {
	gorm.Model
	Title    string
	Author   string
	Category string
	Content  string `grom:"type:text;"`
	Comments string
	Tags string
	Brief string
	ThemeImage string
}