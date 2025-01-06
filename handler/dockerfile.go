package handler

import (
	"WebDM/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Dockerfilereturn(c *gin.Context) {
	dirs, err := service.FindDockerDirs("template")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get images",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": dirs,
	})
}
