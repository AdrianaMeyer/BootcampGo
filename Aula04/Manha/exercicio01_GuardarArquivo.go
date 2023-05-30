package main

import (
	"fmt"
	"os"
)

type Produto struct {
	id int
	preco float64
	quantidade int
}

func main() {

	produto1 := Produto {
		id: 12464,
		preco: 12.50,
		quantidade: 45,
	}

	/*produto2 := Produto {
		id: 23547,
		preco: 50.49,
		quantidade: 130,
	}
	*/

	infoProduto := fmt.Sprintf("ID;Preco;Quantidade;%d;%.2f;%d", produto1.id, produto1.preco, produto1.quantidade)
	
	conteudo := []byte(infoProduto)
	erro := os.WriteFile("./relacaoProdutos.csv", conteudo, 0644)

	if erro != nil {
		fmt.Println("Erro ao escrever as informacoes no arquivo", erro)
	}

}