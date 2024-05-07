package api

import (
	"gorm.io/gorm"
)

// Product struct
type Product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Count Products
func (Product) Count(db *gorm.DB) int64 {
	var cnt int64 = 0
	db.Model(&Product{}).Count(&cnt)
	return cnt
}

// Soft Delete all Products
func (Product) SoftDeleteAll(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&Product{})
}

// Hard Delete all Products
func (Product) HardDeleteAll(db *gorm.DB) {
	db.Unscoped().Where("1 = 1").Delete(&Product{})
}

// Get Products from Database
func (p *Product) GetProducts(db *gorm.DB, limit int, offset int) []Product {
	var products []Product
	db.Limit(limit).Offset(offset).Find(&products)
	return products
}

// Get Product or nil from Database
func (p *Product) GetProduct(db *gorm.DB, id int) (*Product, error) {
	var product Product
	gdb := db.First(&product, "ID = ?", id)
	if gdb.Error != nil {
		return nil, gdb.Error
	} else {
		return &product, nil
	}
}

// Update Product in Database
func (p *Product) UpdateProduct(db *gorm.DB, product *Product) (*Product, error) {
	p.Name = product.Name
	p.Price = product.Price
	db.Save(p)
	return p.GetProduct(db, p.ID)
}

// Delete Product from Database
func (p *Product) DeleteProduct(db *gorm.DB) {
	db.Delete(p)
}

// Create Product in Database
func (p *Product) CreateProduct(db *gorm.DB) (*Product, error) {
	db.Create(p)
	return p.GetProduct(db, p.ID)
}
