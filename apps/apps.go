package apps

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type AppServerConfig struct {
	// Port is specifies the TCP address for the server to listen on,
	// example (port = 8080)
	Port string
	// timeout set to second time
	// default is 30 second
	Timeout int
}

type App interface {
	GetAppSecret() string
	Command(commandName string, handler HandlerCommand)
	Start(param AppServerConfig)
}

type AppConfig struct {
	AppSecret string `json:"app_secret"`
}

type app struct {
	AppSecret  string `json:"app_secret"`
	EngineApps *gin.Engine
}

func (ths *app) setup(cfg AppConfig) error {
	if cfg.AppSecret == "" {
		return fmt.Errorf("client Setup(): invalid app secret string")
	}
	ths.AppSecret = cfg.AppSecret

	ths.EngineApps = InitServer(cfg.AppSecret)
	return nil
}

func (ths *app) GetAppSecret() string {
	return ths.AppSecret
}

func NewApps(cfg AppConfig) App {
	var c app
	c.setup(cfg)
	return &c
}
