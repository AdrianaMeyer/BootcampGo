package main

import (
	"fmt"
	"errors"
)

func calculaMedia(notas[] float64) (float64, error) {
	var media float64 = 0.0
	soma := 0.0
	contador := 0.0

	for _, nota := range notas {

		if nota < 0.0 {
			return 0, errors.New("Não sao aceitas notas com valores negativos")
		} else {
			soma += nota
			contador++
		}
	}

	media = soma / contador
	return media, nil
}

func main() {

	notasAluno1 := []float64{1.5, 4.9, 5, 8.5, 10}
	mediaAluno, err := calculaMedia(notasAluno1)

	//notasAluno2 := []float64{2, 7, -5, 9.2, 10.1}
	//mediaAluno, err := calculaMedia(notasAluno2)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Printf("Media do aluno: %.2f\n", mediaAluno)
	}
}