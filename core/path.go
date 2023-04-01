package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func InitPath() string {
	rootDir, _ := os.Getwd()
	filePath := filepath.Join(rootDir, "files")
	fmt.Println(filePath)
	return filePath
}
