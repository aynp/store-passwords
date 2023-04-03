package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func loginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	fmt.Println(username, password)
	c.JSON(http.StatusOK, gin.H{
		"test": "test",
	})
}
