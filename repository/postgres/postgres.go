package postgres

import (
	"github.com/jmoiron/sqlx"
	"github.com/nurfan27/era-product-service/model"
	"github.com/nurfan27/era-product-service/repository"
)

// RepositoryPostgres ...
type RepositoryPostgres struct {
	conn   *sqlx.DB
	dbName string
}

// NewRepositoryPostgres ...
func NewRepositoryPostgres(conn *sqlx.DB) repository.Postgres {
	return &RepositoryPostgres{
		conn: conn,
	}
}

func (ps RepositoryPostgres) AddProduct(params model.Product) (model.Product, error) {

	sqlStatement := `
		INSERT INTO products (name, price, description, quantity)
		VALUES ($1, $2, $3, $4)
		RETURNING id`

	id := 0
	err := ps.conn.QueryRow(sqlStatement, params.Name, params.Price, params.Description, params.Quantity).Scan(&id)
	if err != nil {
		return params, err
	}

	params.ID = int64(id)

	return params, nil
}

func (ps RepositoryPostgres) ListProduct(param string) ([]model.Product, error) {

	sqlStatement := "SELECT id, name, price, description, quantity FROM products "
	var products []model.Product

	err := ps.conn.Select(&products, sqlStatement)
	if err != nil {
		return products, err
	}

	return products, nil
}
