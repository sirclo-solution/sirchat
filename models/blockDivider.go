package models

import (
	"errors"
	"fmt"
	"regexp"
)

type DividerBlockObject struct {
	// Color is the color of the divider
	// This field is optional
	Color string `json:"color"`

	// padding top is the padding top on divider
	// This field can only accept units : px, rem, em, ex, pt, in, pc, mm, cm
	// Example value: 16px
	// This field is optional
	PaddingTop string `json:"padding_top"`

	// padding bottom is the padding bottom on divider
	// This field can only accept units : px, em, ex, pt, in, pc, mm, cm
	// Example value: 16px
	// This field is optional
	PaddingBottom string `json:"padding_bottom"`
}

// DividerBlock is a subtype of block. It represents a divider block. It
// will render as a divider line between blocks.
type dividerBlock struct {
	block

	// Divider contains the DividerBlockObject that holds the detail of divider block
	Divider DividerBlockObject `json:"divider"`
}

// Validate performs validation to the DividerBlock.
func (ths *dividerBlock) Validate() (bool, []error) {
	// DividerBlock validation implementation
	var errs []error
	if ths.Type != MBTDivider {
		errs = append(errs, errors.New("invalid container block type"))
	}

	if ths.Divider.PaddingTop != "" {
		matchPaddingTop, errPaddingTop := regexp.MatchString(`^([0-9]+)(px|rem|em|ex|pt|in|pc|mm|cm)$`, ths.Divider.PaddingTop)
		if errPaddingTop != nil {
			errs = append(errs, fmt.Errorf("error when validating divider padding top: %s", errPaddingTop.Error()))
		} else if !matchPaddingTop {
			errs = append(errs, fmt.Errorf("invalid divider padding top: %s", ths.Divider.PaddingTop))
		}
	}

	if ths.Divider.PaddingBottom != "" {
		matchPaddingBottom, errPaddingBottom := regexp.MatchString(`^([0-9]+)(px|rem|em|ex|pt|in|pc|mm|cm)$`, ths.Divider.PaddingBottom)
		if errPaddingBottom != nil {
			errs = append(errs, fmt.Errorf("error when validating divider padding bottom: %s", errPaddingBottom.Error()))
		} else if !matchPaddingBottom {
			errs = append(errs, fmt.Errorf("invalid divider padding bottom: %s", ths.Divider.PaddingBottom))
		}
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewDividerBlock returns a new instance of a divider block to be rendered
func NewDividerBlock(dividerObj DividerBlockObject) *dividerBlock {
	block := dividerBlock{}
	block.Type = MBTDivider
	block.Divider = DividerBlockObject{
		Color:         dividerObj.Color,
		PaddingTop:    dividerObj.PaddingTop,
		PaddingBottom: dividerObj.PaddingBottom,
	}

	return &block
}
