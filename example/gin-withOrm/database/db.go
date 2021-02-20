package database

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"

	"gin-orm-prac/models"
)

var db *gorm.DB

func InitModel() {
	db, err := gorm.Open("mysql", "root:1qaz!QAZ@tcp(localhost:3306)/gorm-todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err);
	}
	defer db.Close();

	db.AutoMigrate(&models.Todo{})
}