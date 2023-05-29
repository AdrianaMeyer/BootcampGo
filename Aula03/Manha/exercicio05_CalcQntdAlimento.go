package main

import (
	"fmt"
	"errors"
)

const (
	cachorro = "cachorro"
	gato = "gato"
	hamster = "hamster"
	tarantula = "tarantula"
)


func quantComidaCachorro(quantAnimal float64) float64 {
	quantComida := quantAnimal * 10
	return quantComida
}

func quantComidaGato(quantAnimal float64) float64 {
	quantComida := quantAnimal * 5
	return quantComida
}

func quantComidaHamster(quantAnimal float64) float64 {
	quantComida := quantAnimal * 0.250
	return quantComida
}

func quantComidaTarantula(quantAnimal float64) float64 {
	quantComida := quantAnimal * 0.150
	return quantComida
}

func alertaErro(quantAnimal float64) float64 {
	return 0
}

func Animal(tipoAnimal string) (func(quantAnimal float64) float64, error) {
	switch tipoAnimal {
	case "cachorro":
		return quantComidaCachorro, nil
	case "gato":
		return quantComidaGato, nil
	case "hamster": 
		return quantComidaHamster, nil
	case "tarantula": 
		return quantComidaTarantula, nil
	}
	return alertaErro, errors.New("Tipo de animal inválido")
}


func main() {
	animalCachorro, err := Animal(cachorro)
	animalGato, err := Animal(gato)
	animalHamster, err := Animal(hamster)
	animalTarantula, err := Animal(tarantula)

	if err != nil {
		fmt.Println("Erro:", err)
	} else {

		quantComidaCachorro := animalCachorro(5)
		quantComidaGato := animalGato(2)
		quantComidaHamster := animalHamster(3)
		quantComidaTarantula := animalTarantula(8)

		quantTotalKilos := quantComidaCachorro + quantComidaGato + quantComidaHamster + quantComidaTarantula

		fmt.Printf("Quantidade comida cachorro em Kg: %.3f\n", quantComidaCachorro)
		fmt.Printf("Quantidade comida gato em Kg: %.3f\n", quantComidaGato)
		fmt.Printf("Quantidade comida hasmter em Kg: %.3f\n", quantComidaHamster)
		fmt.Printf("Quantidade comida tarantula em Kg: %.3f\n", quantComidaTarantula)
		fmt.Printf("TOTAL em Kg: %.3f\n", quantTotalKilos)
	}

}