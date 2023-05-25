package exercicio03

var (
	nome string = "José"
	sobrenome string = "Silva"
	idade int = 18
	licencaParaDirigir bool = true
	estaturaDaPessoa float32 = 1.65
	quantidadeDeFilhos int = 2
)

func Nome() string {
	return nome
}

func Sobrenome() string {
	return sobrenome
}

func Idade() int {
	return idade
}

func Dirige() bool {
	return licencaParaDirigir
}

func Estatura() float32 {
	return estaturaDaPessoa
}

func Filhos() int {
	return quantidadeDeFilhos
}


