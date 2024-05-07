package main

import (
	"github.com/GMcD/api-semaphore/api"
)

// Initialize and Run App against Env Db
func main() {

	// Setup Env
	api.Env()

	// Setup App
	a := api.App{}
	a.Initialize()

	// Seed a few records
	AddProducts(a, 4)

	// Determine port for HTTP service.
	port := api.GetEnv("PORT", "8100")

	// Serve On PORT
	a.Run("0.0.0.0:" + port)
}
