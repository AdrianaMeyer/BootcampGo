package requisicoes

import (
	"net/http"
	"github.com/gin-gonic/gin"

)

func Ola(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"mensagem": "Olá, Adriana!",
	})
}
