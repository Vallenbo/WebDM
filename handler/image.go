package handler

import (
	"WebDM/dto"
	"WebDM/service"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

var imageService, _ = service.NewImageService()

// CreateImage 方法用于处理创建镜像的 HTTP 请求
func CreateImage(c *gin.Context) {
	input := &dto.ImageInput{}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid request data",
			"error":   err.Error(),
		})
		return
	}

	_ = imageService.CreateImage(input)
	c.JSON(http.StatusCreated, gin.H{
		"image": "ok",
	})
}

// 返回所有镜像
func GetAllImages(c *gin.Context) {
	images, err := imageService.ListImages()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to get images",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"images": images,
	})
}

// DeleteImage 方法用于处理删除镜像的 HTTP 请求
func DeleteImage(c *gin.Context) {
	imageidID := c.Query("id")
	_, err := imageService.Deleteimage(imageidID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Delete Image request data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "image deleted successfully",
	})
}

// SearchImage 方法用于处理搜索镜像的 HTTP 请求
func SearchImage(c *gin.Context) {
	imageName := c.Query("imageName")
	images, err := imageService.SearchImage(imageName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Search Image request data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": images,
	})
}

// pullImage 方法用于拉取镜像
func PullImage(c *gin.Context) {
	imageName := c.Query("id")
	image, err := imageService.PullImage(imageName)
	written, err := io.Copy(c.Writer, image)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Search Image request data",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": written,
	})
}
