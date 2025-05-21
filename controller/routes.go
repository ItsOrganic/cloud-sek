package controller

import (
	_ "cloud-sek/docs"
	"cloud-sek/handler"
	"cloud-sek/service"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(postService *service.PostService) {
	handler.PostService = postService
	router := gin.Default()
	router.POST("/create", handler.CreatePost)
	router.GET("/post/:id", handler.GetPostById)
	router.GET("/post/:id/comments", handler.GetCommentsByPostID)
	router.POST("/post/:id/comment", handler.CreateComment)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":8080")
}
