package eurostat

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	INDICATOR_Population_PATH = "https://ec.europa.eu/eurostat/databrowser-backend/api/extraction/1.0/LIVE/false/json/en/DEMO_R_GIND3__custom_2044874?cacheId=1643320800000-2.5.10%2520-%25202022-01-25%252010%253A43"
	POPULATION_BASE_YEAR      = 2000
	POPULATION_YEARS          = 2021 - POPULATION_BASE_YEAR
)

func GetPopulation() (population []types.Stat, err error) {
	res, err := processEuroStatResponse(http.Get(INDICATOR_VRP_PATH))
	if err != nil {
		return nil, err
	}

	locations := make([]string, len(res.Dimension.Geo.Category.Index))
	for location, index := range res.Dimension.Geo.Category.Index {
		locations[index] = res.Dimension.Geo.Category.Label[location]
	}

	for i, r := range res.Value {
		index, _ := strconv.Atoi(i)
		population = append(population, types.Stat{
			Timestamp: time.Date(index%POPULATION_YEARS+POPULATION_BASE_YEAR, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String(),
			Location:  locations[index/POPULATION_YEARS],
			Name:      strconv.Itoa(len(res.Value)),
			Value:     r,
			Country:   COUNTRY,
			Section:   "population",
		})
	}

	return population, nil
}
