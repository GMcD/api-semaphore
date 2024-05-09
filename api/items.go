package api

import "gorm.io/gorm"

// Item Struct
type Item struct {
	gorm.Model
	Description string `gorm:"index" json:"description"`
	Price       uint   `json:"price"`
}

// Count Items
func (Item) Count(db *gorm.DB) int64 {
	var cnt int64 = 0
	db.Model(&Item{}).Count(&cnt)
	return cnt
}

// Soft Delete all Items
func (Item) SoftDeleteAll(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Item{})
}

// Hard Delete all Items
func (Item) HardDeleteAll(db *gorm.DB) {
	db.Unscoped().Where("1 = 1").Delete(&Item{})
}

// Get Items from Database
func (i *Item) GetItems(db *gorm.DB, limit int, offset int) []Item {
	var items []Item
	db.Limit(limit).Offset(offset).Find(&items)
	return items
}

// Get Item from Database
func (i *Item) GetItem(db *gorm.DB, desc string) *Item {
	var item Item
	db.First(&item, "description = ?", desc)
	return &item
}

// Find Item if present in Database
func (i *Item) FindItem(db *gorm.DB, desc string) *gorm.DB {
	var item Item
	return db.Where("description = ?", desc).Find(&item)
}

// Create Item in Database
func (i *Item) CreateItem(db *gorm.DB) *Item {
	db.Create(i)
	return i.GetItem(db, i.Description)
}

// Update Item in Database
func (i *Item) UpdateItem(db *gorm.DB, item *Item) *Item {
	i.Description = item.Description
	i.Price = item.Price
	db.Save(i)
	return i.GetItem(db, i.Description)
}

// Delete Item in Database
func (i *Item) DeleteItem(db *gorm.DB) {
	db.Delete(i)
}
