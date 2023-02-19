package main

import (
	"log"

	"github.com/nurfan27/era-product-service/repository/postgres"
	route "github.com/nurfan27/era-product-service/transport/http"
	db "github.com/nurfan27/era-product-service/util/database"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func main() {
	conn, err := db.GetSqlxConnection()
	if err != nil {
		log.Fatalf("failed to connect database")
	}
	defer conn.Close()

	postgresRepo := postgres.NewRepositoryPostgres(conn)

	route.Serve(postgresRepo)
}
