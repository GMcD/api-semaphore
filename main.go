package main

import (
	"log"
	"os"

	"github.com/GMcD/api-semaphore/api"
)

// Initialize and Run App against Default Db
func main() {
	a := api.App{}
	a.Default()

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	a.Run(":" + port)
}
