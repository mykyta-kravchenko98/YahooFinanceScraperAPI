package models

import (
	"time"

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
