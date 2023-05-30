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

type Produto interface {
	CalcularCusto() float64
}

type Ecommerce interface {
	Total() float64
	Adicionar(p produto) bool
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
		listaProdutos: listaProd,
	}
	return novaLoja
}


func (p *produto) CalcularCusto() float64 {

	if p.tipoProduto == "pequeno" {
		return p.preco
	} else if p.tipoProduto == "medio" {
		return p.preco + (p.preco * 0.03)
	} else {
		return p.preco + (p.preco * 0.06) + 2500
	}
}

func (l *loja) Total() float64 {
	var total float64 = 0.0
	
	produtosDaLoja := l.listaProdutos

	for _, produto := range produtosDaLoja {
		total = total + produto.CalcularCusto()
	}
	return total
}

func (l *loja) Adicionar(p produto) bool{
	l.listaProdutos = append(l.listaProdutos, p)
	
	return true
}


func main() {

	produtoCadastrado1 := novoProduto(P, "Camiseta", 35.00)	
	fmt.Printf("Produto 1: %v\n", produtoCadastrado1)

	produtoCadastrado2 := novoProduto(M, "Mesa", 1200.00)	
	fmt.Printf("Produto 2: %v\n", produtoCadastrado2)

	produtoCadastrado3 := novoProduto(G, "Geladeira", 5300.00)	
	fmt.Printf("Produto 3: %v\n", produtoCadastrado3)

	produtoCadastrado4 := novoProduto(M, "Poltrona", 800.00)	
	fmt.Printf("Produto 4: %v\n", produtoCadastrado4)

	listaDeProdutos := []produto{produtoCadastrado1, produtoCadastrado2, produtoCadastrado3}

	lojaCadastrada := novaLoja(listaDeProdutos)
	fmt.Printf("Nova loja: %v\n", lojaCadastrada)

	fmt.Printf("Custo do produto P: %.2f\n", produtoCadastrado1.CalcularCusto())
	fmt.Printf("Custo do produto M: %.2f\n", produtoCadastrado2.CalcularCusto())
	fmt.Printf("Custo do produto G: %.2f\n", produtoCadastrado3.CalcularCusto())

	fmt.Printf("TOTAL DA LOJA: %.2f\n", lojaCadastrada.Total())

	fmt.Printf("CADASTRO DO NOVO PRODUTO REALIZADO? %t\n", lojaCadastrada.Adicionar(produtoCadastrado4))
	fmt.Printf("Estoque atualizado: %v\n", lojaCadastrada)
}
