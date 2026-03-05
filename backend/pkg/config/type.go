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

type Cron struct {
	Str string `yaml:"str"`
}

type Config struct {
	Server   ServerConfig   `yaml:"server"`
	JWT      JWTConfig      `yaml:"jwt"`
	Database DatabaseConfig `yaml:"database"`
	Cron     Cron           `yaml:"cron"`
}
