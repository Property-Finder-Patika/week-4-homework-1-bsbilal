package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	searchDir := "c:/CactusFlutter"
	fileList := []string{}
	err := filepath.Walk(searchDir, func(path string, f os.FileInfo, err error) error {
		fileList = append(fileList, path)
		return nil
	})
	if err != nil {
		// may results with permission errors
		fmt.Println(err)
		return
	}
	for _, file := range fileList {

		fmt.Println(file) //folders
	}
}
