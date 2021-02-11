package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
	"net/http"
	"github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm/dialects/mysql"

	_ "github.com/go-sql-driver/mysql"

	"gin-orm-prac/models"
)

var db *gorm.DB

func initModel() {
	db, err := gorm.Open("mysql", "root:1qaz!QAZ@tcp(localhost:3306)/gorm-todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err);
	}
	defer db.Close();

	db.AutoMigrate(&models.Todo{})
}


func main() {
	initModel()

	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.POST("/todo", CreateTodo)
		v1.GET("/todos", FetchAllTodo)
		v1.GET("/todo/:id", FetchSingleTodo)
	// 	v1.PUT("/todo/:id", UpdateTodo)
		v1.DELETE("/todo/:id", DeleteTodo)
	}

	router.Run()
}

type FormInput struct {
	// public members should be started with upper case, private with lowercase
	Title string `json:"title" binding:"required"`
	Description string `json:"description" binding:""`
}

func CreateTodo(context *gin.Context) {
	form := &FormInput{}

	db, err := gorm.Open("mysql", "root:1qaz!QAZ@tcp(localhost:3306)/gorm-todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err);
	}
	defer db.Close();

	if context.Bind(form) == nil {
		todo := models.Todo{
			Title: form.Title,
			Description: form.Description,
		}

		fmt.Println(todo);

		db.Create(&todo)
		context.JSON(200, gin.H{
			"statusCode": http.StatusOK,
			"title": form.Title,
			"description": form.Description,
			"message": "新增成功",
			"methods": "POST",
		})
	} else {
		context.JSON(500, gin.H{
			"statusCode": 500,
			"message": "新增錯誤",
			"methods": "POST",
		})
	}
}

func FetchAllTodo(context *gin.Context) {

	db, err := gorm.Open("mysql", "root:1qaz!QAZ@tcp(localhost:3306)/gorm-todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err);
	}
	defer db.Close();

	var todos []models.Todo
	db.Find(&todos);
	context.JSON(200, gin.H{
		"statusCode": http.StatusOK,
		"todos": todos,
		"message": "查詢成功",
		"methods": "GET",
	})

}

func FetchSingleTodo(context *gin.Context) {
	db, err := gorm.Open("mysql", "root:1qaz!QAZ@tcp(localhost:3306)/gorm-todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err);
	}
	defer db.Close();

	fmt.Println("id: => ")
	fmt.Println(context.Param("id"));

	todoId := context.Param("id")

	var todo models.Todo

	db.First(&todo,"id = ?",todoId)

	context.JSON(200, gin.H{
		"statusCode": http.StatusOK,
		"todo": todo,
		"message": "查詢成功",
		"methods": "GET",
	})

}

func DeleteTodo(context *gin.Context) {
	db, err := gorm.Open("mysql", "root:1qaz!QAZ@tcp(localhost:3306)/gorm-todo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err);
	}
	defer db.Close();

	todoId := context.Param("id")

	var todo models.Todo

	db.Delete(todo, todoId);

	context.JSON(200, gin.H{
		"statusCode": http.StatusOK,
		"message": "刪除成功",
		"methods": "DELETE",
		"deleteId": todoId,
	})
}