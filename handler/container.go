package handler

import (
	"WebDM/dto"
	"WebDM/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var containerService, _ = service.NewContainerService()

// CreateContainer 方法用于处理创建容器的 HTTP 请求
func CreateContainer(c *gin.Context) {
	containerInput := &dto.ContainerConfig{}

	c.BindJSON(containerInput)
	containerID, err := containerService.CreateContainer(containerInput)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to create container",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"containerID": containerID,
	})
}
func StartContainer(c *gin.Context) {
	containerID := c.Query("id")
	containerService.StartContainer(containerID)

	c.JSON(http.StatusCreated, gin.H{
		"containerID": containerID,
	})
}

// StopContainer 方法用于处理停止容器的 HTTP 请求
func StopContainer(c *gin.Context) {
	containerID := c.Query("id")
	err := containerService.StopContainer(containerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to stop container",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "container stopped successfully",
	})
}

// StopContainer 方法用于处理停止容器的 HTTP 请求
func StopremoveContainer(c *gin.Context) {
	containerID := c.Query("id")
	err := containerService.StopremoveContainer(containerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to stop container",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "container Stop and remove successfully",
	})
}

// ListContainers 方法用于处理查询所有容器列表的 HTTP 请求
func ListContainers(c *gin.Context) {
	containers, err := containerService.ListContainers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to list containers",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"containers": containers,
	})
}

// MonitorContainer 方法用于处理通过容器 ID 监控容器
func MonitorContainer(c *gin.Context) {
	containerID := c.Query("id")
	stats, err := containerService.MonitorContainer(containerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "failed to monitor container",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"stats": stats,
	})
}

// LogsContainer 方法用于处理通过容器 ID 输出容器日志信息的 HTTP 请求
func LogsContainer(c *gin.Context) {
	containerID := c.Query("id")
	logs, err := containerService.PrintContainerLogs(containerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Container Logs request data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"logs": logs,
	})
}

// GetContainerVersion 方法用于处理返回容器版本信息的 HTTP 请求
func GetDockerVersion(c *gin.Context) {
	versioninfo := containerService.GetDockerVersion()
	c.JSON(http.StatusOK, gin.H{
		"version": versioninfo,
	})
}

// VolumesContainer 方法用于处理通过容器 ID 返回容器卷信息的 HTTP 请求
func VolumesContainer(c *gin.Context) {
	containerID := c.Query("id")
	Volumesinfo, err := containerService.GetContainerVolumes(containerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid Volumesinfo request data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"volumes": Volumesinfo,
	})
}

func ContainerInspect(c *gin.Context) {
	containerID := c.Query("id")
	dataJson, err := containerService.ContainerInspect(containerID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid ContainerInspect request Json data",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"volumes": dataJson,
	})
}