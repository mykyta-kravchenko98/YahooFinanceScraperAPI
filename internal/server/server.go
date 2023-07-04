package server

import (
	"sync"
	"yahooscraper/internal/configs"
	"yahooscraper/internal/controllers"
	"yahooscraper/internal/db/repositories"
	"yahooscraper/internal/models"
	"yahooscraper/internal/router"
	"yahooscraper/internal/services"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

var runOnce sync.Once

// Start server
func Init(db *gorm.DB) {
	config := configs.GetConfig()
	port := config.Server.RESTPort
	router := router.New()

	//DB AutoMigration
	db.AutoMigrate(
		&models.Stock{},
	)

	runOnce.Do(func() {
		registrate(router, db)
		router.Logger.Fatal(router.Start(":" + port))
	})
}

// registrate routes
func registrate(e *echo.Echo, db *gorm.DB) {
	rep := repositories.NewCurrencySnapshotDataService(db)
	serv := services.NewStockService(rep)
	stockController := controllers.NewStockController(serv)

	stocks := e.Group("/stocks")

	stocks.GET("", stockController.Get)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}
