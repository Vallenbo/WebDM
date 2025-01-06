package main

import (
	"WebDM/service"
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	dirs, err := service.FindDockerDirs("/path/to/directory")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(dirs)
}
