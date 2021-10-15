package service

import (
	"strings"
)

type Service struct {
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) CheckString(str string) (string, error) {
	str = strings.Trim(str, "\n")

	stringByRunes := []rune(str)
	switch {
	case len(stringByRunes) == 0:
		return "string is empty", nil

	case stringByRunes[len(stringByRunes)-1] != '.':
		return "last element is not '.' ", nil

	case stringByRunes[0] < 'A' || stringByRunes[0] > 'Z':
		return "first element is not big letter", nil
	}

	return "String is OK", nil
}
