package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func signupPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func signup(c *gin.Context) {
	c.String(http.StatusCreated, "DONE")
}
