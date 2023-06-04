package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
	"log"
)

func main() {

	var dadosProduto []string

	arquivo, err := os.Open("./relacaoProdutos.csv")

	if err != nil {
		fmt.Println("Houve erro ao abrir o arquivo", err)
		arquivo.Close()
	}

	leitor := bufio.NewReader(arquivo)

	for {	
		dado, err := leitor.ReadString(';')
		dado = strings.Trim(dado, ";")

		if err == io.EOF {
			break
		}

		dadosProduto = append(dadosProduto, dado)
		
	}
	arquivo.Close()

	var total float64 = 0.0

	//CABECALHO
	for i := 0; i < 1; i++ {
		fmt.Printf("%s\t %s\t %s\t", dadosProduto[i], dadosProduto[i+1], dadosProduto[i+2])

	}

	//CORPO
	for j := 3; j < len(dadosProduto); {

		preco, err := strconv.ParseFloat(dadosProduto[j+1], 64)
		quant, err := strconv.ParseFloat(dadosProduto[j+2], 64)

		if err != nil {
		log.Fatal("Erro ao converter a string para float64:", err)
		}

		total = total + (preco * quant)

		fmt.Printf("%s\t %s\t %s\t", dadosProduto[j], dadosProduto[j+1], dadosProduto[j+2])
		j += 3

	}

	fmt.Printf("\n%s\t %.2f\t", "TOTAL: ", total)
	
}

