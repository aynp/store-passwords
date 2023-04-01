package main

import (
	"fmt"
	"github.com/aynp/store-passwords/src/config"
	"github.com/aynp/store-passwords/src/server"
)

func init() {
	config.SetConfig()
}

func main() {
	fmt.Println("Hello World")
	server.Serve()
}
