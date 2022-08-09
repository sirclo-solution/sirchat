package models

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/sirclo-solution/sirchat/apps"
)

type MessageComponentType string

const (
	MCTDialog       MessageComponentType = "dialog"
	MCTDrawer       MessageComponentType = "drawer"
	MCTNotification MessageComponentType = "notification"
	MCTMessage      MessageComponentType = "message"
)

type IComponent interface {
	GetType() MessageComponentType
	Validate() (bool, []error)
	Compose() ([]byte, error)
	Send() (interface{}, error)
}

// `component` struct type is meant to be embedded to a component
// subtype. `component` provides the embedding structs with field `Type`
// and the basic methods for a component.
type component struct {
	IComponent `json:"-"`
	Type       MessageComponentType `json:"type"`
}

// GetType returns the type of the component. Use this method as the
// alternative for getting the value of field `Type` conventionally,
// such as when handling structs as IComponent.
func (ths *component) GetType() MessageComponentType {
	return ths.Type
}

// Compose returns the JSON object representation (`[]byte`) of the
// embedding struct. It calls the Validate function in the embedding
// component/block and all of the nested blocks it may have.
func (ths *component) Compose() ([]byte, error) {
	var errs []error

	if valid, err := ths.Validate(); !valid {
		errs = append(errs, err...)
	}

	if len(errs) > 0 {
		log.Printf("component.Compose() %+q\n", errs)
		return nil, apps.NewAppsError(http.StatusBadRequest, errors.New("invalid component/blocks"), "invalid component/blocks")
	}

	res, err := json.Marshal(ths.IComponent)
	if err != nil {
		return nil, apps.NewAppsError(http.StatusInternalServerError, err, "error when marshaling component")
	}

	return res, nil
}

// Send returns the JSON string representation of the embedding
// struct. It calls the component's Compose function in the process.
func (ths *component) Send() (interface{}, error) {
	result, err := ths.Compose()
	if err != nil {
		return nil, err
	}

	return string(result), nil
}
