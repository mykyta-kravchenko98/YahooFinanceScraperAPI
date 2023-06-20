package yahoo

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
	"yahooscraper/internal/models"

	"github.com/gocolly/colly"
	"github.com/rs/zerolog/log"
)

var (
	HeaderLengthNotEqualToLineDataLength error = errors.New("Func:convertToStocks. Headers length not equal line data length")
	UnknowErrorDuringDataParsing         error = errors.New("Unknow error during data parsing")
)

// Starting the process of parsing and returning parsed data
func Start() ([]models.Stock, error) {
	headers, stocksCount := parseHeaders()
	data := parseData(stocksCount)

	if len(data) == 0 {
		return nil, UnknowErrorDuringDataParsing
	}

	result := convertToStocks(data, headers)

	return result, nil
}

// Parse header
func parseHeaders() (headers []string, stocksCount int64) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"),
		colly.AllowedDomains("finance.yahoo.com"),
		colly.MaxBodySize(0),
		colly.AllowURLRevisit(),
		colly.Async(true),
	)

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       500 * time.Millisecond,
	})
	c.OnRequest(func(r *colly.Request) {
		log.Debug().Msg(fmt.Sprintf("Visiting to %s", r.URL.String()))
	})
	c.OnHTML("thead", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			el.ForEach("th", func(_ int, el *colly.HTMLElement) {
				//ignore 52 Week Range
				if el.Text != "52 Week Range" {
					headers = append(headers, el.Text)
				}
			})
		})
	})

	c.Visit("https://finance.yahoo.com/most-active")

	c.OnHTML("span", func(e *colly.HTMLElement) {
		if strings.Contains(e.Text, "result") {
			objects := strings.Split(e.Text, " ")
			//extract number of stocks in most-active sections
			stocksCount, _ = strconv.ParseInt(objects[2], 10, 64)
		}
	})

	c.Wait()

	return headers, stocksCount
}

// Parse body
func parseData(stocksCount int64) (data [][]string) {
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 11_2_1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/88.0.4324.182 Safari/537.36"),
		colly.AllowedDomains("finance.yahoo.com"),
		colly.MaxBodySize(0),
		colly.AllowURLRevisit(),
		colly.Async(true),
	)

	// Set max Parallelism and introduce a Random Delay
	c.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 2,
		Delay:       500 * time.Millisecond,
	})
	c.OnRequest(func(r *colly.Request) {
		log.Debug().Msg(fmt.Sprintf("Visiting to %s", r.URL.String()))
	})

	c.OnHTML("tbody", func(e *colly.HTMLElement) {
		e.ForEach("tr", func(_ int, el *colly.HTMLElement) {
			dataSlice := []string{}
			el.ForEach("td", func(_ int, el *colly.HTMLElement) {
				dataSlice = append(dataSlice, el.Text)
			})
			//remove the last element which is the empty string
			dataSlice = dataSlice[:len(dataSlice)-1]

			//add to overall slice
			data = append(data, dataSlice)
		})
	})

	stocksCountInt := int(stocksCount)
	for i := 0; i <= stocksCountInt; i += 100 {
		count := 100
		capacity := stocksCountInt - i
		if capacity < 100 {
			count = capacity
		}
		c.Visit(fmt.Sprintf("https://finance.yahoo.com/most-active?offset=%d&count=%d", i, count))
	}

	c.Wait()

	return data
}

func convertToStocks(data [][]string, headers []string) (result []models.Stock) {
	lineCount := 0

	for _, line := range data {
		lineCount++
		if len(line) != len(headers) {
			log.Error().Err(HeaderLengthNotEqualToLineDataLength).Msg(fmt.Sprintf("Exception: func:convertToStocks, data line:%d. Header count not equal line data count.", lineCount))
			continue
		}

		stock := models.Stock{}
		var convertError error

		stock.Symbol = line[0]
		stock.Name = line[1]

		stock.Price, convertError = strconv.ParseFloat(line[2], 64)

		if convertError != nil {
			log.Error().Err(convertError).Msg(fmt.Sprintf("Exception: func:convertToStocks, data line:%d. Price data '%s' can`t be convert from string to float64.", lineCount, line[2]))
			continue
		}

		stock.Change, convertError = strconv.ParseFloat(line[3], 64)

		if convertError != nil {
			log.Error().Err(convertError).Msg(fmt.Sprintf("Exception: func:convertToStocks, data line:%d. Change data '%s' can`t be convert from string to float64.", lineCount, line[3]))
			continue
		}

		if percentChangeWithoutSign, ok := strings.CutSuffix(line[4], "%"); ok {
			stock.PercentChange, convertError = strconv.ParseFloat(percentChangeWithoutSign, 64)

			if convertError != nil {
				log.Error().Err(convertError).Msg(fmt.Sprintf("Exception: func:convertToStocks, data line:%d. Change data '%s' can`t be convert from string to float64.", lineCount, line[4]))
				continue
			}
		} else {
			log.Error().Err(convertError).Msg("Exception during extracting PercentChange without sign.")
			continue
		}

		stock.Volume = line[5]
		stock.AvgVolumeFor3Month = line[6]
		stock.MarketCap = line[7]
		stock.PERatio = line[8]

		result = append(result, stock)
	}

	return result
}
