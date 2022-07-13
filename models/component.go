package models

type MessageComponentType string

const (
	MCTDialog MessageComponentType = "dialog"
	MCTDrawer MessageComponentType = "drawer"
)

type IComponent interface {
	NewDialog() *Component
	NewDrawer() *Component
}

type Component struct {
	Type   MessageComponentType
	Title  Title
	Action Action
	Blocks Blocks
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
