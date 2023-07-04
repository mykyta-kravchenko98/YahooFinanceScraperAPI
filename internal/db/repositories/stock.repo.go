package repositories

import (
	"github.com/mykyta-kravchenko98/YahooFinanceScraperAPI/internal/models"

	"gorm.io/gorm"
)

type StockDataService interface {
	Create(stock models.Stock) error
	BatchCreate(stocks []models.Stock) error
	GetLastUpdatedData() ([]models.Stock, error)
}

// return StockDataService instance
func NewCurrencySnapshotDataService(db *gorm.DB) StockDataService {
	iDBSvc := &stockRepository{
		database: db,
	}
	return iDBSvc
}

// currencyRepository implements StockDataService
type stockRepository struct {
	database *gorm.DB
}

// Saving one stock
func (stockRepository *stockRepository) Create(stock models.Stock) error {
	result := stockRepository.database.Create(&stock)

	return result.Error
}

// Saving batch of stocks
func (stockRepository *stockRepository) BatchCreate(stocks []models.Stock) error {
	result := stockRepository.database.Create(&stocks)

	return result.Error
}

// Returning the last updated batch of stocks
func (stockRepository *stockRepository) GetLastUpdatedData() ([]models.Stock, error) {
	var stocks []models.Stock
	subquery := stockRepository.database.Table("stocks").Select("created_unix").Order("created_unix").Limit(1)
	result := stockRepository.database.Where("created_unix = (?)", subquery).Find(&stocks)

	return stocks, result.Error
}
