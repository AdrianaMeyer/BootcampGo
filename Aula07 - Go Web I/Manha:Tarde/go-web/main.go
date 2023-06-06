package main

import (
	"github.com/gin-gonic/gin"
	"github.com/AdrianaMeyer/BootcampGo/exibirInfos"
)

func main() {

	router := gin.Default()

	group := router.Group("/")
	{
		group.GET("/", exibirInfos.GetByQuery)
		group.GET("/Ola", exibirInfos.Ola)
		group.GET("/users", exibirInfos.GetAll)
		group.GET("/users/:Id", exibirInfos.GetById)
	}
	router.Run()

}