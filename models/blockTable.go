package models

// TableBlock defines a new block of type section
type TableBlock struct {
	Block
	Type  MessageBlockType  `json:"type"`
	Table *TableBlockObject `json:"table,omitempty"`
}

type TableBlockObject struct {
	BlockObject
	Header []HeaderObject `json:"header"`
	Body   [][]Body       `json:"body"`
}

type HeaderObject struct {
	Type string            `json:"type"`
	Text *TextHeaderObject `json:"text,omitempty"`
}

type TextHeaderObject struct {
	Align string `json:"align,omitempty"`
	Body  string `json:"body"`
}

func (s TableBlock) Validate() bool {
	// TableBlock validation implementation

	return true
}

// NewTableBlock returns a new instance of a section block to be rendered
func NewTableBlock(tableHeader []HeaderObject, body [][]Body) *TableBlock {
	block := TableBlock{
		Type: MBTTable,
		Table: &TableBlockObject{
			Header: tableHeader,
			Body:   body,
		},
	}

	return &block
}
