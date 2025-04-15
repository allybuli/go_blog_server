package utils

import (
	"io/fs"
	"os"
	"server/global"

	"gopkg.in/yaml.v3"
)

const configFile = "config.yaml"

// 从yaml配置文件中读取数据并返回字节数组
func LoadYAML() ([]byte, error) {
	return os.ReadFile(configFile)
}

// 将全局配置对象保存到yaml文件中
func SaveYAML() error {
	data, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	return os.WriteFile(configFile, data, fs.ModePerm)
}
