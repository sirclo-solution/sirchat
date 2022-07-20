package models

type Title struct {
	Text string `json:"text"`
	Icon string `json:"icon"`
}

func NewTitle(text, icon string) Title {
	return Title{
		Text: text,
		Icon: icon,
	}
}
