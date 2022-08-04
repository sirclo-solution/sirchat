package models

type Subheading struct {
	Text    string `json:"text"`
	Divider bool   `json:"divider"`
}

func NewSubheading(text string) Subheading {
	return Subheading{
		Text: text,
	}
}
