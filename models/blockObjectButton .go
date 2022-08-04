package models

type ButtonBlockObjectColor string
type ButtonBlockObjectVariant string
type ButtonBlockObjectType string

const (
	ButtonBlockObjectColorPrimary   ButtonBlockObjectColor   = "primary"
	ButtonBlockObjectColorSecondary ButtonBlockObjectColor   = "secondary"
	ButtonObjectVariantContained    ButtonBlockObjectVariant = "contained"
	MBTTAction                      ButtonBlockObjectType    = "action"
	MBTTCancel                      ButtonBlockObjectType    = "cancel"
	MBTTSubmit                      ButtonBlockObjectType    = "submit"
)

type ButtonBlockObject struct {
	Type     ButtonBlockObjectType    `json:"type"`
	Label    string                   `json:"label"`
	Color    ButtonBlockObjectColor   `json:"color,omitempty"`
	Variant  ButtonBlockObjectVariant `json:"variant,omitempty"`
	Action   *Action                  `json:"action,omitempty"`
	Icon     string                   `json:"icon"`
	Disbaled bool                     `json:"disabled"`
}
