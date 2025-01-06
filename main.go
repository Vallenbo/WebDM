package main

import (
	"WebDM/handler"
	"WebDM/util"
	"fmt"
	"github.com/docker/docker/pkg/stack"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	log.Println("Web Server starting on port 8080")
	r := gin.Default() // 创建一个默认的路由引擎
	r.Use(func(context *gin.Context) {
		defer func() {
			err := recover()
			switch err.(type) {
			case error:
				err := err.(error)
				fmt.Printf(err.Error())
				stack.Dump()
				context.AbortWithStatusJSON(500, gin.H{})
			}
		}()
		context.Next()
	})
	r.GET("/monitor", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Failed to upgrade to websocket:", err)
			return
		}
		defer conn.Close()

		tick := time.Tick(time.Second * 3)
		for range tick {
			info := util.GetSystemInfo()
			err = conn.WriteJSON(info)
			if err != nil {
				fmt.Println(err.Error())
				stack.Dump()
				return
			}
		}
	})

	//r.LoadHTMLGlob("templates/*")
	//r.POST("/", controller.IndexHandlc)
	r.POST("/CreateContainer", handler.CreateContainer)
	r.POST("/RunContainer", handler.StartContainer)
	r.POST("/StopContainer", handler.StopContainer)
	r.POST("/StopremoveContainer", handler.StopremoveContainer)
	r.POST("/ListContainers", handler.ListContainers)
	r.POST("/MonitorContainer", handler.MonitorContainer)
	r.POST("/LogsContainer", handler.LogsContainer)
	r.POST("/GetDockerVersion", handler.GetDockerVersion)
	r.POST("/VolumesContainer", handler.VolumesContainer)
	r.POST("/ContainerInspect", handler.ContainerInspect)

	r.POST("/CreateImage", handler.CreateImage)
	r.POST("/dockerfilereturn", handler.Dockerfilereturn)
	r.POST("/GetAllImages", handler.GetAllImages)
	r.POST("/DeleteImage", handler.DeleteImage)
	r.POST("/PullImage", handler.PullImage)
	r.POST("/SearchImage", handler.SearchImage)
	//r.POST("/pullImage", handler.PullImage)

	err := r.Run(":8080")
	if err != nil {
		fmt.Println("Background-Server start failed error:", err)
	}
}