package validator

import (
	"context"
	"errors"
)

type Validator struct {
}

func New() *Validator {
	return &Validator{}
}

func (v *Validator) Validate(_ context.Context, str string) error {
	stringByRunes := []rune(str)
	switch {
	case len(stringByRunes) == 0:
		return errors.New("string is empty")

	case stringByRunes[len(stringByRunes)-1] != '.':
		return errors.New("last element is not '.' ")

	case stringByRunes[0] < 'A' || stringByRunes[0] > 'Z':
		return errors.New("first element is not big letter")
	}

	return nil
}
