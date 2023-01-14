package config

import (
	_ "embed"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
	"os"
)

//go:embed config.yaml
var CacheFile []byte

// RunFileConfig 设置配置信息
func (cfg *Config) RunFileConfig() {
	if err := yaml.Unmarshal(cfg.FileCache, &cfg.StructConfig); err != nil {
		logrus.Error("config file read yaml error:", err)
		panic(err)
	}
}

// FileRead 读取文件内容，如果是使用本地文件
// config.yaml 应当存放在和项目相同的路径
func (cfg *Config) FileRead() {
	data, err := os.ReadFile("config.yaml")
	if err != nil {
		logrus.Error("config file read local error:", err)
		panic(err)
	}
	cfg.FileCache = data
}
