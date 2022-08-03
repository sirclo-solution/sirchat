package models

type MessageComponentType string

const (
	MCTDialog       MessageComponentType = "dialog"
	MCTDrawer       MessageComponentType = "drawer"
	MCTNotification MessageComponentType = "notification"
	MCTMessage      MessageComponentType = "message"
)

type IComponent interface {
	GetType() MessageComponentType
	Validate() (bool, error)
	Compose() ([]byte, []error)
}

type Component struct {
	Type MessageComponentType `json:"type"`
}

func (ths *Component) GetType() MessageComponentType {
	return ths.Type
}
