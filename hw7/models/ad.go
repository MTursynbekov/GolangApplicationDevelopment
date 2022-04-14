package models

import "github.com/lib/pq"

type Ad struct {
	ID          uint64         `json:"id" gorm:"primary_key"`
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Description string         `json:"description"`
	Images      pq.StringArray `json:"images"`
	Price       int            `json:"price"`
}
