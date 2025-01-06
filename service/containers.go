package service

import (
	"WebDM/dto"
	"WebDM/util"
	"fmt"
	"github.com/docker/docker/api/types"
)

type ContainerService struct {
	dockerUtil *util.DockerUtil
}

func NewContainerService() (*ContainerService, error) {
	dockerUtil, err := util.NewDockerUtil()
	if err != nil {
		return nil, err
	}
	return &ContainerService{dockerUtil: dockerUtil}, nil
}

func (s *ContainerService) ListContainers() ([]dto.Container, error) {
	return s.dockerUtil.ListContainers()
}

func (s *ContainerService) CreateContainer(config *dto.ContainerConfig) (string, error) {
	// 创建并启动容器
	ID, err := s.dockerUtil.RunContainer(config.Name, config.Image, config.PortMap, config.Env)
	if err != nil {
		panic(err)
	}
	return ID, nil
}

// 停止容器
func (s *ContainerService) StopContainer(id string) error {
	if err := s.dockerUtil.StopContainer(id); err != nil {
		return fmt.Errorf("failed to stop container: %v", err)
	}
	return nil
}

func (s *ContainerService) ContainerInspect(id string) (types.ContainerJSON, error) {
	dataJson, err := s.dockerUtil.ContainerInspect(id)
	if err != nil {
		return types.ContainerJSON{}, fmt.Errorf("failed to stop container: %v", err)
	}
	return dataJson, nil
}

func (s *ContainerService) StartContainer(id string) {
	s.dockerUtil.StartContainer(id)
}

// 停止并删除容器
func (s *ContainerService) StopremoveContainer(id string) error {
	if err := s.dockerUtil.StopContainer(id); err != nil {
		return fmt.Errorf("failed to stop container: %v", err)
	}
	if err := s.dockerUtil.StopremoveContainer(id); err != nil {
		return fmt.Errorf("failed to stop and remove container: %v", err)
	}
	return nil
}

// MonitorContainer 方法通过参数容器 ID 监控容器
func (s *ContainerService) MonitorContainer(id string) (string, error) {
	str, err := s.dockerUtil.MonitorContainer(id)
	if err != nil {
		return "", fmt.Errorf("failed to Monitor Container: %v", err)
	}
	return str, nil
}

func (s *ContainerService) PrintContainerLogs(id string) (string, error) {
	str, err := s.dockerUtil.PrintContainerLogs(id)
	if err != nil {
		return "", fmt.Errorf("failed to Monitor Container: %v", err)
	}
	return str, nil
}
func (s *ContainerService) GetDockerVersion() types.Version {
	versioninfo := s.dockerUtil.GetDockerVersion()
	return versioninfo
}

func (s *ContainerService) GetContainerVolumes(id string) ([]string, error) {
	volumesinfo, err := s.dockerUtil.GetContainerVolumes(id)
	if err != nil {
		return nil, fmt.Errorf("failed to Monitor Container: %v", err)
	}
	return volumesinfo, nil
}