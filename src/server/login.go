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
	uname, pass := c.PostForm("username"), c.PostForm("password")
	fmt.Println(uname, pass)
	c.JSON(http.StatusOK, gin.H{
		"test": "test",
	})
}
