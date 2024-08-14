package main

import (
	"database/sql"
	db3 "github.com/codeedu/go-hexagonal/adapters/db"
	"github.com/codeedu/go-hexagonal/application"
)

func main() {

	db, _ := sql.Open("sqlite3", "db.sqlite")
	productDbAdapter := db3.NewProductDb(db)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product Exemplo", 30)

	productService.Enable(product)
}
