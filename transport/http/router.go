package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/nurfan27/era-product-service/repository"
)

func Serve(postgresRepo repository.Postgres) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	adapter := NewAdapter(postgresRepo)

	// Routes
	e.POST("/api/v1/products", adapter.AddProduct)

	// Start server
	e.Logger.Fatal(e.Start(":8000"))
}
