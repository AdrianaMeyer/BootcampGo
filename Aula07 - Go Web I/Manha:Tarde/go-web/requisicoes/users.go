package requisicoes 

type Users struct {
	Id 				int
    Nome 			string
    Sobrenome 		string
    Email 			string
	Idade 			int
    Altura 			float64
    Ativo 			string
    DataDeCriacao 	string
}

var users []Users