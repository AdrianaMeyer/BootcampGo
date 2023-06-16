package calc_test

import (

	"github.com/stretchr/testify/assert"
)


func TestSubtract(t *testing.T) {
	result := calc.Subtract(5, 3)
	expectedResult := 0

	assert.Equal(t, expectedResult, result, "the result must be equal to expected result")
}