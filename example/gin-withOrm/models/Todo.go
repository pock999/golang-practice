package models

type Todo struct {
	Id uint `gorm:"AUTO_INCREMENT"`
	Title string `json:"title"`
	Description string `json:"decreption"`
}