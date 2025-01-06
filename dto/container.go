package dto // 定义 Container 结构体，包含 Docker 容器的所有信息
import "github.com/docker/docker/api/types"

type Container struct {
	ID          string     `json:"id"`          // 容器 ID
	Name        string     `json:"name"`        // 容器名称
	Image       string     `json:"image"`       // 容器使用的镜像名称
	CreatedTime int64      `json:"createdTime"` // 容器创建时间
	Status      string     `json:"status"`      // 容器状态
	Ports       types.Port `json:"ports,omitempty"`
	PublicPort  uint16     `json:"publicPort,omitempty"`
	PrivatePort uint16     `json:"privatePort,omitempty"`
	Labels      string     `json:"labels"`  // 容器标签信息
	Command     string     `json:"command"` //容器内运行命令
}

type ContainerConfig struct {
	Name    string              `json:"name"`
	Image   string              `json:"image"`
	Env     []map[string]string `json:"env"`
	PortMap []map[string]string `json:"portMap"`
}