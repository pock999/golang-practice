package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	apiServer := gin.New()
	apiServer.GET("/hello", HelloWorld)
	apiServer.GET("/test", TestJson)

	apiServer.GET("/api/test", TestGet)
	apiServer.POST("/api/test", TestPost)
	apiServer.PUT("/api/test", TestPut)
	apiServer.DELETE("/api/test", TestDel)

	apiServer.Run(":3388")
}

func HelloWorld(context *gin.Context) {
  context.JSON(http.StatusOK, "Hello World")
}

func TestJson(context *gin.Context) {
	text := context.Query("text")
	if(text != "") {
		context.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message": text,
			"methods": "GET by query",
		})
	} else {
		context.JSON(http.StatusOK, gin.H{
			"statusCode": http.StatusOK,
			"message": "test",
			"methods": "GET by query",
		})
	}
}

// post data by json
type FormInput struct {
	// public members should be started with upper case, private with lowercase
	Input string `json:"input" binding:"required"`
}

func TestGet(context *gin.Context) {
	form := &FormInput{}

	if context.ShouldBindJSON(form) == nil {
		if form.Input != "" {
			context.JSON(200, gin.H{
				"statusCode": http.StatusOK,
				"message": form.Input,
				"methods": "GET",
			})
		} else {
			context.JSON(400, gin.H{
				"statusCode": 400,
				"message": "your input is empty",
				"methods": "GET",
			})
		}
	} else {
		fmt.Println("error");
	}
}

func TestPost(context *gin.Context) {
	form := &FormInput{}

	if context.Bind(form) == nil {
		if form.Input != "" {
			context.JSON(200, gin.H{
				"statusCode": http.StatusOK,
				"message": form.Input,
				"methods": "POST",
			})
		} else {
			context.JSON(400, gin.H{
				"statusCode": 400,
				"message": "your input is empty",
				"methods": "POST",
			})
		}
	} else {
		fmt.Println("error");
	}
}

func TestPut(context *gin.Context) {
	form := &FormInput{}

	if context.Bind(form) == nil {
		if form.Input != "" {
			context.JSON(200, gin.H{
				"statusCode": http.StatusOK,
				"message": form.Input,
				"methods": "PUT",
			})
		} else {
			context.JSON(400, gin.H{
				"statusCode": 400,
				"message": "your input is empty",
				"methods": "PUT",
			})
		}
	} else {
		fmt.Println("error");
	}
}

func TestDel(context *gin.Context) {
	form := &FormInput{}

	if context.Bind(form) == nil {
		if form.Input != "" {
			context.JSON(200, gin.H{
				"statusCode": http.StatusOK,
				"message": form.Input,
				"methods": "DELETE",
			})
		} else {
			context.JSON(400, gin.H{
				"statusCode": 400,
				"message": "your input is empty",
				"methods": "DELETE",
			})
		}
	} else {
		fmt.Println("error");
	}
}