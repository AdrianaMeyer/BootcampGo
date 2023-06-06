package requisicoes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"

)

func GetAll(c *gin.Context) {
	err := LeDadosJson(c)

	if err != nil {
		log.Println(err)
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func GetById(c *gin.Context) {
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
	nome := c.Query("nome")
	sobrenome := c.Query("sobrenome")
	email := c.Query("email")
	idade := c.Query("idade")
	altura := c.Query("altura")
	ativo := c.Query("ativo")
	dataCriacao := c.Query("dataCriacao")
	var usersFilter []Users
	
	err := LeDadosJson(c)
	
	if err != nil {
		log.Println(err)
	} 

	for _, user := range users {

		if user.Nome == nome {
			usersFilter = append(usersFilter, user)
		}

		if user.Sobrenome == sobrenome {
			usersFilter = append(usersFilter, user)
		}

		if user.Email == email {
			usersFilter = append(usersFilter, user)
		}

		if strconv.Itoa(user.Idade) == idade {
			usersFilter = append(usersFilter, user)
		}

		if strconv.FormatFloat(user.Altura, 'f', -1, 64) == altura {
			usersFilter = append(usersFilter, user)
		}

		if strconv.FormatBool(user.Ativo) == ativo {
			usersFilter = append(usersFilter, user)
		}

		if user.DataDeCriacao == dataCriacao {
			usersFilter = append(usersFilter, user)
		}

	}

	if len(usersFilter) > 0 {
		c.JSON(http.StatusOK, usersFilter)
		return
	} else {
		c.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
	}
	
	
}