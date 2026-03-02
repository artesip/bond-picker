package config

type ServerConfig struct {
	Port int `yaml:"port"`
}

type JWTConfig struct {
	Path string `yaml:"key_path"`
}

type DatabaseConfig struct {
	Url string `yaml:"url"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	JWT      JWTConfig      `yaml:"jwt"`
	Database DatabaseConfig `yaml:"database"`
}
