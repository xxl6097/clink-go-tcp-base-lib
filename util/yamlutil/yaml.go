package yamlutil

import (
	"github.com/xxl6097/go-glog/glog"
	"gopkg.in/yaml.v3"
	"os"
)

func Decode[T any](path string) *T {
	var file *os.File
	var err error
	if path == "" {
		glog.Info("read same dir config file")
		file, err = os.Open("app.yaml")
	} else {
		file, err = os.Open(path)
	}
	defer file.Close()
	if err != nil {
		glog.Error("Error opening file ", err)
		return nil
	}
	// 创建解析器
	decoder := yaml.NewDecoder(file)
	var t *T
	// 解析 YAML 数据
	err = decoder.Decode(t)
	if err != nil {
		glog.Error("Error decoding YAML", err)
	} else {
		t = nil
	}
	return t
}

func Unmarshal[T any](content string) *T {
	var t *T
	err := yaml.Unmarshal([]byte(content), t)
	if err != nil {
		glog.Error(err)
		return nil
	}
	return t
}
