package handlers

import (
	"net/http"
	"user-management-api/models"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	//id := c.Param("id")
	// Fetch user by ID from database
	user := models.User{ID: 1, Name: "Aladin ka gin"}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save user to database
	c.JSON(http.StatusCreated, user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save user to database
	c.JSON(http.StatusCreated, user)
}
