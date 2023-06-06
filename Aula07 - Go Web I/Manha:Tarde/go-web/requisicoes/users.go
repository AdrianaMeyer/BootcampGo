package requisicoes 

type Users struct {
	Id 				int
    Nome 			string
    Sobrenome 		string
    Email 			string
	Idade 			int
    Altura 			float64
    Ativo 			bool
    DataDeCriacao 	string
}

var users []Users