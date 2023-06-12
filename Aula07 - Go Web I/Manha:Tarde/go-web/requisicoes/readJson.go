package requisicoes

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"os"
	"fmt"

)

func LeDadosJson(c *gin.Context) error {
	data, err := os.ReadFile("users.json")
	
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Erro ao ler o arquivo JSON"})
		return fmt.Errorf("Erro ao ler o arquivo JSON: %v", err)
	}
	e := json.Unmarshal(data, &users)

	if e != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter o JSON"})
		return fmt.Errorf("Erro ao fazer o parse do JSON: %v", e)
	}

	return nil

}