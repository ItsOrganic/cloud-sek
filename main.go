package main

import (
	"cloud-sek/apploader"
	"cloud-sek/controller"
	"cloud-sek/database"
	"cloud-sek/service"
)

func main() {
	apploader.Init()
	database.InitDbConn()
	postRepo := database.PostRepository{}
	postService := service.NewPostService(&postRepo)
	controller.InitRoutes(postService)
}
