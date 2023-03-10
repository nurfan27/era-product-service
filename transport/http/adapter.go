package http

import (
	"context"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/nurfan27/era-product-service/action"
	"github.com/nurfan27/era-product-service/model"
	"github.com/nurfan27/era-product-service/repository"
)

type Adapter struct {
	db repository.Postgres
}

func (adp *Adapter) AddProduct(c echo.Context) error {
	ctx := context.Background()

	req := new(model.Product)
	err := c.Bind(req)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	act := action.NewCreate(adp.db)
	result, _ := act.Handle(ctx, *req)
	var resp Response

	data := map[string]interface{}{
		"id":          result.ID,
		"name":        result.Name,
		"price":       result.Price,
		"description": result.Description,
		"quantity":    result.Quantity,
	}

	resp.SetSuccessResponse(http.StatusOK, data)
	return c.JSON(resp.Code, resp)
}

func (adp *Adapter) ListProduct(c echo.Context) error {
	ctx := context.Background()

	req := c.QueryParam("sort")

	act := action.NewList(adp.db)
	result, _ := act.Handle(ctx, req)
	var resp Response
	var products []map[string]interface{}

	for _, v := range result {
		product := map[string]interface{}{
			"id":          v.ID,
			"name":        v.Name,
			"price":       v.Price,
			"description": v.Description,
			"quantity":    v.Quantity,
		}
		products = append(products, product)
	}

	data := map[string]interface{}{
		"products": products,
	}

	resp.SetSuccessResponse(http.StatusOK, data)
	return c.JSON(resp.Code, resp)
}

func NewAdapter(postgresRepo repository.Postgres) *Adapter {
	return &Adapter{
		db: postgresRepo,
	}
}
