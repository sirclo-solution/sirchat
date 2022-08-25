package apps

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirclo-solution/sirchat/logger"
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
	engineApps *gin.Engine
}

// setup use to setup application
// this method it doesn't need to be called in your app, because it's already called here
// This method includes creating a new server / initial Server
func (ths *app) setup(cfg AppConfig) error {
	if cfg.AppSecret == "" {
		logger.Get().ErrorWithoutSTT("Error setup", "Error", "invalid app secret string")
		return fmt.Errorf("client Setup(): invalid app secret string")
	}
	ths.AppSecret = cfg.AppSecret

	ths.engineApps = initServer(cfg.AppSecret)
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

// GetAuthSirclo used to get authorization SIRCLO (only use internal SIRCLO)
func GetAuthSirclo(c context.Context) (string, error) {
	sircloAuth, ok := c.Value(sircloAuth).(string)
	if !ok {
		return "", errors.New("authorization sirclo invalid")
	}
	return sircloAuth, nil
}

// BindRequestBody used to bind request body
func BindRequestBody(c context.Context, b interface{}) error {
	byteVal, ok := c.Value(reqBody).([]byte)
	if !ok {
		return errors.New("invalid request")
	}
	if err := json.Unmarshal(byteVal, b); err != nil {
		return err
	}
	return nil
}
