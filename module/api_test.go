package main_test

import (
	"log"
	"testing"

	"github.com/GMcD/api-semaphore/module/api"
)

func TestOrmSetup(t *testing.T) {
	teardownSuite := SetupSuite(t)
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
	teardownSuite := SetupSuite(t)
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
	teardownSuite := SetupSuite(t)
	defer teardownSuite(t)

	item := &api.Item{}
	items := item.GetItems(o.DB, records, 0)

	if len(items) != records {
		t.Errorf("Should be %v Items, not %v.", records, len(items))
	}
}

func TestGetItem(t *testing.T) {
	teardownSuite := SetupSuite(t)
	defer teardownSuite(t)

	coat := &api.Item{}
	coat = coat.GetItem(o.DB, "Coat")

	if coat.Description != "Coat" {
		t.Errorf("Should be same as '%v', not '%v'.", coat.Description, "Coat")
	}

}

func TestCreateItem(t *testing.T) {
	teardownSuite := SetupSuite(t)
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
	teardownSuite := SetupSuite(t)
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
	teardownSuite := SetupSuite(t)
	defer teardownSuite(t)

	hat := &api.Item{Description: "Hat", Price: 98}
	hat.CreateItem(o.DB)

	hat.DeleteItem(o.DB)

	res := hat.FindItem(o.DB, hat.Description)

	if res.RowsAffected != 0 {
		t.Errorf("Item '%v' should have been deleted.", hat.Description)
	}
}
