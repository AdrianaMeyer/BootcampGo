package main

import "fmt"

func main() {
	
	var employees = map[string]int{"Benjamin": 20, "Manuel": 26, "Brenda": 19, "Dario": 44, "Pedro": 30}
	var contador int = 0

	fmt.Println()
	fmt.Println("Idade do Benjamin: ", employees["Benjamin"], "anos")
	fmt.Println()

	for key, element := range employees {
		if element > 21 {
			fmt.Println("O funcionario ", key, "tem mais de 21 anos")
			contador++
		}
	}

	fmt.Println("Total de funcionarios com mais de 21 anos: ", contador)

	fmt.Println()
	employees["Federico"] = 25
	fmt.Println("Incluindo o funcionario Federico: ", employees)

	fmt.Println()
	delete (employees, "Pedro") 
	fmt.Println("Deletando o Pedro: ", employees)
}