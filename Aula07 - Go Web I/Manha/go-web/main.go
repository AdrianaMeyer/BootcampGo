package main

import (
	"github.com/gin-gonic/gin"
	"github.com/AdrianaMeyer/BootcampGo/exibirInfos"
)

func main() {

	router := gin.Default()

	group := router.Group("/")
	{
		group.GET("/ola", exibirInfos.Ola)
		group.GET("/users", exibirInfos.GetAll)
	}
	router.Run()

}