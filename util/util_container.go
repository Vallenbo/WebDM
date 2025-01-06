package util

import (
	"WebDM/dto"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"io"
	"log"
	"os"
	"strings"
)

// 创建启动容器
func (du *DockerUtil) RunContainer(name string, image string, ports []map[string]string, envs []map[string]string) (string, error) {
	// 创建容器配置
	config := &container.Config{
		Image: image,
		Env:   convertMapToSlice(envs),
	}

	// 创建容器主机配置
	hostConfig := &container.HostConfig{
		PortBindings: convertPortBindings(ports),
	}

	// 创建容器
	resp, err := du.cli.ContainerCreate(context.Background(), config, hostConfig, nil, nil, name)
	if err != nil {
		return "", err
	}

	// 启动容器
	if err := du.cli.ContainerStart(context.Background(), resp.ID, types.ContainerStartOptions{}); err != nil {
		return "", err
	}
	return resp.ID, nil
}

// 停止容器
func (du *DockerUtil) StopContainer(containerID string) error {
	if err := du.cli.ContainerStop(context.Background(), containerID, container.StopOptions{}); err != nil {
		return err
	}
	return nil
}

func (du *DockerUtil) ContainerInspect(containerID string) (types.ContainerJSON, error) {
	dataJson, err := du.cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return types.ContainerJSON{}, err
	}

	return dataJson, nil
}

// 删除容器
func (du *DockerUtil) StopremoveContainer(containerID string) error {
	if err := du.cli.ContainerStop(context.Background(), containerID, container.StopOptions{}); err != nil {
		return err
	}
	if err := du.cli.ContainerRemove(context.Background(), containerID, types.ContainerRemoveOptions{}); err != nil {
		return err
	}
	return nil
}

func (du *DockerUtil) StartContainer(containerID string) {
	err := du.cli.ContainerStart(context.Background(), containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
}

// 查询所有容器列表
func (du *DockerUtil) ListContainers() ([]dto.Container, error) {
	// 调用 Docker API 查询所有容器列表
	containers, err := du.cli.ContainerList(context.Background(), types.ContainerListOptions{All: true})
	if err != nil {
		return nil, err
	}

	var result []dto.Container
	// 将查询到的容器信息封装到 Container 结构体中
	for _, acontainer := range containers {
		if len(acontainer.Names) <= 0 || len(acontainer.Ports) <= 0 { // 判断names、ports是否为空
			acontainer.Ports = []types.Port{{IP: "0.0.0.0", PrivatePort: 0, PublicPort: 0, Type: "no"}}
		}
		if len(acontainer.Command) <= 0 {
			acontainer.Command = ""
		}
		result = append(result, dto.Container{
			CreatedTime: acontainer.Created,
			Status:      acontainer.Status,
			ID:          acontainer.ID[:12],                           // 长 Container ID 截取为短 ID
			Name:        strings.TrimPrefix(acontainer.Names[0], "/"), // 去除 name 中的斜杠前缀
			Image:       acontainer.Image,
			Command:     acontainer.Command,
			Ports:       acontainer.Ports[0]})
	}
	return result, nil
}

// MonitorContainer 方法通过参数容器 ID 监控容器
func (du *DockerUtil) MonitorContainer(containerID string) (string, error) {
	// 创建容器监控选项
	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	}

	// 监控容器
	containerLogs, err := du.cli.ContainerLogs(context.Background(), containerID, options)
	if err != nil {
		fmt.Errorf("failed to monitor container: %v", err)
	}

	// 打印容器日志
	defer containerLogs.Close()
	bytes, err := io.ReadAll(containerLogs)
	content := string(bytes)
	for {
		_, err = io.Copy(os.Stdout, containerLogs)
		if err != nil && err != io.EOF {
			return "", fmt.Errorf("failed to read container logs: %v", err)
		}
		if err == nil || err == io.EOF {
			return content, nil
		}
	}
}

// PrintContainerLogs 方法通过参数容器 ID 输出容器日志信息
func (du *DockerUtil) PrintContainerLogs(containerID string) (string, error) {
	// 创建容器监控选项
	options := types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     true,
	}

	// 监控容器日志
	containerLogs, err := du.cli.ContainerLogs(context.Background(), containerID, options)
	if err != nil {
		return "", fmt.Errorf("failed to monitor container logs: %v", err)
	}

	// 打印容器日志
	defer containerLogs.Close()
	reader, err := io.ReadAll(containerLogs)
	if err == io.EOF {
		return "nil", nil
	}
	if err != nil {
		return err.Error(), nil
	}
	content := string(reader)
	return content, nil
}

// GetContainerVersion 方法返回容器信息
func (du *DockerUtil) GetDockerVersion() types.Version {
	versioninfo, err := du.cli.ServerVersion(context.Background())
	if err != nil {
		log.Println("GetContainerVersion error :", err)
		panic(err)
	}

	return versioninfo
}

// GetContainerVolumes 方法通过参数容器 ID 返回容器卷信息
func (du *DockerUtil) GetContainerVolumes(containerID string) ([]string, error) {
	// 获取容器信息
	containerinfo, err := du.cli.ContainerInspect(context.Background(), containerID)
	if err != nil {
		return nil, fmt.Errorf("failed to inspect container: %v", err)
	}

	// 获取容器卷信息
	var volumes []string
	for _, mount := range containerinfo.Mounts {
		if mount.Type == "volume" {
			volumes = append(volumes, mount.Name)
		}
	}

	return volumes, nil
}