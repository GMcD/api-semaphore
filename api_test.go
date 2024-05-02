package main_test

import (
	"log"
	"testing"

	"github.com/GMcD/api-semaphore/api"
)

// gorm database instance
var o api.Orm

// Number of Seed Records
const records = 3

// Set Up Gorm Database
func setupSuite(t *testing.T) func(t *testing.T) {
	t.Log("Setting up database for tests")
	o.SetupDb()

	// Seed Records
	coat := &api.Item{Description: "Coat", Price: 123}
	coat.CreateItem(o.DB)
	shoes := &api.Item{Description: "Shoes", Price: 234}
	shoes.CreateItem(o.DB)
	trousers := &api.Item{Description: "Trousers", Price: 150}
	trousers.CreateItem(o.DB)
	//	o.CreateItem(&api.Item{Description: "Trousers", Price: 150})

	// Close Connection
	return func(t *testing.T) {
		t.Log("Tearing down database")
		api.Item{}.HardDeleteAll(o.DB)
		// Probably Unneccesary
		sqlDB, err := o.DB.DB()
		if err != nil {
			log.Fatalln(err)
		}
		sqlDB.Close()
	}
}

func TestOrmSetup(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	desc := "Coat"
	var item api.Item
	i := item.GetItem(o.DB, desc)
	if i.Description != desc {
		t.Errorf("Failed to retrieve '%v', received '%v'", i.Description, desc)
	}

	log.Printf("Item: %v", i)
}

func TestSoftDeleteAll(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	cnt := api.Item{}.Count(o.DB)
	if cnt != records {
		t.Errorf("Should be %v Items, not %v.", records, cnt)
	}
	api.Item{}.SoftDeleteAll(o.DB)
	cnt = api.Item{}.Count(o.DB)
	if cnt != 0 {
		t.Errorf("Should be 0 Items, not %v.", cnt)
	}
}

func TestGetItems(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	item := &api.Item{}
	items := item.GetItems(o.DB, records, 0)

	if len(items) != records {
		t.Errorf("Should be %v Items, not %v.", records, len(items))
	}
}

func TestGetItem(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	coat := &api.Item{}
	coat = coat.GetItem(o.DB, "Coat")

	if coat.Description != "Coat" {
		t.Errorf("Should be same as '%v', not '%v'.", coat.Description, "Coat")
	}

}

func TestCreateItem(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	hat := &api.Item{Description: "Hat", Price: 98}
	hat.CreateItem(o.DB)

	saved := &api.Item{}
	saved = saved.GetItem(o.DB, hat.Description)

	if saved.Description != hat.Description {
		t.Errorf("Should be same as '%v', not '%v'.", hat.Description, saved.Description)
	}

}

func TestUpdateItem(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	coat := &api.Item{}
	coat = coat.GetItem(o.DB, "Coat")
	coat.Price = 1000

	coat.UpdateItem(o.DB, coat)

	expensive := &api.Item{}
	expensive = coat.GetItem(o.DB, "Coat")

	if expensive.Price != 1000 {
		t.Errorf("Should be same as '%v', not '%v'.", expensive.Price, 1000)
	}

}

func TestDeleteItem(t *testing.T) {
	teardownSuite := setupSuite(t)
	defer teardownSuite(t)

	hat := &api.Item{Description: "Hat", Price: 98}
	hat.CreateItem(o.DB)

	hat.DeleteItem(o.DB)

	res := hat.FindItem(o.DB, hat.Description)

	if res.RowsAffected != 0 {
		t.Errorf("Item '%v' should have been deleted.", hat.Description)
	}
}
