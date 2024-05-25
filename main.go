package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	port := ":8080"

	router := gin.Default()
	router.StaticFS("/", http.Dir("./frontend/dist/"))

	log.Print("Listening on http://localhost" + port)
	router.Run(port)
}
