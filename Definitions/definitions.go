package definitions

import "errors"

var (
	// ErrEmptyString is the error in case of empty string
	ErrEmptyString = errors.New("string cannot be empty")

	// ErrEmptyList is the error in case of empty list
	ErrEmptyList = errors.New("list cannot be empty")

	// ErrEmptyLabel is the error in case of empty label
	ErrEmptyLabel = errors.New("label cannot be empty")

	// ErrInvalidDNAString is the error in case of invalid DNA string
	ErrInvalidDNAString = errors.New("invalid dna string")

	// ErrNotPositiveValue is the error in case of not positive input value
	ErrNotPositiveValue = errors.New("not positive input value")

	// ErrUnequalLength is the error in case of the input string lengths are unequal
	ErrUnequalLength = errors.New("input string lengths are unequal")
)
