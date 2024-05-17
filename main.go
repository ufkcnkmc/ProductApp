package main

import (
	"ProductApp/common/app"
	"ProductApp/common/postgresql"
	"ProductApp/controller"
	"ProductApp/persistence"
	"ProductApp/service"
	"context"
	"github.com/labstack/echo/v4"
)

func main() {
	ctx := context.Background()
	e := echo.New()
	configurationManager := app.NewConfigurationManager()

	dbPool := postgresql.GetConnectionPool(ctx, configurationManager.PostgreSqlConfig)

	productRepository := persistence.NewProductRepository(dbPool)

	productService := service.NewProductService(productRepository)

	productController := controller.NewProductController(productService)

	productController.RegisterRoutes(e)
	e.Start("localhost:8080")

}
