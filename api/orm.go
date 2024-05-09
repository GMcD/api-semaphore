package api

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Orm struct {
	DB *gorm.DB
}

// Open Database from DSN
func (o *Orm) GetGormDb(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	o.DB = db
}

// Create Database Tables for Items and Products
func (o *Orm) SetupDb() {
	connectionString := GetDsn()
	o.GetGormDb(connectionString)

	o.DB.Migrator().DropTable(&Item{}, &Product{})
	o.DB.AutoMigrate(&Item{}, &Product{})
}
