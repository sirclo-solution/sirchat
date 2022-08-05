package modules

import (
	"fmt"

	"github.com/sirclo-solution/sirchat/models"
)

type Client interface {
	GetAppSecret() string
	Command(commandName string, handler func())
	Send(response models.IComponent) error
}

type ClientConfig struct {
	AppSecret string `json:"app_secret"`
}

type client struct {
	AppSecret string `json:"app_secret"`
}

func (ths *client) setup(cfg ClientConfig) error {
	if cfg.AppSecret == "" {
		return fmt.Errorf("client Setup(): invalid app secret string")
	}
	ths.AppSecret = cfg.AppSecret

	return nil
}

func (ths *client) GetAppSecret() string {
	return ths.AppSecret
}

func (ths *client) Command(commandName string, handler func()) {

}

func (ths *client) Send(response models.IComponent) error {
	/*replace _ with a variable. e.x. jsonStr*/ _, errs := response.Compose()
	if len(errs) != 0 {
		return fmt.Errorf("client Send(): %+q", errs)
	}

	// send jsonStr to BE via http server

	return nil
}

func NewClient(cfg ClientConfig) Client {
	var c client
	c.setup(cfg)
	return &c
}
