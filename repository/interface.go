package repository

import (
	"github.com/nurfan27/era-product-service/model"
)

type Postgres interface {
	AddProduct(params model.Product) (model.Product, error)
	ListProduct(param string) ([]model.Product, error)
}
