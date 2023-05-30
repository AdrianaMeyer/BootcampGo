package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	arquivo, erro := os.ReadFile("./relacaoProdutos.csv")

	if erro != nil {
		fmt.Println("Houve erro ao ler o arquivo", erro)
	}

	dadosProduto := strings.Slipt(arquivo, ";")

	for _, dado := range dadosProduto {
		fmt.Println(dado)
	}

	//EXERCICIO EM AJUSTES
}
