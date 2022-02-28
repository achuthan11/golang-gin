package main

import (
	"golearning/controller"
	"golearning/service"

	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func getVideos(ctx *gin.Context) {
	ctx.JSON(200, videoController.FindAll())
}
func putVideo(ctx *gin.Context) {
	ctx.JSON(200, videoController.Save(ctx))
}

func main() {
	server := gin.Default()

	server.GET("/videos", getVideos)

	server.POST("/videos", putVideo)

	server.Run(":8080")
}
