package main

import (
	"math/rand"
	"fmt"
	"time"
)


func main() {
	array1 := rand.Perm(100)
	array2 := rand.Perm(1000)
	array3 := rand.Perm(10000)

	fmt.Printf("\nARRAY DE 100 ELEMENTOS\n")

	fmt.Println("---- Ordenacao por Insercao ----")
	inicio := time.Now()
	ordenaInsercao(array1)
	fim := time.Now()
	duracao := fim.Sub(inicio)

	fmt.Println("Array ordenado. Tempo de execucao: ", duracao)

	fmt.Println("---- Ordenacao por Selecao ----")
	inicio = time.Now()
	ordenaSelecao(array1)
	fim = time.Now()
	duracao = fim.Sub(inicio)

	fmt.Println("Array ordenado. Tempo de execucao: ", duracao)


	fmt.Printf("\nARRAY DE 1000 ELEMENTOS\n")

	fmt.Println("---- Ordenacao por Insercao ----")
	inicio = time.Now()
	ordenaInsercao(array2)
	fim = time.Now()
	duracao = fim.Sub(inicio)

	fmt.Println("Array ordenado. Tempo de execucao: ", duracao)

	fmt.Println("---- Ordenacao por Selecao ----")
	inicio = time.Now()
	ordenaSelecao(array2)
	fim = time.Now()
	duracao = fim.Sub(inicio)

	fmt.Println("Array ordenado. Tempo de execucao: ", duracao)


	fmt.Printf("\nARRAY DE 10000 ELEMENTOS\n")

	fmt.Println("---- Ordenacao por Insercao ----")
	inicio = time.Now()
	ordenaInsercao(array3)
	fim = time.Now()
	duracao = fim.Sub(inicio)

	fmt.Println("Array ordenado. Tempo de execucao: ", duracao)

	fmt.Println("---- Ordenacao por Selecao ----")
	inicio = time.Now()
	ordenaSelecao(array3)
	fim = time.Now()
	duracao = fim.Sub(inicio)

	fmt.Println("Array ordenado. Tempo de execucao: ", duracao)


	//TODO: IMPLEMENTAR ORDENACAO POR GRUPO

}


func ordenaInsercao(array []int) {

	for i := 1; i < len(array); i++ {
		atual := array[i]
		anterior := i - 1
		for anterior >= 0 && array[anterior] > atual {
			array[anterior+1] = array[anterior]
			anterior--
		}
		array[anterior+1] = atual
	}

}

func ordenaSelecao(array []int) {
	for i := 0; i < len(array)-1; i++ {
		minIndex := i

		for j := i + 1; j < len(array); j++ {
			if array[j] < array[minIndex] {
				minIndex = j
			}
		}

		array[i], array[minIndex] = array[minIndex], array[i]
	}
}