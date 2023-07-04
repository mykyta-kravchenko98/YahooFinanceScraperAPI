package controllers

import (
	"net/http"

	"github.com/mykyta-kravchenko98/YahooFinanceScraperAPI/internal/models"
	"github.com/mykyta-kravchenko98/YahooFinanceScraperAPI/internal/services"
	"github.com/mykyta-kravchenko98/YahooFinanceScraperAPI/utils"

	"github.com/labstack/echo/v4"
)

type StockController struct {
	stockService services.StockService
}

func NewStockController(svc services.StockService) *StockController {
	return &StockController{
		stockService: svc,
	}
}

// Get godoc
// @Summary Get actually stocks
// @Description Get an actually stocks. Auth not required
// @ID get
// @Tags stocks
// @Accept  json
// @Produce  json
// @Success 200 {object} []StockResponceDto
// @Failure 400 {object} utils.Error
// @Failure 500 {object} utils.Error
// @Router /stocks/ [get]
func (stockController *StockController) Get(ctx echo.Context) error {
	result, err := stockController.stockService.GetActuallyStocks()
	if err != nil {
		return ctx.JSON(http.StatusNotFound, utils.NotFound())
	}

	dtos := make([]models.StockResponceDto, 0, len(result))
	for _, item := range result {
		dtos = append(dtos, item.ToDTO())
	}

	return ctx.JSON(http.StatusOK, dtos)
}
