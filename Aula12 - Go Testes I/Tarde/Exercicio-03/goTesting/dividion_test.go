package goTesting_test

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"errors"

	"github.com/AdrianaMeyer/BootcampGo/goTesting"
)

func TestDivision(t *testing.T) {
	result, _ := goTesting.Division(10, 2)
	expectedResult := 5

	assert.Equal(t, result, expectedResult, "o resultado esperado e o resultado devem ser iguais")
}

func TestDivisionWithZero(t *testing.T) {
	_, err := goTesting.Division(10, 0)
	expectedResult := errors.New("O denominador não pode ser 0")

	assert.Equal(t, expectedResult, err, "o resultado esperado e o resultado devem ser iguais")
}