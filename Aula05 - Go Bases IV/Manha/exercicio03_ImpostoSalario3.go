package main

import (
	"fmt"
	"os"
)

func validaErro(salario int) error {

	if salario <= 15000 {
		return fmt.Errorf("Erro: o mínimo tributável é 15.000 e o salário informado é: %d", salario)
	} 
	return nil
}

func main() {
	var salario int = 18000

	err := validaErro(salario)

	if err != nil { 
		fmt.Println(err) 
		os.Exit(1) 
	}
	fmt.Printf("Salario informado: %d. Necessário pagamento de imposto!\n", salario)
}
