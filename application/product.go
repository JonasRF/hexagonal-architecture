package application

import "errors"

type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Desable() error
	GetId() string
	GetName() string
	GetStatus() string
	getPrice() float64
}

const (
	DISABLED = "disabled"
	ENABLED  = "enabled"
)

type Product struct {
	ID     string
	Name   string
	Price  float64
	Status string
}

//func (p *Product) IsValid() error {
//
//}

func (p *Product) Enable() error {
	if p.Price > 0 {
		p.Status = ENABLED
		return nil
	}
	return errors.New("the price must be greater than zero to enable the product")
}

//func (p *Product) Desable() error {
//
//}

func (p *Product) GetId() string {
	return p.ID
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetStatus() string {
	return p.Status
}

func (p *Product) getPrice() float64 {
	return p.Price
}
