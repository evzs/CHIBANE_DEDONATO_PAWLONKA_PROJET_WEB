package web_project

import (
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"os"
)

// Base config structure

type TomlConf struct {
	WebServer *WebServer `toml:"web_server"`
}

// Server part

type WebServer struct {
	Config *Config `toml:"config"`
}

type Config struct {
	IP   string `toml:"ip"`
	Port string `toml:"port"`
}

func NewConfig() *TomlConf {
	conf := new(TomlConf)

	content, _ := os.ReadFile("back/config.toml")
	err := toml.Unmarshal(content, conf)
	if err != nil {
		fmt.Println(err)
	}

	return conf
}
