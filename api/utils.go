package api

import (
	"fmt"
	"log"
	"os"
)

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
	user := GetEnv("APP_DB_USERNAME", "postgres")
	password := GetEnv("APP_DB_PASSWORD", "postgres")
	dbname := GetEnv("APP_DB_NAME", "postgres")
	connectionString :=
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbname)
	log.Print(connectionString)
	return connectionString
}
