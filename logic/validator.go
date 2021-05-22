package logic

import "errors"

type Validator interface {
	ValidateConverterArgs(...string) error
}

type validate struct {
}

func NewValidator() Validator {
	return &validate{}
}

func (v validate) ValidateConverterArgs(args ...string) error {
	for _, input := range args {
		switch input {
		// acceptable file type
		case "csv", "json":
		// acceptable contract type
		case "b3":
		default:
			return errors.New("invalid args")
		}
	}
	return nil
}
