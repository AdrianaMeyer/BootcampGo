package main

import (
	"fmt"
	"os"
)

type meuErroCustomizado struct {
	status string
	msg    string
}

func (m *meuErroCustomizado) Error() string {
	return fmt.Sprintf("%s: %v", m.status, m.msg)
}

func validaErro(salario int) error {

	if salario <= 15000 {
		return &meuErroCustomizado{
			status: "Erro",
			msg: "O salário digitado não está dentro do valor mínimo",
		}
	}
	return nil
}

func main() {
	var salario int = 2000
	
	err := validaErro(salario)

	if err != nil { 
		fmt.Println(err) 
		os.Exit(1) 
	}
	fmt.Printf("Necessário pagamento de imposto!\n")
}
