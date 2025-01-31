package main

import (
	"fmt"
	"os"

	"github.com/almarino_meli/grupo-5-wave-15/cmd/server"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	// app
	// - config
	cfg := &server.ConfigServerChi{
		ServerAddress:            os.Getenv("SERVER_ADDR"),
		ProductsLoaderFilePath:   os.Getenv("PRODUCTS_FILE"),
		SellersLoaderFilePath:    os.Getenv("SELLERS_FILE"),
		BuyersLoaderFilePath:     os.Getenv("BUYERS_FILE"),
		EmployeesLoaderFilePath:  os.Getenv("EMPLOYEES_FILE"),
		SectionsLoaderFilePath:   os.Getenv("SECTIONS_FILE"),
		WarehousesLoaderFilePath: os.Getenv("WAREHOUSES_FILE"),
	}

	app := server.NewServerChi(cfg)
	// - run
	if err := app.Run(); err != nil {
		fmt.Println(err)
		return
	}
}