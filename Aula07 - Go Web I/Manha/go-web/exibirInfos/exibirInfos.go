package exibirInfos

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"io/ioutil"
	"log"
)

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

func Ola(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"mensagem": "Olá, Adriana!",
	})
}

func GetAll(c *gin.Context) {
	filePath := "users.json"
	data, err := ioutil.ReadFile(filePath)
	
	if err != nil {
		log.Println("Erro ao ler o arquivo JSON:", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Erro ao ler o arquivo JSON"})
		return
	}

	var user []Users
	e := json.Unmarshal(data, &user)

	if e != nil {
		log.Println("Erro ao fazer o parse do JSON:", e)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter o JSON"})
		return
	}

	c.JSON(http.StatusOK, user)
}