package goTesting

import "errors"

func Division(num1, num2 int) (int, error) {
	if num2 == 0 {
		return 0, errors.New("O denominador não pode ser 0")
	}

	result := num1 / num2
	return result, nil
}