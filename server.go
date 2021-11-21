package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type PostDataV1 struct {
	Content string `json:"content" binding:"required"`
}

type SearchDataV1 struct {
	Q string `form:"q" binding:"required"`
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
	err := c.BindJSON(&postParams)

	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": "an error occurred",
		})
		return
	}

	result := IndexDocument(postParams.Content)
	c.JSON(200, gin.H{
		"id": result.Id,
	})
}

func searchControllerV1(c *gin.Context) {
	var query SearchDataV1
	err := c.Bind(&query)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": "an error occurred",
		})
		return
	}

	result := Search(query.Q)

	fmt.Println(query.Q)

	c.JSON(200, gin.H{
		"docs": result,
	})
}