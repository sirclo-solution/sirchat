package models

import "errors"

// collapsibleBlock is a subtype of block. It represents a collapsible block and
// holds a CollapsibleBlockObject in the field `collapsible`.
type collapsibleBlock struct {
	block
	// collapsible contains the CollapsibleBlockObject that holds the detail of collapsible block
	Collapsible CollapsibleBlockObject `json:"collapsible"`
}

// CollapsibleBlockObject holds the detail of the collapsibleBlock.
type CollapsibleBlockObject struct {
	// Title is the text that will be shown as the title of collapsible.
	// Title is required.
	Title string `json:"title"`

	// Collapsed that determines whether its contents are collapsed true `closed` or false `opened``
	Collapsed bool `json:"collapsed"`

	// Content is the content of the collapsible, contains blocks.
	// This field is required.
	Content []IBlock `json:"content"`
}

// Validate performs validation to the CardBlock.
func (ths *collapsibleBlock) Validate() (bool, []error) {
	// CardBlock validation implementation
	var errs []error
	if ths.Type != MBTCollapsible {
		errs = append(errs, errors.New("invalid collapsible block type"))
	}

	if ths.Collapsible.Title == "" {
		errs = append(errs, errors.New("field Title is required"))
	}

	if len(ths.Collapsible.Content) == 0 {
		errs = append(errs, errors.New("field Content must contain one block"))
	} else {
		for _, v := range ths.Collapsible.Content {
			if valid, err := v.Validate(); !valid {
				errs = append(errs, err...)
			}
		}
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewCollapsibleBlock returns a new instance of a collapsible block to be rendered
func NewCollapsibleBlock(collapsibleObj CollapsibleBlockObject) *collapsibleBlock {
	obj := CollapsibleBlockObject{
		Title:     collapsibleObj.Title,
		Collapsed: collapsibleObj.Collapsed,
		Content:   collapsibleObj.Content,
	}

	var block collapsibleBlock
	block.Type = MBTCollapsible
	block.Collapsible = obj

	return &block
}

// AddContentCollapsible adds blocks (`IBlock`) to the field `Blocks` in
// the embedded struct Collapsible.Content `collapsibleBlock`
func (ths *collapsibleBlock) AddContentCollapsible(blocks ...IBlock) {
	ths.Collapsible.Content = append(ths.Collapsible.Content, blocks...)
}
