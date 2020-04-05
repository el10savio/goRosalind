package definitions

import "errors"

var (
	// ErrEmptyString is the error in case of empty string
	ErrEmptyString = errors.New("string cannot be empty")

	// ErrInvalidDNAString is the error in case of invalid DNA string
	ErrInvalidDNAString = errors.New("invalid dna string")

	// ErrNotPositiveValue is the error in case of not positive input value
	ErrNotPositiveValue = errors.New("not positive input value")
)
