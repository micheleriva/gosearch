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

type UpdateDataV1 struct {
	ID 		string `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func RunServer() {
	r := gin.Default()
	r.GET("/ping", pingController)
	r.POST("/v1/insert", insertDocControllerV1)
	r.PUT("/v1/update", updateDocControllerV1)
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

	result := IndexDocument(postParams.Content, "")
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

	c.JSON(200, gin.H{
		"docs": result,
	})
}

func updateDocControllerV1(c *gin.Context) {
	var data UpdateDataV1
	err := c.Bind(&data)
	if err != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": "an error occurred",
		})
		return
	}

	updateError := UpdateDocument(data.ID, data.Content)
	if updateError != nil {
		fmt.Println(err)
		c.JSON(500, gin.H{
			"error": "an error occurred",
		})
		return
	}

	c.JSON(200, gin.H{
		"result": fmt.Sprintf("Document %s updated successfully", data.ID),
	})
}
