package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {

	var dadosClientes []string

	arquivo, err := os.Open("./customers.txt")

	if err != nil {
		panic("O arquivo indicado não foi encontrado ou está danificado")	
	}
	
	leitor := bufio.NewReader(arquivo)

	for {	
		dado, err := leitor.ReadString('\n')
		dado = strings.TrimSpace(dado)

		dadosClientes = append(dadosClientes, dado)

		if err == io.EOF {
			break
		}
	}

	arquivo.Close()
	
	fmt.Println("Dados dos clientes: ", dadosClientes)
	
}
