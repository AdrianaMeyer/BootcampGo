package main

import (
	"fmt"
)

type Produto struct {
	nome string
	preco float64
	quantidade int
}

type Servico struct {
	nome string
	preco float64
	minTrab int
}


type Manutencao struct {
	nome string
	preco float64
}


func novoProduto(nome string, preco float64, quantidade int) Produto {
	novoProduto := Produto{
		nome: nome,
		preco: preco,
		quantidade: quantidade,
	}
	
	return novoProduto
}

func novoServico(nome string, preco float64, minTrab int) Servico {
	novoServico:= Servico{
		nome: nome,
		preco: preco,
		minTrab: minTrab,
	}
	
	return novoServico
}


func novaManutencao(nome string, preco float64) Manutencao {
	novaManutencao := Manutencao{
		nome: nome,
		preco: preco,
	}
	
	return novaManutencao
}


func somarProdutos(produtos []Produto) float64 {
	precoTotal := 0.0

	for _, p := range produtos {
		precoTotal = precoTotal + (p.preco * float64(p.quantidade))
	}
	return precoTotal
}


func somarServicos(servicos []Servico) float64 {
	precoTotal := 0.0
	minutos := 0
	horasTrab := 0
	precoPorMin := 0.0

	for _, s := range servicos {
		minutos = s.minTrab

		for minutos > 60 {
			horasTrab = horasTrab + 1
			minutos = minutos - 60
		}
		
		if minutos < 30 {
			minutos = 30
		}

		precoPorMin = s.preco / float64(60) 
		precoTotal = precoTotal + (s.preco * float64(horasTrab)) + (precoPorMin * float64(minutos))

	}

	return precoTotal
}

 
func somarManutencao(manutencoes []Manutencao) float64 {
	precoTotal := 0.0

	for _, m := range manutencoes {
		precoTotal = precoTotal + m.preco
	}
	return precoTotal
}


func main() {

	println("------- PRODUTOS CRIADOS -------")
	produto1:= novoProduto("cafeteira", 129.90, 4)
	produto2:= novoProduto("playstation", 4500.50, 2)
	fmt.Println(produto1, produto2)

	println("------- SERVICOS CRIADOS -------")
	servico1:= novoServico("instalacao", 50.0, 130)
	servico2:= novoServico("pintura", 310.0, 80)
	fmt.Println(servico1, servico2)
	
	println("------- MANUTENCOES CRIADAS -------")
	manutencao1:= novaManutencao("reparo", 230.0)
	manutencao2:= novaManutencao("troca peca", 78.0)
	fmt.Println(manutencao1, manutencao2)

	sliceProdutos := []Produto{produto1, produto2}
	sliceServicos := []Servico{servico1, servico2}
	sliceManutencao := []Manutencao{manutencao1, manutencao2}

	totalProd := somarProdutos(sliceProdutos)
	totalServ := somarServicos(sliceServicos)
	totalManut := somarManutencao(sliceManutencao)

	total := totalProd + totalServ + totalManut

	fmt.Printf("O valor total dos produtos é: R$ %.2f\n", totalProd)
	fmt.Printf("O valor total das manutencoes é: R$ %.2f\n", totalManut)
	fmt.Printf("O valor total dos servicos é: R$ %.2f\n", totalServ)
	fmt.Printf("TOTAL GERAL: R$ %.2f\n ", total)

}