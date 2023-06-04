package main

import (
	"fmt"
	"os"
	"strings"
)

type Produto struct {
	Id int
	Preco float64
	Quantidade int
}

func NovoProduto(id int, preco float64, quant int) Produto {
	novoProduto := Produto {
		Id: id,
		Preco: preco,
		Quantidade: quant,
	}
	return novoProduto
}

func main() {

	produto1 := NovoProduto(101, 12.50, 15)
	produto2 := NovoProduto(102, 150.90, 23)
	produto3 := NovoProduto(103, 204.99, 19)

	produtos := []Produto {produto1, produto2, produto3}

	var linhas []string

	linhas = append(linhas, "ID;Preco;Quantidade;")

	for _, produto := range produtos {
		linha := fmt.Sprintf("%d;%.2f;%d;", produto.Id, produto.Preco, produto.Quantidade)
		linhas = append(linhas, linha)
	}

	conteudoCsv := []byte(strings.Join(linhas, "\n"))

	err := os.WriteFile("./relacaoProdutos.csv", conteudoCsv, 0644)

	if err != nil {
		fmt.Println("Erro ao escrever as informacoes no arquivo", err)
	}

	fmt.Println("Dados salvos com sucesso")

}