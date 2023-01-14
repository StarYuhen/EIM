package config

// 用于储存配置文件的信息

type Config struct {
	FileCache []byte
	StructConfig
}

type Function interface {
	RunFileConfig()
	FileRead()
}

type StructConfig struct {
	DataBase DataBase `yaml:"DataBase"`
}

type DataBase struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Schema   string `yaml:"schema"`
	Addr     string `yaml:"addr"`
}
