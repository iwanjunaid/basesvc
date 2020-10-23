package server

// ServerList :
type ServerList struct {
	Rest Server // main REST config for this service
}

// Server :
type Server struct {
	TLS     bool   `yaml:"tls"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	Timeout int    `yaml:"timeout"`
}
