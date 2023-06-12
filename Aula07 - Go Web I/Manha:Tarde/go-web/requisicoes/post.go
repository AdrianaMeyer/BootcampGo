package requisicoes

import (
	"net/http"
	"github.com/gin-gonic/gin"
    "fmt"
    "log"
)

func SaveUser() gin.HandlerFunc {
	return func(c *gin.Context) {
        token := c.GetHeader("token")
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token inválido - você não tem permissao para fazer esta solicitação",
			})
			return
		}

        var req Users
        err := c.ShouldBindJSON(&req);
		if  err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

        invalidFields := validateFields(req)
		if invalidFields != nil {
			message := fmt.Sprintf("Existem campos obrigatórios que devem ser preenchidos: %v", invalidFields)
			c.JSON(http.StatusBadRequest, message)
			return
		}

        lastID, err := LastID(c)
        if err != nil {
            log.Println(err)
        }
		lastID++
		req.Id = lastID
		users = append(users, req)

	    if err != nil {
		fmt.Println("Erro ao salvar as informacoes no arquivo", err)
	    }

		c.JSON(http.StatusCreated, users)
	}
}

func LastID(c *gin.Context) (int, error) {

    err := LeDadosJson(c)
	if err != nil {
		return 0, err
	} 

	if len(users) == 0 {
		return 0, nil
	}

	lastUser := users[len(users)-1]
	return lastUser.Id, nil

}

func validateFields(req Users) []string {

	emptyfields := []string{}

	if req.Nome == "" {
		emptyfields = append(emptyfields, "Nome")
	}

	if req.Sobrenome == "" {
		emptyfields = append(emptyfields, "Sobrenome")
	}

	if req.Email == "" {
		emptyfields = append(emptyfields, "Email")
	}

	if req.Idade == 0 {
		emptyfields = append(emptyfields, "Idade")
	}

	if req.Altura == 0 {
		emptyfields = append(emptyfields, "Altura")
	}

    if req.Ativo == "" {
		emptyfields = append(emptyfields, "Ativo")
	}

	if req.DataDeCriacao == "" {
		emptyfields = append(emptyfields, "Data de Criação")
	}

	if len(emptyfields) != 0 {
		return emptyfields
	} else {
		return nil
	}
}