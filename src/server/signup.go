package server

import (
	"github.com/aynp/store-passwords/src/database"
	"github.com/aynp/store-passwords/src/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func signupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", nil)
}

func signup(c *gin.Context) {
	name := c.PostForm("name")
	username := c.PostForm("username")
	password := c.PostForm("password")

	//hashedPass := hash(password);

	user := models.User{
		Name:     name,
		Username: username,
		Password: password,
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"userId": user.Id,
	})
}
