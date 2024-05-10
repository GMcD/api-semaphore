package main

import (
	"github.com/GMcD/api-semaphore/module/api"
)

// Initialize and Run App against Env Db
func main() {

	// Setup Env
	api.Env()

	// Setup App
	a := api.App{}
	a.Initialize()

	// Determine port for HTTP service.
	port := api.GetEnv("PORT", "8100")

	// Serve On PORT
	a.Run("0.0.0.0:" + port)
}
