package main

import (
	"fmt"
	"os"
	"errors"
)

func validaErro(salario int) error {

	if salario <= 15000 {
		return errors.New("O salário digitado não está dentro do valor mínimo")
	}
	return nil
}

func main() {
	var salario int = 1000
	
	err := validaErro(salario)

	if err != nil { 
		fmt.Println(err) 
		os.Exit(1) 
	}
	fmt.Printf("Necessário pagamento de imposto!\n")
}

