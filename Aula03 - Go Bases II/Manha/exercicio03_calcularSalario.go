package main

import (
	"fmt"
)

func calculaSalario(categoria string, qntHoras float64) float64{
	salario := 0.0
	
	if categoria == "C" {
		salario = 1000 * qntHoras
	} else if categoria == "B" {
		if qntHoras <= 160 { 
			salario = (1500 * qntHoras)
		} else {
			salario = (1500 * qntHoras)
			salario = salario + (salario * 0.2)
		}
	} else if categoria == "A" {
		if qntHoras <= 160 { 
			salario = (3000 * qntHoras)
		} else {
			salario = (3000 * qntHoras)
			salario = salario + (salario * 0.5)
		}
	} else {
		return 0
	}

	return salario
}

func main() {

	fmt.Printf("Salario do funcionário da categoria C: %.2f\n", calculaSalario("C", 162))
	fmt.Printf("Salario do funcionário da categoria B: %.2f\n", calculaSalario("B", 176))
	fmt.Printf("Salario do funcionário da categoria A: %.2f\n", calculaSalario("A", 172))
	fmt.Printf("Salario do funcionário da categoria INEXISTENTE: %.2f\n", calculaSalario("F", 180))
}