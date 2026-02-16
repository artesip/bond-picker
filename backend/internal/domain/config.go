package domain

type Server struct {
	Port uint16 `yaml:"port"`
}

type JWT struct {
	Path string `yaml:"key_path"`
}

type Config struct {
	Server Server `yaml:"server"`
	JWT    JWT    `yaml:"jwt"`
}
