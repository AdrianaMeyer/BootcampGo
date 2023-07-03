package main_test

import (
	"testing"

	calc "github.com/AdrianaMeyer/BootcampGo"
)

func TestFibonacci(t *testing.T){
	tests := []struct {
		arg  int
		expectedResult int
	}{{2, 1}, {6, 8}, {10, 55}, {30, 832040}}

	for i, test := range tests {
		result := calc.Fibonacci(test.arg)
		if result != test.expectedResult {
			t.Errorf("Teste[%d]: Fibonacci(%d) retornou %d, mas o resultado esperado era %d",
				i, test.arg, result, test.expectedResult)
		}
	}
}