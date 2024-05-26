package main

import (
	"news/newsletter"

	"log"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

func main() {
	port := ":8080"

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	router.StaticFS("/", http.Dir("./frontend/dist/")) // Serve the frontend files

	router.POST("/file/upload", newsletter.FileUpload)

	newsLetterApi := router.Group("newsletter")
	{
		newsLetterApi.POST("/create", newsletter.Create)
	}

	log.Println("Listening on http://localhost" + port)
	router.Run(port)
}
