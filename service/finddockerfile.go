package service

import (
	"os"
	"path/filepath"
)

func FindDockerDirs(root string) ([]string, error) {
	var dockerDirs []string
	// 读取指定目录下的所有文件和目录
	files, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	// 检查每个目录是否包含Dockerfile
	for _, file := range files {
		if file.IsDir() {
			dirPath := filepath.Join(root, file.Name())
			if _, err := os.ReadFile(filepath.Join(dirPath, "Dockerfile")); err == nil {
				dockerDirs = append(dockerDirs, file.Name())
			}
		}
	}

	return dockerDirs, nil
}