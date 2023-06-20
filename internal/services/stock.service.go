package services

import (
	"yahooscraper/internal/db/repositories"
	"yahooscraper/internal/models"
)

type stockService struct {
	stockRepository repositories.StockDataService
}

type StockService interface {
	GetActuallyStocks() ([]models.Stock, error)
}

func NewStockService(stockRepository repositories.StockDataService) StockService {
	return &stockService{stockRepository: stockRepository}
}

// Returning actually batch of stocks
func (stockService *stockService) GetActuallyStocks() ([]models.Stock, error) {
	stocks, err := stockService.stockRepository.GetLastUpdatedData()
	if err != nil {
		return nil, err
	}

	return stocks, nil
}
