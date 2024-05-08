package api

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

var once sync.Once

// Grab Environment once
func Env() {
	once.Do(func() {
		// Find .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Panicf("Error loading .env file: %s", err)
		}
	})
}

// Provide Default for Env Var
func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// Publish the method
func GetEnv(key, fallback string) string {
	return getEnv(key, fallback)
}

// Get Database Connection String
// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=5432 sslmode=disable TimeZone=Europe/London"

func GetDsn() string {
	host := GetEnv("APP_DB_HOST", "0.0.0.0")
	port := GetEnv("APP_DB_PORT", "5432")
	user := GetEnv("APP_DB_USERNAME", "postgres")
	password := GetEnv("APP_DB_PASSWORD", "postgres")
	dbname := GetEnv("APP_DB_NAME", "postgres")
	connectionString :=
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	log.Print(connectionString)
	return connectionString
}
