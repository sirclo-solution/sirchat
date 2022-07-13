package models

type Title struct {
	Text string
	Icon string
}

func NewTitle(text, icon string) Title {
	return Title{
		Text: text,
		Icon: icon,
	}
}
