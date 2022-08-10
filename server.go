package main

import (
	"gin_chim/controllers"
	"gin_chim/middlewares"
	"gin_chim/mysql"
	"gin_chim/service"

	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
)

var (
	videoService    service.VideoService        = service.New()
	VideoController controllers.VideoController = controllers.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Static("/css", "./templates/css")

	server.LoadHTMLGlob("templates/*.html")

	server.Use(gin.Recovery(), middlewares.Logger(),
		gindump.Dump(), mysql.Cors())
	//middlewares.BasicAuth()
	apiRoutes := server.Group("/api")
	{
		apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "OK!!",
			})
		})

		apiRoutes.GET("/videos", func(ctx *gin.Context) {
			ctx.JSON(200, VideoController.FindAll())
		})
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := VideoController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Video input is valid!"})
			}
		})
		apiRoutes.GET("/users/:id", controllers.GetUserDetail)
		apiRoutes.GET("/users/", controllers.GetUser)
		apiRoutes.POST("/login/", controllers.Login)
		apiRoutes.PUT("/users/:id", controllers.UpdateUser)
		apiRoutes.POST("/users", controllers.PostUser)
	}
	viewRoutes := server.Group("/view")
	{
		viewRoutes.GET("/videos", VideoController.ShowAll)
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
