package util

import (
	"WebDM/dto"
	"archive/tar"
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"gopkg.in/ini.v1"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"
)

// 定义 DockerUtil 结构体，用于封装所有 Docker 相关操作
type DockerUtil struct {
	cli *client.Client
}

// 实例化 DockerUtil 结构体
func NewDockerUtil() (*DockerUtil, error) {
	// 从配置文件中获取地址和认证信息
	cfg, err := ini.Load("conf/docker_conf.ini")
	if err != nil {
		fmt.Printf("Failed to read docker_conf.ini file: %v\n", err)
		return nil, err
	}
	ip := cfg.Section("docker").Key("ip").String()
	port := cfg.Section("docker").Key("port").String()

	host := fmt.Sprintf("tcp://%s:%s", ip, port)
	// 连接 Docker 客户端
	cli, err := client.NewClientWithOpts(
		client.WithHost(host),
		client.WithHTTPClient(&http.Client{Transport: http.DefaultTransport}),
		client.WithAPIVersionNegotiation(),
	)
	if err != nil {
		return nil, err
	}

	return &DockerUtil{cli: cli}, nil
}

// 查询所有镜像列表
func (du *DockerUtil) ListImages() ([]dto.Image, error) {
	// 调用 Docker API 查询所有镜像列表
	images, err := du.cli.ImageList(context.Background(),
		types.ImageListOptions{})
	if err != nil {
		return nil, err
	}

	var result []dto.Image
	for _, image := range images { // 将查询到的镜像信息封装到 Image 结构体中
		result = append(result, dto.Image{Size: image.Size / 1024 / 1024, CreatedTime: image.Created, Name: image.ID, Tag: image.RepoTags[0]})
	}
	return result, nil
}

// 创建镜像
//
//	func (du *DockerUtil) CreatefromImage(name string, dockerfileContent string)  {
//		//containerID := "my-container"
//		//commitOpts := types.ContainerCommitOptions{
//		//	Reference: "my-image:latest",
//		//}
//		//resp, err := cli.ContainerCommit(ctx, containerID, commitOpts)
//		//if err != nil {
//		//	panic(err)
//		//}
//		createTarGz("")
//	}

func createTarGz(source, target string) error {
	// 打开源文件
	file, err := os.Open(source)
	if err != nil {
		return err
	}
	defer file.Close()

	// 创建目标文件
	targetFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer targetFile.Close()

	// 创建gzip压缩器
	gzWriter := gzip.NewWriter(targetFile)
	defer gzWriter.Close()

	// 创建tar打包器
	tarWriter := tar.NewWriter(gzWriter)
	defer tarWriter.Close()

	// 获取源文件的信息
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// 创建tar头部信息
	header := &tar.Header{
		Name: filepath.Base(source),
		Mode: int64(fileInfo.Mode()),
		Size: fileInfo.Size(),
	}

	// 写入tar头部信息
	err = tarWriter.WriteHeader(header)
	if err != nil {
		return err
	}

	// 将源文件内容写入tar包
	_, err = io.Copy(tarWriter, file)
	if err != nil {
		return err
	}

	return nil
}
func (du *DockerUtil) CreatefromImage(imagename, tag string, dockerfileContent string) error {
	// 创建 Dockerfile 临时文件
	path := fmt.Sprintf("tmp/Dockerfile")
	gzPath := fmt.Sprintf("tmp/Dockerfile.tar.gz")
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	file.WriteString(dockerfileContent)

	// 构建镜像选项
	//reader, err := createTarReader(dockerfileContent)
	//tarF, _ := io.ReadAll(dockerfileContent)
	//tarF := bytes.NewBuffer([]byte(dockerfileContent)) // archive.TarWithOptions("./tmp", &archive.TarOptions{})
	//if err != nil {
	//	panic(err)
	//}
	createTarGz(path, gzPath)

	tarF, err := os.Open(gzPath)
	if err != nil {
		panic(err)
	}

	//defer tarF.Close()

	buildOptions := types.ImageBuildOptions{
		Dockerfile: "Dockerfile",
		Context:    tarF,
		Tags:       []string{imagename + ":" + tag},
	}
	// 调用 Docker API 构建镜像
	buildResponse, err := du.cli.ImageBuild(context.Background(), tarF, buildOptions)
	if err != nil {
		return err
	}

	// 解析构建镜像的输出
	defer buildResponse.Body.Close()
	outputBytes, err := io.ReadAll(buildResponse.Body)
	if err != nil {
		return err
	}
	output := string(outputBytes)
	if strings.Contains(output, "error") {
		return fmt.Errorf("error occurred while building image: %s%v", output)
	}

	return nil
}

func createTarReader(dockerfileContent string) (*tar.Reader, error) {
	tarBuffer := new(bytes.Buffer)
	tarWriter := tar.NewWriter(tarBuffer)

	// 写入 Dockerfile。
	header := &tar.Header{
		Name:    "Dockerfile",
		Size:    int64(len(dockerfileContent)),
		Mode:    0644,
		ModTime: time.Now(),
	}
	err := tarWriter.WriteHeader(header)
	if err != nil {
		return nil, err
	}
	_, err = tarWriter.Write([]byte(dockerfileContent))
	if err != nil {
		return nil, err
	}

	// 写入上下文文件。
	err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.Mode().IsRegular() {
			return nil
		}
		header, err = tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}
		header.Name = path
		err = tarWriter.WriteHeader(header)
		if err != nil {
			return err
		}
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(tarWriter, file)
		return err
	})
	if err != nil {
		return nil, err
	}

	err = tarWriter.Close()
	if err != nil {
		return nil, err
	}

	return tar.NewReader(tarBuffer), nil
}

// 辅助函数：将端口映射字符串转换为 Docker API 所支持的格式
func convertPortBindings(ports []map[string]string) nat.PortMap {
	result := make(nat.PortMap)
	for _, kv := range ports {
		containerPort, _ := kv["containerPort"]
		hostPort := kv["hostPort"]
		result[nat.Port(containerPort+"/tcp")] = []nat.PortBinding{
			{
				HostPort: hostPort,
			},
		}
	}

	return result
}

// 辅助函数：将 map[string]string 转换为 []string
func convertMapToSlice(m []map[string]string) []string {
	var result []string
	for _, kv := range m {
		result = append(result, fmt.Sprintf("%s=%s", kv["key"], kv["value"]))
	}

	return result
}

// BuildImage 方法使用容器创建镜像
func (du *DockerUtil) BuildImage(containerID string) error {
	resp, err := du.cli.ContainerCommit(context.Background(), containerID, types.ContainerCommitOptions{})
	if err != nil {
		fmt.Printf("BuildImage err : %v\n", err)
		return err
	}
	fmt.Printf("image ID is：%v \n", resp)
	return nil
}

// Deleteimages  方法删除镜像
func (du *DockerUtil) Deleteimage(imageID string) (bool, error) {
	resp, err := du.cli.ImageRemove(context.Background(), imageID, types.ImageRemoveOptions{})
	if err != nil {
		fmt.Printf("Remove Image err :%v", err)
		return false, err
	}
	fmt.Printf("image ID is： %v\n", resp)
	return true, nil
}

// SearchImage 方法通过参数镜像 ID 搜索镜像
func (du *DockerUtil) SearchImage(imageName string) ([]dto.Imagepull, error) {
	imagesResults, err := du.cli.ImageSearch(context.Background(), imageName, types.ImageSearchOptions{Limit: 20}) // 搜索结果数量限制为10个
	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to send HTTP request: %v", err)
	}
	var result []dto.Imagepull
	for _, image := range imagesResults { // 将查询到的镜像信息封装到 Image 结构体中
		result = append(result, dto.Imagepull{
			Name:        image.Name,
			Description: image.Description,
			StarCount:   image.StarCount,
			IsOfficial:  image.IsOfficial,
			IsAutomated: image.IsAutomated,
		})
	}

	return result, nil
}

func (du *DockerUtil) PullImage(imageID string) (io.ReadCloser, error) {
	// 创建并启动容器
	resp, err := du.cli.ImagePull(context.Background(), imageID, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	return resp, nil
}