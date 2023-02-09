package main

import (
	handle "Hoa/handle"
	"github.com/gin-gonic/gin"
)

func setupRouter() {
	r := gin.Default()
	content := r.Group("/content")
	{
		content.POST("/advise", handle.AddAdvise())
	}
	r.Run(":8080")
}
