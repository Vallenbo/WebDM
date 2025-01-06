package service

import (
	"WebDM/dto"
	"WebDM/util"
	"fmt"
	"io"
	"os"
)

type ImageService struct {
	dockerUtil *util.DockerUtil
}

func NewImageService() (*ImageService, error) {
	dockerUtil, err := util.NewDockerUtil()
	if err != nil {
		return nil, err
	}
	return &ImageService{dockerUtil: dockerUtil}, nil
}

func (s *ImageService) ListImages() ([]dto.Image, error) {
	return s.dockerUtil.ListImages()
}

func (s *ImageService) CreateImage(input *dto.ImageInput) error {
	// 读取镜像模板文件
	templatePath := fmt.Sprintf("template/%s/Dockerfile", input.TemplateName)
	templateContent, err := os.ReadFile(templatePath)
	if err != nil {
		panic(err)
	}

	// 添加自定义命令
	dockerfileContent := string(templateContent)
	for _, s := range input.TemplateCommand {
		dockerfileContent += fmt.Sprintf("\nRUN %s", s)
	}

	// 创建新的 Docker 镜像
	err = s.dockerUtil.CreatefromImage(input.Name, input.Tag, dockerfileContent)
	if err != nil {
		panic(err)
	}

	return nil
}

func (s *ImageService) SearchImage(imageName string) ([]dto.Imagepull, error) {
	result, err := s.dockerUtil.SearchImage(imageName)
	if err != nil {
		return nil, fmt.Errorf("failed to Monitor Container: %v", err)
	}
	return result, nil
}

func (s *ImageService) PullImage(imageID string) (io.ReadCloser, error) {
	result, err := s.dockerUtil.PullImage(imageID)
	if err != nil {
		return nil, fmt.Errorf("failed to Monitor Container: %v", err)
	}
	return result, nil
}

func (s *ImageService) Deleteimage(imageID string) (bool, error) {
	str, err := s.dockerUtil.Deleteimage(imageID)
	if err != nil {
		return false, fmt.Errorf("failed to Monitor Container: %v", err)
	}
	return str, nil
}
