package models

import "time"

type User struct {
	Id         string     `json:"id"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	User_type  string     `json:"user_type"`
	Created_at time.Time  `json:"created_at"`
	Updated_at *time.Time `json:"updated_at"`
	Deleted_at *time.Time `json:"deleted_at"`
}

type CreateUserModel struct {
	Username   string     `json:"username"`
	Password   string     `json:"password"`
	User_type  string     `json:"user_type"`
}

type UpdateUserModel struct {
	Id            string `json:"id" binding:"required"`
	Password   string     `json:"password"`
}
