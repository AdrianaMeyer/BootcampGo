package main

import (
	"fmt"
	"errors" 
)

const (
	Max = "maximo"
	Med = "media"
	Min = "minimo"
)


func calculaMax(valores ...float64) float64 {
	maiorValor := 0.0
	for _, valor := range valores {
		if valor > maiorValor {
			maiorValor = valor
		}
	}
	return maiorValor
}

func calculaMedia(valores ...float64) float64 {
	media := 0.0 
	soma:= 0.0
	qnt := 0.0

	for _, valor := range valores {
		soma += valor
		qnt++
	}
	media = soma / qnt
	return media
}

func calculaMin(valores ...float64) float64 {
	menorValor := valores[0]

	for i := 1; i < len(valores); i++ { 
		if valores[i] < menorValor {
			menorValor = valores[i]
		}
	}

	return menorValor
}

func alertaErro(valores ...float64) float64 {
	return 0
}

func operacao(operacao string) (func(valores ...float64) float64, error) {
	switch operacao {
		case "maximo":
			return calculaMax, nil
		case "media":
			return calculaMedia, nil
		case "minimo": 
			return calculaMin, nil
	}
	return alertaErro, errors.New("A operacao é inválida e não pode ser realizada")
}



func main() {

	maxFunc, err := operacao(Max)
	medFunc, err := operacao(Med)
	minFunc, err := operacao(Min)

	maxValor := maxFunc(2, 4, 7, 6, 9, 23, 15, 8, 1)
	medValor := medFunc(2, 4, 7, 6, 9, 23, 15, 8, 1)
	minValor := minFunc(2, 4, 7, 6, 9, 23, 15, 8, 1)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {
		fmt.Println("Maior valor:", maxValor)
		fmt.Printf("Média: %.2f\n", medValor)
		fmt.Println("Menor valor:", minValor)
	}
}