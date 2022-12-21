package models

import "time"

type Product struct {
	Id         string     `json:"id"`
	CategoryId string     `json:"category_id" binding:"required"`
	Title      string     `json:"title"`
	Descrip    string     `json:"description"`
	Price      string     `json:"price"`
	Created_at time.Time  `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at"`
}

type CreateProductModel struct {
	CategoryId string `json:"category_id" binding:"required"`
	Title      string `json:"title"`
	Descrip    string `json:"description"`
	Price      string `json:"price"`
}

type UpdateProductModel struct {
	Id    string `json:"id" binding:"required"`
	Title string `json:"title"`
	Price      string `json:"price"`
}
