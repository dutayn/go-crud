package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Category struct {
	gorm.Model
	Name     string    `json:"name"`
	Enable   bool      `json:"enable"`
	Products []Product `gorm:"ForeignKey:CategoryID" json:"categories"`
}

type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Enable      bool    `json:"enable"`
	CategoryID  uint    `json:"category_id"`
	Images      []Image `gorm:"ForeignKey:ProductID" json:"products"`
}

type Image struct {
	gorm.Model
	Name      string `json:"name"`
	File      string `json:"file"`
	Enable    bool   `json:"enable"`
	ProductID uint   `json:"product_id"`
}

// DBMigrate will create and migrate the tables, and then make the some relationships if necessary
func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Category{}, &Product{}, &Image{})
	db.Model(&Product{}).AddForeignKey("category_id", "categories(id)", "CASCADE", "CASCADE")
	db.Model(&Image{}).AddForeignKey("product_id", "products(id)", "CASCADE", "CASCADE")

	return db
}
