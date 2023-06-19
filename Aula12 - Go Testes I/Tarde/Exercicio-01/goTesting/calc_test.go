package goTesting_test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/AdrianaMeyer/BootcampGo/goTesting"
)


func TestSubtract(t *testing.T) {
	result := goTesting.Subtract(5, 3)
	expectedResult := 2

	assert.Equal(t, expectedResult, result, "the result must be equal to expected result")
}