package main

import (
	"fmt"
	"errors"
)

func CalculaSalario(horasMensais int, valorHora float64) (float64, error) {
	salario := 0.0

	if horasMensais <= 0 {
		return 0, fmt.Errorf("Erro: a quantidade de horas trabalhadas nao pode ser um valor negativo.")
	} else if horasMensais > 0 && horasMensais < 80  {
		return 0, errors.New("Erro: o trabalhador não pode ter trabalhado menos de 80 horas por mes")
	} else {	
		salario = float64(horasMensais) * valorHora

		if salario >= 20000 {
			salario = salario - (salario * 0.1)
		}

		return salario, nil
	}

}

func main() {
	
	salario, erro2 := CalculaSalario(90, 8.9)

	erro1 := fmt.Errorf("%w", erro2)

	if erro1 != nil {
		fmt.Println(errors.Unwrap(erro1))
	} 
	
	fmt.Printf("O salario total é: R$ %.2f\n", salario)

}