package conf

import (
	"WebDM/util"
	"fmt"
	"github.com/docker/docker/client"
	"gopkg.in/ini.v1"
)

type ContainerService struct {
	dockerUtil *util.DockerUtil
}

func DockerCon() (*ContainerService, error) {
	// 从配置文件中获取地址和认证信息
	cfg, err := ini.Load("docker_conf.ini")
	if err != nil {
		fmt.Printf("Failed to read docker_conf.ini file: %v\n", err)
		return nil, err
	}

	ip := cfg.Section("docker").Key("ip").String()
	port := cfg.Section("docker").Key("port").String()

	host := fmt.Sprintf("tcp://%s:%s", ip, port)
	// 使用认证信息创建Docker客户端
	_, err = client.NewClientWithOpts(client.WithHost(host))
	if err != nil {
		fmt.Printf("Failed to create Docker client: %v\n", err)
		return nil, err
	}
	println(host, "connection success! ")
	return &ContainerService{}, nil
}