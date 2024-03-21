package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vogonwann/gorm-gin/model"
	"github.com/vogonwann/gorm-gin/utils"
)

func Login(c *gin.Context) {
	var user model.User

	// Check user credentials and generate a JWT token
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	// Check if credentials are valid
	// TODO: replace the logic with real authentication
	if user.Username == "user" && user.Password == "password" {
		// Generate a JWT token
		token, err := utils.GenerateToken(user.Id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "error generating JWT token"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}
}

// Function for registering a new user
// TODO: replace the logic with real registration
func Register(c *gin.Context) {
	var user model.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid data"})
		return
	}

	// TODO: remember to securely hash password before storing them
	user.Id = 1 // just for demonstration purposes
	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}
