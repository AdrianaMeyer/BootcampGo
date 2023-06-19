package goTesting_test

import (
	"testing"
	"github.com/stretchr/testify/assert"

	"github.com/AdrianaMeyer/BootcampGo/goTesting"
)

func TestOrdering(t *testing.T) {
	arrayTest := []int{22, 9, 5, 1, 3, 8}

	result := goTesting.Ordering(arrayTest)
	expectedResult := []int{1, 3, 5, 8, 9, 22}

	assert.Equal(t, expectedResult, result, "os arrays devem ser iguais")
}