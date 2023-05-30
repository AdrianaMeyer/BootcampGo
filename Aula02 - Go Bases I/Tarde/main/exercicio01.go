package main

import "fmt"

func main() {
	var palavra string
	palavra = "exemplo"
	
	var tamanho int = len(palavra)
	fmt.Println("Quantidade de letras na palavra: ", tamanho)

    for i := 0; i < len(palavra); i++ {
        fmt.Printf("%c\n",palavra[i])
    }
}
