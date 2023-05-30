package main

import (
	"fmt"
)

type usuarios struct {
	nome string
	sobrenome string
	idade int
	email string
	senha string
}

func AtualizarNome(usuario *usuarios, novoNome string, novoSobrenome string) {
	usuario.nome = novoNome
	usuario.sobrenome = novoSobrenome
}

func AtualizarIdade(usuario *usuarios, novaIdade int) {
	usuario.idade = novaIdade
}

func AtualizarEmail(usuario *usuarios, novoEmail string) {
	usuario.email = novoEmail
}

func AtualizarSenha(usuario *usuarios, novoSenha string) {
	usuario.senha = novoSenha
}

func main() {

	usuario1 := usuarios{
		nome: "Jose",
		sobrenome: "Silva",
		idade: 15,
		email: "jose.ze@email.com",
		senha: "senha123",
	}

	fmt.Println("\n----------- Dados Originais --------")
	fmt.Println(usuario1)

	AtualizarNome(&usuario1, "Zezinho", "Silva e Silva")
	AtualizarIdade(&usuario1, 28)
	AtualizarEmail(&usuario1, "jose.zezinho@email.com")
	AtualizarSenha(&usuario1, "Senha12345")
	fmt.Println("\n----------- Dados Atualizados --------")
	fmt.Println(usuario1)
}