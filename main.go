package main

import (
	"news/newsletter"

	"log"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := newsletter.NewDatabase()
	if err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
	db.RunMigrations()

	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.Use(static.Serve("/", static.LocalFile("./frontend/dist/", true))) // Serve the frontend files

	router.POST("/file/upload", newsletter.UploadFile)
	router.Use(static.Serve("/file", static.LocalFile("/tmp/", false)))

	newsLetterApi := router.Group("newsletter")
	{
		newsLetterApi.GET("/", newsletter.GetAll)
		newsLetterApi.POST("/create", newsletter.Create)
		newsLetterApi.POST("/:id/send", newsletter.Send)
	}

	port := ":8080"
	log.Println("Listening on http://localhost" + port)
	router.Run(port)
}
