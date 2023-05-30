package main

import (
	"fmt"
	"github.com/AdrianaMeyer/BootcampGo/exercicio01"
	"github.com/AdrianaMeyer/BootcampGo/exercicio02"
	"github.com/AdrianaMeyer/BootcampGo/exercicio03"
	"github.com/AdrianaMeyer/BootcampGo/exercicio04"
)

func main() {
	fmt.Println()
	fmt.Println("Exercício 1 - Imprimindo o nome na tela")
	fmt.Println("---------------------------")
	fmt.Println("Meu nome é: ", exercicio01.Nome()) 
	fmt.Println("Minha idade é: ", exercicio01.Idade())

	fmt.Println()
	fmt.Println("Exercício 2 - Clima")
	fmt.Println("---------------------------")
	fmt.Println("Temperatura:", exercicio02.Temperatura(), "C")
	fmt.Println("Umidade:", exercicio02.Umidade())
	fmt.Println("Pressao Atmosferica:", exercicio02.PressaoAtmosferica(), "ATM")

	fmt.Println()
	fmt.Println("Exercício 3 - Declaração de variáveis")
	fmt.Println("---------------------------")
	fmt.Println("Nome:", exercicio03.Nome())
	fmt.Println("Sobrenome:", exercicio03.Sobrenome())
	fmt.Println("Idade:", exercicio03.Idade())
	fmt.Println("Pode dirigir:", exercicio03.Dirige())
	fmt.Println("Estatura:", exercicio03.Estatura())
	fmt.Println("Filhos:", exercicio03.Filhos())


	fmt.Println()
	fmt.Println("Exercício 4 - Tipos de dados")
	fmt.Println("---------------------------")
	fmt.Println("Nome:", exercicio04.Nome())
	fmt.Println("Sobrenome:", exercicio04.Sobrenome())
	fmt.Println("Idade:", exercicio04.Idade())
	fmt.Println("Verdadeiro ou falso:", exercicio04.Boolean())
	fmt.Println("Salário: R$", exercicio04.Salario())

}
