package requisicoes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func SaveUser() gin.HandlerFunc {
	return func(c *gin.Context) {
        var lastID int
		var req Users
        token := c.GetHeader("token")
		
		if token != "123456" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Error": "Token inválido - você não tem permissao para fazer esta solicitação",
			})
			return
		}

        if len(users) == 0 {
            lastID = 0
        } else {
            lastUser := len(users) -1
            lastID =  users[lastUser].Id     
        }

        err := c.ShouldBindJSON(&req);
		if  err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		lastID++
		req.Id = lastID
		users = append(users, req)

		c.JSON(http.StatusCreated, users)
	}
}
