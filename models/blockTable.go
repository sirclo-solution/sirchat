package models

import "errors"

// TableBlock is a subtype of block. It represents a table block.
type TableBlock struct {
	block
	Type  MessageBlockType  `json:"type"`
	Table *TableBlockObject `json:"table,omitempty"`
}

// TableBlockObject holds the detail of the TableBlock. The field `Body`
// holds a 3D array of IBlock. The outermost array represents the table rows.
// The middle array represents the table columns. The innermost array represents
// data in each column.
type TableBlockObject struct {
	appendable
	Header []HeaderObject `json:"header"`
	Body   [][][]IBlock   `json:"body"`
}

// HeaderObject is the struct for field `Header` in TableBlockObject. Header
// can only be in text form.
type HeaderObject struct {
	Type string            `json:"type"`
	Text *TextHeaderObject `json:"text,omitempty"`
}

// TextHeaderObject is the struct for field `Text` in HeaderObject
type TextHeaderObject struct {
	Align string `json:"align,omitempty"`
	Body  string `json:"body"`
}

// Validate performs validation to the TableBlock.
func (s TableBlock) Validate() (bool, []error) {
	var errs []error
	if s.Type != MBTTable {
		errs = append(errs, errors.New("invalid table block type"))
	}

	for _, row := range s.Table.Body {
		for _, column := range row {
			for _, v := range column {
				if valid, err := v.Validate(); !valid {
					errs = append(errs, err...)
				}
			}
		}
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewTableBlock returns a new instance of a table block to be rendered
func NewTableBlock(tableHeader []HeaderObject, body [][][]IBlock) *TableBlock {
	block := TableBlock{
		Type: MBTTable,
	}

	return &block
}
