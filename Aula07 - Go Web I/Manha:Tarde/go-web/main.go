package main

import (
	"github.com/gin-gonic/gin"
	"github.com/AdrianaMeyer/BootcampGo/requisicoes"
)

func main() {

	router := gin.Default()

	group := router.Group("/")
	{
		group.GET("/", requisicoes.GetByQuery)
		group.GET("/Ola", requisicoes.Ola)
		group.GET("/users", requisicoes.GetAll)
		group.GET("/users/:Id", requisicoes.GetById)

		group.POST("/salvar", requisicoes.SaveUser())
	}
	router.Run()

}