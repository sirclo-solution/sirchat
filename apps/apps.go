package apps

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

// AppServerConfig as a parameter of a Start()
type AppServerConfig struct {
	// Port is specifies the TCP address for the server to listen on,
	// example (port = 8080)
	Port string
	// timeout set to second time
	// default is 30 second
	Timeout int
}

// Interface App is contract from app struct
// this wrap method for needed your application
type App interface {
	GetAppSecret() string

	// Command is method provide grouping route /command
	// this method have 2 parameter
	// first parameter is commandName (ex. /searchProduct)
	// second parameter is handler, this parameter using type HandlerCommand func(*gin.Context) (interface{}, error)
	Command(commandName string, handler HandlerCommand)

	// Start use start your service/application
	Start(param AppServerConfig)
}

// AppConfig is parameter for initial your Apps
type AppConfig struct {
	// AppSecret is secret for verifying request
	AppSecret string `json:"app_secret"`
}

// struct app owned by App interface
type app struct {
	// AppSecret is secret for verifying request
	AppSecret string `json:"app_secret"`

	// EngineApps used for configuring your application with the server
	EngineApps *gin.Engine
}

// setup use to setup application
// this method it doesn't need to be called in your app, because it's already called here
// This method includes creating a new server / initial Server
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

// NewApps is method for initial your application
func NewApps(cfg AppConfig) App {
	var c app
	c.setup(cfg)
	return &c
}

func GetAuthSirclo(c *gin.Context) string {
	return c.GetString(SircloAuthorization)
}

func BindRequestBody(c *gin.Context, b any) error {
	var byteVal []byte
	if val, ok := c.Get(SirchatRequestBody); ok && val != nil {
		byteVal, _ = val.([]byte)
	}

	if err := json.Unmarshal(byteVal, b); err != nil {
		return err
	}
	return nil
}
