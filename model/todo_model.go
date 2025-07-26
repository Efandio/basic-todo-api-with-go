package model

import "time"

type Todo struct {
	ID       string	`json:"id"`
	Title    string	`json:"title"`
	Complete bool	`json:"complete"`
	CreateAt time.Time	`json:"create_at"`
	UpdateAt time.Time	`json:"update_at"`
	UserID	string	`json:"user_id"`
}