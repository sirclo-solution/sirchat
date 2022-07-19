package models

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MessageComponentType string

const (
	MCTDialog MessageComponentType = "dialog"
	MCTDrawer MessageComponentType = "drawer"
)

type IComponent interface {
	NewDialog() *Component
	NewDrawer() *Component
	Compose() ([]byte, []error)
}

type Component struct {
	Type   MessageComponentType
	Title  Title
	Action Action
	Blocks []IBlock
}

func NewApp() IComponent {
	return &Component{}
}

func (c *Component) NewDialog() *Component {
	return &Component{
		Type: MCTDialog,
	}
}

func (c *Component) NewDrawer() *Component {
	return &Component{
		Type: MCTDrawer,
	}
}

func (c *Component) Compose() ([]byte, []error) {
	var errs []error

	for i, v := range c.Blocks {
		if err := validateBlock(v); err != nil {
			err := fmt.Errorf("component.Blocks index %d: %s", i, err.Error())
			errs = append(errs, err)
		}
	}

	if len(errs) > 0 {
		return nil, errs
	}

	res, err := json.Marshal(c)
	if err != nil {
		return nil, []error{errors.New("error when marshaling block")}
	}

	return res, nil
}