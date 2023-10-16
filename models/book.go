package models

import "gorm.io/gorm"

type Author struct {
	Name string `json:"name"`
	ID   uint   `json:"id"`
}

type Book struct {
	gorm.Model
	AuthorID    uint   `json:"author_id"`
	Author      uint   `json:"-" gorm:"foreignKey:AuthorID"`
	Title       string `json:"title"`
	ReleaseYear string `json:"release_year"`
}
