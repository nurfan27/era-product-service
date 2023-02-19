package action

import (
	"context"

	"github.com/nurfan27/era-product-service/util/errors"

	"github.com/nurfan27/era-product-service/model"
	"github.com/nurfan27/era-product-service/repository"
)

type List struct {
	e    errors.UniError
	repo repository.Postgres
}

func (c *List) Handle(ctx context.Context, req model.Product) (*model.Product, *errors.UniError) {
	var result model.Product

	result, err := c.repo.AddProduct(req)
	if err != nil {
		return nil, c.e.ErrProcess("error insert data to table product")
	}

	return &result, nil
}

func NewList(repo repository.Postgres) *List {
	return &List{
		repo: repo,
	}
}
