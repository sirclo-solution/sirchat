package models

import "errors"

// TableBlock defines a new block of type section
type TableBlock struct {
	block
	Type  MessageBlockType  `json:"type"`
	Table *TableBlockObject `json:"table,omitempty"`
}

type TableBlockObject struct {
	appendable
	Header []HeaderObject `json:"header"`
	Body   [][][]IBlock   `json:"body"`
}

type HeaderObject struct {
	Type string            `json:"type"`
	Text *TextHeaderObject `json:"text,omitempty"`
}

type TextHeaderObject struct {
	Align string `json:"align,omitempty"`
	Body  string `json:"body"`
}

func (s TableBlock) Validate() (bool, []error) {
	// TableBlock validation implementation
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

// NewTableBlock returns a new instance of a section block to be rendered
func NewTableBlock(tableHeader []HeaderObject, body [][][]IBlock) *TableBlock {
	block := TableBlock{
		Type: MBTTable,
		Table: &TableBlockObject{
			Header: tableHeader,
			Body:   body,
		},
	}

	return &block
}
