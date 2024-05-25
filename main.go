package main

import (
	"log"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8080"

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.StaticFS("/", http.Dir("./frontend/dist/"))

	log.Print("Listening on http://localhost" + port)
	router.Run(port)
}
