package models

import "errors"

// CardBlock is a subtype of block. It represents a card block and
// holds a CardBlockObject in the field `Card`.
type cardBlock struct {
	block
	// Card contains the CardBlockObject that holds the detail of card block
	Card *CardBlockObject `json:"card"`
}

// CardBlockObject holds the detail of the CardBlock. CardBlockObject
// can contain other blocks.
type CardBlockObject struct {
	appendable

	// Headers is the header of the card, contains blocks.
	// This field is optional.
	Headers []IBlock `json:"headers,omitempty"`
}

// Validate performs validation to the CardBlock.
func (ths *cardBlock) Validate() (bool, []error) {
	// CardBlock validation implementation
	var errs []error
	if ths.Type != MBTCard {
		errs = append(errs, errors.New("invalid card block type"))
	}

	for _, v := range ths.Card.Headers {
		if valid, err := v.Validate(); !valid {
			errs = append(errs, err...)
		}
	}

	for _, v := range ths.Card.Blocks {
		if valid, err := v.Validate(); !valid {
			errs = append(errs, err...)
		}
	}

	if len(errs) > 0 {
		return false, errs
	}

	return true, nil
}

// NewCardBlock returns a new instance of a card block to be rendered
func NewCardBlock(cardObj CardBlockObject) *cardBlock {
	obj := CardBlockObject{
		Headers: cardObj.Headers,
	}

	var block cardBlock
	block.Type = MBTCard
	block.Card = &obj

	return &block
}

// AddCardHeader adds blocks (`IBlock`) to the field `Blocks` in
// the embedded struct Card.Header `cardBlock`
func (ths *cardBlock) AddCardHeader(blocks ...IBlock) {
	ths.Card.Headers = append(ths.Card.Headers, blocks...)
}
