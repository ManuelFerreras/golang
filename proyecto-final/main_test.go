package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	result, err := add(1, 2)   // Call function from the main package.
	assert.Nil(t, err)         // Check that there was no error.
	assert.Equal(t, 3, result) // Check that the result is correct.
}

func TestDivision(t *testing.T) {
	result, err := division(4, 2) // Call function from the main package.
	assert.Nil(t, err)            // Check that there was no error.
	assert.Equal(t, 2, result)    // Check that the result is correct.
}

func TestDivisionByZero(t *testing.T) {
	result, err := division(4, 0) // Call function from the main package.
	assert.NotNil(t, err)         // Check that there was an error.
	assert.Equal(t, 0, result)    // Check that the result is correct.
}
