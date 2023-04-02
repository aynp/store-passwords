package server

import (
	"github.com/gin-gonic/gin"
	"log"
)

func Serve() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	router.GET("/ping", ping)

	router.GET("/login", loginPage)
	router.POST("/login", login)

	router.GET("/signup", signupPage)
	router.POST("/signup", signup)

	if err := router.Run(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
