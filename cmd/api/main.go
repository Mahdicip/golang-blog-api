package main

import (
	"log"

	"blog-api/config"
	"blog-api/internal/routes"
)

func main() {
	config.ConnectDatabase()

	router := routes.SetupRouter()

	log.Println("🚀 Server running on port 3000")
	router.Run(":3000")
}