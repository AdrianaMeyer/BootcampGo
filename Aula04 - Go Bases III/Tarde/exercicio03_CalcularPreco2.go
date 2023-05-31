package main

import (
	"fmt"
)

type produto struct {
	nome string
	preco float64
	quantidade float64
}

type servico struct {
	nome string
	preco float64
	minTrab float64
}


type manutencao struct {
	nome string
	preco float64
}


func novoProduto(nome string, preco float64, quantidade float64) produto {
	novoProduto := produto{
		nome: nome,
		preco: preco,
		quantidade: quantidade,
	}
	
	return novoProduto
}

func novoServico(nome string, preco float64, minTrab float64) servico {
	novoServico:= servico{
		nome: nome,
		preco: preco,
		minTrab: minTrab,
	}
	
	return novoServico
}


func novaManutencao(nome string, preco float64) manutencao {
	novaManutencao := manutencao{
		nome: nome,
		preco: preco,
	}
	
	return novaManutencao
}


func somarProdutos(produtos []produto) float64 {
	precoTotal := 0.0

	for _, p := range produtos {
		precoTotal = precoTotal + (p.preco * p.quantidade)
	}
	return precoTotal
}

/* AJUSTAR 

func somarServicos(servicos []servico) float64 {
	precoTotal := 0.0

	for _, s := range servicos {
		minResto := s.minTrab % 60

		if minResto < 30 {
			horasResto = 0.5
		}

		horasTrab = (s.minTrab - minResto) / 60
		horasTotais := horasTrab + horasResto

		precoTotal = precoTotal + (s.preco * horasTotais)
	}

	return precoTotal
}

*/
 
func somarManutencao(manutencoes []manutencao) float64 {
	precoTotal := 0.0

	for _, m := range manutencoes {
		precoTotal = precoTotal + m.preco
	}
	return precoTotal
}


func main() {

	println("------- PRODUTOS CRIADOS -------")
	produto1:= novoProduto("cafeteira", 129.90, 4)
	produto2:= novoProduto("playstation", 4500.50, 20)
	fmt.Println(produto1, produto2)

	println("------- SERVICOS CRIADOS -------")
	servico1:= novoServico("instalacao", 50.0, 45)
	servico2:= novoServico("pintura", 310.0, 80)
	fmt.Println(servico1, servico2)
	
	println("------- MANUTENCOES CRIADAS -------")
	manutencao1:= novaManutencao("reparo", 230.0)
	manutencao2:= novaManutencao("troca peca", 78.0)
	fmt.Println(manutencao1, manutencao2)

	sliceProdutos := []produto{produto1, produto2}
	//sliceServicos := []servico{servico1, servico2}
	sliceManutencao := []manutencao{manutencao1, manutencao2}

	totalProd := somarProdutos(sliceProdutos)
	//totalServ := somarServicos(sliceServicos)
	totalManut := somarManutencao(sliceManutencao)

	//total := totalProd + totalServ + totalManut

	fmt.Printf("O valor total dos produtos é: R$ %.2f\n", totalProd)
	fmt.Printf("O valor total das manutencoes é: R$ %.2f\n ", totalManut)

}