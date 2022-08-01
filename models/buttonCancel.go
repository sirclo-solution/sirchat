package models

type CancelButton struct {
	Button
}

func (s CancelButton) Validate() (bool, error) {
	// CancelButton validation implementation

	return true, nil
}

// NewCancelButton returns a new instance of a section block to be rendered
func NewCancelButton(label string) *CancelButton {
	var button CancelButton
	button.Label = label
	button.Type = MBTTCancel

	return &button
}
