package models

import (
	"github.com/almarino_meli/grupo-5-wave-15/pkg/errors"
	"github.com/go-playground/validator/v10"
)

type ID struct {
	ID int `validate:"omitempty,numeric,min=0"`
}

func NewID(id int) (ID, error) {
	if id < 0 {
		return ID{-1}, &errors.ValidationError{Message: "ID must be greater than 0"}
	}
	return ID{id}, nil
}

func (id *ID) Validate(validate *validator.Validate) error {
	return validate.Struct(id)
}
