package db

import (
	"database/sql"
	"github.com/codeedu/go-hexagonal/application"
)

type ProductDb struct {
	db *sql.DB
}

func (p *ProductDb) Get(id string) (application.ProductInterface, error) {
	var product application.Product
	stmp, err := p.db.Prepare("select id, name, price, status from products where id=?")
	if err != nil {
		return nil, err
	}
	err = stmp.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)
	if err != nil {
		return nil, err
	}
	return &product, nil
}
