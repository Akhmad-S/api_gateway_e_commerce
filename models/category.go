package models

import "time"

type Category struct {
	Id            string     `json:"id"`
	CategoryTitle string     `json:"category_title"`
	Created_at    time.Time  `json:"created_at"`
	Updated_at    *time.Time `json:"updated_at"`
	Deleted_at    *time.Time `json:"deleted_at"`
}

type CreateCategoryModel struct {
	CategoryTitle string     `json:"category_title"`
}

type UpdateCategoryModel struct {
	Id       string `json:"id" binding:"required"`
	CategoryTitle string     `json:"category_title"`
}
