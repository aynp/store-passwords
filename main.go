package main

import (
	"fmt"
	"github.com/aynp/store-passwords/src/config"
)

func init() {
	config.SetConfig()
}

func main() {
	fmt.Println("Hello World")
}
