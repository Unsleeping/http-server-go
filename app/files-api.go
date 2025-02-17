package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func FilesHandler(path string, directory string)(int, map[string]string, []byte) {
	fileName := strings.TrimPrefix(path, "/files/")
	filePath := filepath.Join(directory, fileName)
	fileInfo, err := os.Stat(filePath)

	if err != nil {
		if os.IsNotExist(err) {
			return 404, nil, []byte{}
		} 

		fmt.Println("Error getting file info:", err)
		return 500, nil, []byte{}
	}

	fileSize := fileInfo.Size()

	responseHeaders := map[string]string {
		"Content-Type": "application/octet-stream",
		"Content-Length": fmt.Sprintf("%d", fileSize),
	}

	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return 500, nil, []byte{}
	}

	return 200, responseHeaders, fileContent
}