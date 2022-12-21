package models

import "time"

type Order struct {
	Id           string    `json:"id"`
	Product_id   string    `json:"product_id"`
	Quantity     int32     `json:"quantity"`
	User_name    string    `json:"user_name"`
	User_address string    `json:"user_address"`
	User_phone   string    `json:"user_phone"`
	Created_at   time.Time `json:"created_at"`
}

type CreateOrderModel struct {
	Product_id   string `json:"product_id"`
	Quantity     int32  `json:"quantity"`
	User_name    string `json:"user_name"`
	User_address string `json:"user_address"`
	User_phone   string `json:"user_phone"`
}

type PackedOrderModel struct {
	Id           string    `json:"id"`
	Product_id   string    `json:"product_id"`
	Quantity     int32     `json:"quantity"`
	User_name    string    `json:"user_name"`
	User_address string    `json:"user_address"`
	User_phone   string    `json:"user_phone"`
	Created_at   time.Time `json:"created_at"`
	Product      Product   `json:"product"`
}
