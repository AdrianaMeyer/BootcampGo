package requisicoes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"

)

func GetAll(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Token inválido - você não tem permissao para fazer esta solicitação",
		})
		return
	}
	err := LeDadosJson(c)

	if err != nil {
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetById(c *gin.Context) {
	token := c.GetHeader("token")
	if token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Token inválido - você não tem permissao para fazer esta solicitação",
		})
		return
	}

	err := LeDadosJson(c)
	
	if err != nil {
		log.Println(err)
	} 
		
	idParam := c.Param("Id")
	id, err := strconv.Atoi(idParam)
		
	if err != nil {
		log.Println("Erro ao converter o ID:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}
		
	for _, user := range users {
		if user.Id == id {
			c.JSON(http.StatusOK, user)
			return
		}
	}
		
	c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		
}
	

func GetByQuery(c *gin.Context){
	token := c.GetHeader("token")
	if token != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Token inválido - você não tem permissao para fazer esta solicitação",
		})
		return
	}

	nome := c.Query("nome")
	sobrenome := c.Query("sobrenome")
	email := c.Query("email")
	idade := c.Query("idade")
	altura := c.Query("altura")
	ativo := c.Query("ativo")
	dataCriacao := c.Query("dataCriacao")
	
	err := LeDadosJson(c)
	if err != nil {
		log.Println(err)
	} 

	resultado := []Users{}

	for _, user := range users {
    if nome != "" && user.Nome != nome {
        continue
    }
    if sobrenome != "" && user.Sobrenome != sobrenome {
        continue
    }
	if email != "" && user.Email != email {
        continue
    }
	if sobrenome != "" && user.Sobrenome != sobrenome {
        continue
    }
	if idade != "" && strconv.Itoa(user.Idade) != idade {
        continue
    }
	if altura != "" && strconv.FormatFloat(user.Altura, 'f', -1, 64) != altura {
        continue
    }
	if ativo != "" && user.Ativo != ativo {
        continue
    }
	if dataCriacao != "" && user.DataDeCriacao != dataCriacao {
        continue
    }

    resultado = append(resultado, user)
	}

	if len(resultado) > 0 {
		c.JSON(http.StatusOK, resultado)
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
	}
	
}