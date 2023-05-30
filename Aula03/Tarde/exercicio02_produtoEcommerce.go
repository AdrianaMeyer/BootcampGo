package main

import (
	"fmt"
)

const (
	P = "pequeno"
	M = "medio"
	G = "grande"
)

type loja struct {
	listaProdutos []produto 
}

type produto struct {
	tipoProduto string
	nome string
	preco float64
}

func (p produto) CalcularCusto() float64 {

	if p.tipoProduto == "pequeno" {
		return p.preco
	} else if p.tipoProduto == "medio" {
		return p.preco + (p.preco * 0.03)
	} else {
		return p.preco + (p.preco * 0.06) + 2500
	}
}

func (l loja) Total() float64 {
	var total float64 = 0.0

	for _, produto := range l.listaProdutos {
		total = total + CalcularCusto()
	}
	return total
}

func (l loja) Adicionar(p produto) bool{
	novoProduto := novoProduto(p.tipoProduto, p.nome, p.preco)
	l := novoProduto 
	
	return true
}

type Produto interface {
	CalcularCusto() float64
}

type Ecommerce interface {
	Total() float64
	Adicionar() bool
}

func novoProduto(tipoProd string, nomeProd string, precoProd float64) produto {

	novoProduto := produto {
		tipoProduto: tipoProd,
		nome: nomeProd,
		preco: precoProd,
	}
	
	return novoProduto
}

func novaLoja(listaProd []produto) loja {

	novaLoja := loja {
		listaProdutos: listaprod,
	}
	return novaLoja
}


func main() {

	produtoCadastrado := novoProduto(P, "Camiseta", 35.00)	
	fmt.Println(produtoCadastrado)

	//EXERCICIO EM AJUSTE
}