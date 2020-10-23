package logging

type LoggerConfig struct {
	Console Logger
	File    Logger
}

type Logger struct {
	Enable bool   `yaml:"enable"`
	JSON   bool   `yaml:"json"`
	Level  string `yaml:"level"`
	Path   string `yaml:"path"`
}
