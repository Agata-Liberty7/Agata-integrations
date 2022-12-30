package belstat

import (
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	INDICATOR_VRP_PATH = "https://ec.europa.eu/eurostat/databrowser-backend/api/extraction/1.0/LIVE/false/json/en/TGS00003?cacheId=1643320800000-2.5.10%2520-%25202022-01-25%252010%253A43"
	VRP_BASE_YEAR      = 2008
	VRP_YEARS          = 2021 - VRP_BASE_YEAR
)

func GetVRP() (vrp []types.Stat, err error) {
	// res, err := processEuroStatResponse(http.Get(INDICATOR_VRP_PATH))
	// if err != nil {
	// 	return nil, err
	// }

	// locations := make([]string, len(res.Dimension.Geo.Category.Index))
	// for location, index := range res.Dimension.Geo.Category.Index {
	// 	locations[index] = res.Dimension.Geo.Category.Label[location]
	// }

	// for i, r := range res.Value {
	// 	index, _ := strconv.Atoi(i)
	// 	vrp = append(vrp, types.Stat{
	// 		Timestamp: time.Date(index%VRP_YEARS+VRP_BASE_YEAR, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String(),
	// 		Location:  locations[index/VRP_YEARS],
	// 		Value:     r,
	// 		Unit:      "million EUR",
	// 	})
	// }

	return vrp, nil
}
