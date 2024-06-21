package fileutil

import (
	"github.com/xxl6097/go-glog/glog"
	"path/filepath"
	"strings"
)

func GetFileNameAndFileExtension(filePath string) (string, string) {
	// 使用 filepath 包提供的函数获取文件名
	fileName := filepath.Base(filePath)
	// 使用 strings 包提供的函数获取文件名和扩展名
	fileNameWithoutExtension := strings.TrimSuffix(fileName, filepath.Ext(fileName))
	fileExtension := filepath.Ext(fileName)
	// 打印文件名和扩展名
	glog.Println("文件名:", fileNameWithoutExtension)
	glog.Println("扩展名:", fileExtension)
	return fileNameWithoutExtension, fileExtension
}
