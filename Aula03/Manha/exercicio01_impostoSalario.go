package main

import(
	"fmt"
)

func impostoSalario(salario float32) float32 {
	var imposto float32

	if salario < 150000 {
		imposto = 0.17
	} else if salario >= 150000 {
		imposto = 0.17 + 0.10
	} else {
		imposto = 0
	}

	return imposto
}

func calculaSalarioComImposto(salario float32, imposto float32) float32{
	salario = salario - (salario * imposto)
	return salario
}


func main() {

	var salarioFunc1 float32 = 50000
	var salarioFunc2 float32 = 150000
	
	fmt.Println("Imposto a ser descontado:", (impostoSalario(salarioFunc1)*100),"%")
	fmt.Println("Salário do funcionario 1 com desconto: R$", calculaSalarioComImposto(salarioFunc1, impostoSalario(salarioFunc1)))
	
	fmt.Println("Imposto a ser descontado:", (impostoSalario(salarioFunc2)*100),"%")
	fmt.Println("Salário do funcionario 2 com desconto: R$", calculaSalarioComImposto(salarioFunc2, impostoSalario(salarioFunc2)))
}

