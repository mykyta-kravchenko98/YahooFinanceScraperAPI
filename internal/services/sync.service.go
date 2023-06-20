package services

import (
	"yahooscraper/internal/dataparser/yahoo"
	"yahooscraper/internal/db/repositories"
)

type syncService struct {
	stockRepository repositories.StockDataService
}

type SyncService interface {
	StartSync() error
}

func NewSyncService(stockRepository repositories.StockDataService) SyncService {
	return &syncService{stockRepository: stockRepository}
}

// Starting stocks data sync
func (syncService *syncService) StartSync() error {
	result, err := yahoo.Start()

	if err != nil {
		return err
	}

	err = syncService.stockRepository.BatchCreate(result)

	return err
}
