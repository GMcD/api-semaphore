package main_test

import (
	"log"
	"testing"

	"github.com/GMcD/api-semaphore/module/api"
	"github.com/google/uuid"
)

// gorm database instance
var o api.Orm

// Number of Seed Records
const records = 3

// Set Up Gorm Database -> Should be Migration
func SetupSuite(t *testing.T) func(t *testing.T) {
	// t.Log("Initializing App routes")
	a.Initialize()
	// t.Log("Setting up database for tests")
	o.SetupDb()

	// Seed Item Records
	coat := &api.Item{Description: "Coat", Price: 123}
	coat.CreateItem(o.DB)
	shoes := &api.Item{Description: "Shoes", Price: 234}
	shoes.CreateItem(o.DB)
	trousers := &api.Item{Description: "Trousers", Price: 150}
	trousers.CreateItem(o.DB)
	//	o.CreateItem(&api.Item{Description: "Trousers", Price: 150})

	// Seed Product Records
	u1, _ := uuid.NewRandom()
	car := &api.Product{ID: u1, Name: "Tesla", Price: 30000}
	car.CreateProduct(o.DB)
	u2, _ := uuid.NewRandom()
	lorry := &api.Product{ID: u2, Name: "Peterbilt", Price: 300000}
	lorry.CreateProduct(o.DB)

	// Close Connection
	return func(t *testing.T) {
		// t.Log("Tearing down database")
		api.Item{}.HardDeleteAll(o.DB)
		api.Product{}.HardDeleteAll(o.DB)
		// Probably Unneccesary
		sqlDB, err := o.DB.DB()
		if err != nil {
			log.Fatalln(err)
		}
		sqlDB.Close()
	}
}
