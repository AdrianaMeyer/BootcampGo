package main

import "fmt"

func main() {
	
	var idade int = 22
	var empregado bool = true
	var tempoDeAtividade int = 2
	var salario float32 = 2000
	var elegivel bool = false
	
	if idade >= 22 && empregado == true && tempoDeAtividade > 1 {
		elegivel = true
	} else {
		fmt.Println("Empréstimo negado")
	}
	
	if elegivel == true && salario < 100000.00 {
		fmt.Println("Empréstimo concedido com juros")
	} else if elegivel == true && salario >= 100000.00 {
		fmt.Println("Empréstimo concedido sem juros")
	} 

}