package main

import (
	// "fmt"
	// "github.com/gin-gonic/gin"
	"goblog/router"
)

func main() {
	r := router.NewRouter()

	r.Run(":8080")
}
