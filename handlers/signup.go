package handlers

import (
	"myserver/models"
	"myserver/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := utils.CreateUser(user); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": err.Error()})
		return
	}

	token, _ := utils.GenerateToken(user.Username)
	c.JSON(http.StatusOK, gin.H{"access_token": token})
}