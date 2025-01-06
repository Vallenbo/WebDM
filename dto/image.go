package dto

// 定义 Image 结构体，包含 Docker 镜像的所有信息
type Image struct {
	Name         string `json:"name"`         // 镜像名称
	Tag          string `json:"tag"`          // 镜像标签
	Digest       string `json:"digest"`       // 镜像摘要
	Size         int64  `json:"size"`         // 镜像大小
	CreatedTime  int64  `json:"createdTime"`  // 镜像创建时间
	Author       string `json:"author"`       // 镜像作者
	Architecture string `json:"architecture"` // 镜像所支持的CPU架构
	OS           string `json:"os"`           // 镜像所支持的操作系统

}

type Imagepull struct {
	Name        string `json:"name"`
	StarCount   int    `json:"star_count"`
	IsOfficial  bool   `json:"is_official"`
	IsAutomated bool   `json:"is_automated"`
	Description string `json:"description"`
}
