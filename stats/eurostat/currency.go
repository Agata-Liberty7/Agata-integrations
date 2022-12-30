package eurostat

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	INDICATOR_CURRENCY_PATH = "https://ec.europa.eu/eurostat/databrowser-backend/api/extraction/1.0/LIVE/false/json/en/TEC00033?cacheId=1643796000000-2.5.10%2520-%25202022-01-25%252010%253A43"
	CURRENCY_BASE_YEAR      = 2010
	CURRENCY_YEARS          = 2021 - CURRENCY_BASE_YEAR
)

func GetCurrency() (currency []types.Stat, err error) {
	res, err := processEuroStatResponse(http.Get(INDICATOR_CURRENCY_PATH))
	if err != nil {
		return nil, err
	}

	currencies := make([]string, len(res.Dimension.Currency.Category.Index))
	for location, index := range res.Dimension.Currency.Category.Index {
		currencies[index] = res.Dimension.Currency.Category.Label[location]
	}

	for i, r := range res.Value {
		index, _ := strconv.Atoi(i)
		currency = append(currency, types.Stat{
			Timestamp: time.Date(index%VRP_YEARS+VRP_BASE_YEAR, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String(),
			Name:      currencies[index/VRP_YEARS],
			Value:     r,
			Unit:      "EUR",
			Country:   COUNTRY,
			Section:   "currency",
		})
	}

	return currency, nil
}
