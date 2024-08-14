package cli

import (
	"fmt"
	"github.com/codeedu/go-hexagonal/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with te price %f and status %s",
			product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())
	}
}
