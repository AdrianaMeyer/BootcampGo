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
	return 0.0
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
	return alertaErro, errors.New(" - O tipo de animal informado é inválido")
}


func main() {
	animalCachorro, err := Animal(cachorro)
	animalGato, err := Animal(gato)
	animalHamster, err := Animal(hamster)
	animalTarantula, err := Animal(tarantula)

	animalInvalido, err := Animal("macaco")
	erro := animalInvalido(10)

	if err != nil {
		fmt.Println("Houve um erro:", erro, err)
	} 

	quantComidaCachorro := animalCachorro(5)
	quantComidaGato := animalGato(2)
	quantComidaHamster := animalHamster(3)
	quantComidaTarantula := animalTarantula(8)

	quantTotalKilos := quantComidaCachorro + quantComidaGato + quantComidaHamster + quantComidaTarantula

	fmt.Printf("Quantidade comida cachorro: %.3fKg\n", quantComidaCachorro)
	fmt.Printf("Quantidade comida gato: %.3fKg\n", quantComidaGato)
	fmt.Printf("Quantidade comida hasmter: %.3fKg\n", quantComidaHamster)
	fmt.Printf("Quantidade comida tarantula: %.3fKg\n", quantComidaTarantula)
	fmt.Printf("TOTAL: %.3fKg\n", quantTotalKilos)

}
