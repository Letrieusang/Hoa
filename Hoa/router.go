package main

import (
	docs "Hoa/docs"
	handle "Hoa/handle"
	"github.com/gin-gonic/gin"
)

func setupRouter() {
	r := gin.Default()

	docs.SwaggerInfo.BasePath = "/content"

	content := r.Group("/content")
	{
		content.POST("/advise", handle.AddAdvise())
		content.POST("/confess", handle.AddConfess())
	}
	r.Run(":8080")
}
