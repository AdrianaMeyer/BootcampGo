package requisicoes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

var lastID int

func SaveUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req Users
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
