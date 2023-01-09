package cli

import (
	"fmt"

	"github.com/XavierCabeto/takeaway/application"
)

func Run(service application.ProductServiceInterface, action string, productId string, productName string, price float64) (string, error) {

	var result = ""

	switch action {
	case "create":
		product, err := service.Create(productName, price)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID %s with the name %s has been created with the price %f",
			product.GetID(), product.GetName(), product.GetPrice())
	case "get":
		res, err := service.Get(productId)
		if err != nil {
			return result, err
		}
		result = fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f",
			res.GetID(), res.GetName(), res.GetPrice())
	default:
		res, err := service.GetAll()
		if err != nil {
			return result, err
		}
		for _, product := range res {
			result += fmt.Sprintf("Product ID: %s\nName: %s\nPrice: %f\n",
				product.GetID(), product.GetName(), product.GetPrice())
		}
	}
	return result, nil
}
