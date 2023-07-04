package models

import (
	"time"

	pb "yahooscraper/pkg/yahooScraper_v1"

	"gorm.io/gorm"
)

// GORM DB Entity
type Stock struct {
	gorm.Model

	Id                 uint `gorm:"primaryKey"`
	Symbol             string
	Name               string
	Price              float64
	Change             float64
	PercentChange      float64
	Volume             string
	AvgVolumeFor3Month string
	MarketCap          string
	PERatio            string
	CreatedAt          time.Time
	CreatedUnix        int64 `gorm:"autoCreateTime"`
}

// DTO model
// StockResponceDto  godoc
type StockResponceDto struct {
	Symbol             string  `json:"symbol"`
	Name               string  `json:"name"`
	Price              float64 `json:"price"`
	Change             float64 `json:"change"`
	PercentChange      float64 `json:"percent_change"`
	Volume             string  `json:"volume"`
	AvgVolumeFor3Month string  `json:"avarage_3_month_volume"`
	MarketCap          string  `json:"market_cap"`
	PERatio            string  `json:"pe_ratio"`
	CreatedAt          string  `json:"created_at_date_utc"`
	CreatedUnix        int64   `json:"created_at_unix"`
} // @name StockResponceDto

func (stock *Stock) ToDTO() StockResponceDto {
	dto := StockResponceDto{}

	dto.Symbol = stock.Symbol
	dto.Name = stock.Name
	dto.Price = stock.Price
	dto.Change = stock.Change
	dto.PercentChange = stock.PercentChange
	dto.Volume = stock.Volume
	dto.AvgVolumeFor3Month = stock.AvgVolumeFor3Month
	dto.MarketCap = stock.MarketCap
	dto.PERatio = stock.PERatio
	dto.CreatedAt = stock.CreatedAt.UTC().String()
	dto.CreatedUnix = stock.CreatedUnix

	return dto
}

func (stock *Stock) ProtoStock() *pb.Stocks {
	ps := pb.Stocks{}

	ps.Symbol = stock.Symbol
	ps.Name = stock.Name
	ps.Price = stock.Price
	ps.Change = stock.Change
	ps.PercentChange = stock.PercentChange
	ps.Volume = stock.Volume
	ps.AvgVolumeFor_3Month = stock.AvgVolumeFor3Month
	ps.MarketCap = stock.MarketCap
	ps.PeRatio = stock.PERatio
	ps.CreatedAt = stock.CreatedAt.UTC().String()
	ps.CreatedUnix = stock.CreatedUnix

	return &ps
}
