package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func GetFilesHandler(path string, directory string) (int, map[string]string, []byte) {
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

func PostFilesHandler(path string, directory string, fileContent string)(int, map[string]string, []byte) {
	fileName := strings.TrimPrefix(path, "/files/")
	filePath := filepath.Join(directory, fileName)

	data := []byte(fileContent)

	err := os.WriteFile(filePath, data, 0644)
	if err != nil {
		return 500, nil, []byte("Internal Server Error")
	}

	responseHeaders := map[string]string {
		"Content-Type": "text/plain",
		"Content-Length": fmt.Sprintf("%d", len(fileContent)),
	}

	return 201, responseHeaders, data
}

func FilesHandler(path string, directory string, method string, fileContent string)(int, map[string]string, []byte) {
	if method == "GET" {
		return GetFilesHandler(path, directory)
	}

	if method == "POST" {
		return PostFilesHandler(path, directory, fileContent)
	}

	return 405, nil, []byte{}
}