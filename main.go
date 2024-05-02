package main

import (
	"github.com/GMcD/api-semaphore/api"
)

// Initialize and Run App against Default Db
func main() {
	a := api.App{}
	a.Initialize()

	// Determine port for HTTP service.
	port := api.GetEnv("PORT", "8100")

	a.Run(":" + port)
}
