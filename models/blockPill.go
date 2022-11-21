package models

import (
	"errors"
	"fmt"
)

// PillBlockObjectType is a type to define type pill
type PillBlockObjectType string

const (
	// primary type will use primary (#1c9aea)
	PillBlockObjectTypePrimary PillBlockObjectType = "primary"

	// secondary_1 type will use secondary (#1c9aea) with grey text (#1c9aea)
	PillBlockObjectTypeSecondary1 PillBlockObjectType = "secondary_1"

	// secondary_2 type will use secondary (#1c9aea) with white text (#1c9aea)
	PillBlockObjectTypeSecondary2 PillBlockObjectType = "secondary_2"

	// success_1 type will use success (#00b969); type will ligther than success_2
	PillBlockObjectTypeSuccess1 PillBlockObjectType = "success_1"

	// success_2 type will use success (#008566); type will darker than success_1
	PillBlockObjectTypeSuccess2 PillBlockObjectType = "success_2"

	// info type will use info (#1478c9); type will lighter than info_2
	PillBlockObjectTypeInfo1 PillBlockObjectType = "info_1"

	// info type will use info (#083f87); type will darker than info_2
	PillBlockObjectTypeInfo2 PillBlockObjectType = "info_2"

	// warning type will use warning (#ffb300)
	PillBlockObjectTypeWarning PillBlockObjectType = "warning"

	// danger type will use danger (#d64241)
	PillBlockObjectTypeDanger PillBlockObjectType = "danger"
)

// PillBlock is a subtype of block. it represents a pill block
type pillBlock struct {
	block
	Pill *PillBlockObject `json:"pill"`
}

// PillBlockObject holds the detail of the PillBlock
type PillBlockObject struct {
	appendable

	// text is content of pill block
	// this field is required
	Text string `json:"text"`

	// type is the background color of the pill
	// primary will be type default value
	Type PillBlockObjectType `json:"type"`
}

// Validate performs validation to the PillBlock. Field `Text`
// should not be empty.
func (ths *pillBlock) Validate() (bool, []error) {
	// PillBlock validation implementation
	var errs []error

	if ths.Pill.Text == "" {
		errs = append(errs, errors.New("text is missing"))
	}

	if typeValid := ths.Pill.Type.validatePillObjectType(); !typeValid {
		errs = append(errs, fmt.Errorf("invalid BadgeBlockObjectType %v", ths.Pill.Type))
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewPillBlock returns a new instance of a section block to be rendered
func NewPillBlock(pillObj *PillBlockObject) *pillBlock {
	var typePill PillBlockObjectType

	typePill = PillBlockObjectTypePrimary

	if string(pillObj.Type) != "" {
		typePill = pillObj.Type
	}

	block := pillBlock{
		Pill: &PillBlockObject{
			Text: pillObj.Text,
			Type: typePill,
		},
	}
	block.Type = MBTPill

	return &block
}

func (t PillBlockObjectType) validatePillObjectType() bool {
	switch t {
	case "":
		return true
	case PillBlockObjectTypePrimary:
		return true
	case PillBlockObjectTypeSecondary1:
		return true
	case PillBlockObjectTypeSecondary2:
		return true
	case PillBlockObjectTypeSuccess1:
		return true
	case PillBlockObjectTypeSuccess2:
		return true
	case PillBlockObjectTypeInfo1:
		return true
	case PillBlockObjectTypeInfo2:
		return true
	case PillBlockObjectTypeWarning:
		return true
	case PillBlockObjectTypeDanger:
		return true
	default:
		return false
	}
}
