package models

import (
	"errors"
	"fmt"
)

// TableBlock is a subtype of block. It represents a table block.
type tableBlock struct {
	block

	// object of the table block
	Table *TableBlockObject `json:"table"`
}

// TableBlockObject holds the detail of the TableBlock.
type TableBlockObject struct {
	appendable

	// Header contains an array of HeaderObject. All table headers should be
	// defined here.
	Header []HeaderObject `json:"header,omitempty"`

	// The field `Body` holds a 3D array of IBlock. The outermost array represents
	// the table rows. The middle array represents the table columns. The innermost
	// array represents data in each column.
	Body [][][]IBlock `json:"body"`
}

// HeaderObject is the struct for field `Header` in TableBlockObject. Header
// can only be in text form.
type HeaderObject struct {
	Type string            `json:"type"`
	Text *TextHeaderObject `json:"text"`
}

// TextHeaderObject is the struct for field `Text` in HeaderObject
type TextHeaderObject struct {
	// Align is positioning column.
	// Align is not required, default value is left
	Align string `json:"align,omitempty"`

	// Body is a content of column
	Body string `json:"body"`
}

// Validate performs validation to the TableBlock.
func (ths *tableBlock) Validate() (bool, []error) {
	var errs []error
	if ths.Type != MBTTable {
		errs = append(errs, errors.New("invalid table block type"))
	}

	if len(ths.Table.Header) > 0 && len(ths.Table.Body) > 0 {
		for _, row := range ths.Table.Body {
			if len(row) != len(ths.Table.Header) {
				errs = append(errs, fmt.Errorf("the number of headers and columns must be the same, header = %v, column = %v", len(ths.Table.Header), len(row)))
			}
		}
	}

	for _, row := range ths.Table.Body {
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
func NewTableBlock() *tableBlock {
	block := tableBlock{
		Table: &TableBlockObject{},
	}
	block.Type = MBTTable

	return &block
}

// AddHeader have 2 params (body and align) but body is required, align have default value is left
func (ths *tableBlock) AddHeader(body string, align string) {
	var headerObj HeaderObject
	var alignText string

	alignText = "left"
	if align != "" {
		alignText = align
	}
	headerObj.Type = "text"
	headerObj.Text = &TextHeaderObject{
		Align: alignText,
		Body:  body,
	}
	ths.Table.Header = append(ths.Table.Header, headerObj)
}

// AddColumn use to add column on body table
func (ths *tableBlock) AddColumn(column ...IBlock) []IBlock {
	return column
}

// AddRow use to add row on body table
func (ths *tableBlock) AddRow(row ...[]IBlock) [][]IBlock {
	return row
}

// AddBody use to add body on table
func (ths *tableBlock) AddBody(rows ...[][]IBlock) {
	ths.Table.Body = rows
}
