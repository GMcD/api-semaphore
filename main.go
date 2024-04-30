package api

import (
	"log"
	"os"
)

// Initialize and Run App against Default Db
func main() {
	a := App{}
	a.Default()

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}

	a.Run(":" + port)
}
