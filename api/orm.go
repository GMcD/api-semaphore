package api

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Orm struct {
	DB *gorm.DB
}

func (o *Orm) GetGormDb(dsn string) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	o.DB = db
}

// Create a Database Table for Items
func (o *Orm) SetupDb() {
	connectionString := GetDsn()
	o.GetGormDb(connectionString)

	o.DB.AutoMigrate(&Item{})
}
