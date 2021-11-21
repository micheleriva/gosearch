package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type PostDataV1 struct {
	content string `json:"content" binding:"required"`
}

func RunServer() {
	r := gin.Default()
	r.GET("/ping", pingController)
	r.POST("/v1/insert", insertDocControllerV1)
	r.GET("/v1/search", searchControllerV1)
	r.Run()
}

func pingController(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func insertDocControllerV1(c *gin.Context) {
	var postParams PostDataV1
	c.BindJSON(&postParams)

	result := IndexDocument(postParams.content)
	c.JSON(200, gin.H{
		"id": result.Id,
	})
}

func searchControllerV1(c *gin.Context) {
	searchQuery, _ := c.GetQuery("q")

	result := Search(searchQuery)

	fmt.Println(result)

	c.JSON(200, gin.H{
		"docs": nil,
	})
}