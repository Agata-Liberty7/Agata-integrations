package usastat

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	INDICATOR_POPULATION_PATH = "https://datausa.io/api/data?drilldowns=State&measures=Population"
)

func GetPopulation() (population []types.Stat, err error) {
	res, err := processUsaStatResponse(http.Get(INDICATOR_POPULATION_PATH))
	if err != nil {
		return nil, err
	}

	for _, item := range res.Data {
		timestamp := time.Date(item.Year, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String()

		population = append(population, types.Stat{
			Timestamp: timestamp,
			Location:  item.State,
			Value:     strconv.Itoa(item.Population),
			Country:   COUNTRY,
			Section:   "population",
		})
	}
	return
}
