package models

type Subheading struct {
	Text string `json:"text"`
}

func NewSubheading(text string) Subheading {
	return Subheading{
		Text: text,
	}
}
