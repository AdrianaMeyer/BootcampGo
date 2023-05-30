package main

import (
	"fmt"
)

type produto struct {
	nome string
	preco float64
	quantidade int
}

type usuario struct {
	nome string
	sobrenome string
	email string
	produtos []produto
}

func novoProduto(nome string, preco float64) produto {
	novoProduto := produto{
		nome: nome,
		preco: preco,
	}

	return novoProduto
}

func novoUsuario(nome string, sobrenome string, email string) usuario {
	novoUsuario := usuario{
		nome: nome,
		sobrenome: sobrenome,
		email: email,
	}

	return novoUsuario
}

func (u *usuario) adicionarProduto(p produto, quant int) {
	p.quantidade = quant
	u.produtos = append(u.produtos, p)
}

func (u *usuario) excluirProduto() {
	u.produtos = nil
}

func main() {

	println("------- PRODUTOS CRIADOS -------")
	produto1:= novoProduto("cafeteira", 129.90)
	fmt.Println(produto1)

	produto2:= novoProduto("playstation", 4500.50)
	fmt.Println(produto2)

	produto3:= novoProduto("tenis", 340.00)
	fmt.Println(produto3)

	println("------- USUARIO CRIADO -------")
	usuario1:= novoUsuario("Maria", "Silva", "maria.silva@email.com")
	fmt.Println(usuario1)

	println("------- ADICIONANDO PRODUTOS -------")
	usuario1.adicionarProduto(produto1, 32)
	fmt.Println(usuario1)

	usuario1.adicionarProduto(produto2, 4)
	fmt.Println(usuario1)

	usuario1.adicionarProduto(produto3, 2)
	fmt.Println(usuario1)

	println("------- EXCLUINDO PRODUTOS -------")
	usuario1.excluirProduto()
	fmt.Println(usuario1)
}