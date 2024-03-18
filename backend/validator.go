package main

import (
	"github.com/go-playground/validator"
)

func Validate(input JokeInput) error {
	// Create a new validator using the go-playground/validator package
	validate := validator.New()

	// Validate the input considering the struct tags declared in the JokeInput struct
	err := validate.Struct(input)
	if err != nil {
		return err
	}
	return nil
}