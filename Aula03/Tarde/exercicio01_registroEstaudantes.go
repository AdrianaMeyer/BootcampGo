package main

import (
	"fmt"
)

type Aluno struct {
	nome string 
	sobrenome string 
	rg string 
	dataAdmissao string 

}

func (aluno Aluno) detalhes() {
	fmt.Println("--------- Dados do Aluno ------- \t")
	fmt.Println("Nome:", aluno.nome)
	fmt.Println("Sobrenome:", aluno.sobrenome)  
	fmt.Println("RG:", aluno.rg)  
	fmt.Println("Data de Admissão:", aluno.dataAdmissao)       
}


func main() {
	aluno1 := Aluno {
		nome: "João",
		sobrenome: "Silva",
		rg: "11.234.567-8",
		dataAdmissao: "12/04/2010",
	}

	aluno2 := Aluno {
		nome: "Maria",
		sobrenome: "Ferreira",
		rg: "12.134.567-9",
		dataAdmissao: "01/02/2012",
	}

	aluno1.detalhes()
	aluno2.detalhes()
}