package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hakuuww/go-gin/controllers"
	"log"
	"github.com/hakuuww/go-gin/middlewares"
)

func main() {
	server := gin.Default()
	server.Use(middlewares.MyAuth())

	videocontroller := controllers.NewVideoController()

	videoGroup := server.Group("/videos")
	videoGroup.Use(middlewares.MyLogger())

	server.StaticFile("/sickfifaclip", "./resources/fifaclip.mp4")

	videoGroup.GET("/", videocontroller.GetAll)
	videoGroup.POST("/", videocontroller.Create)
	videoGroup.PUT("/:id", videocontroller.Update)
	videoGroup.DELETE("/:id", videocontroller.Delete)

	log.Fatalln(server.Run("localhost:8080"))
}
