package fedstat

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	INDICATOR_PRICE_AVERAGE_PATH = "https://www.fedstat.ru/indicator/dataGrid.do?id=57786&lineObjectIds=0&lineObjectIds=30611&lineObjectIds=57831&lineObjectIds=57956&lineObjectIds=57939&columnObjectIds=3&columnObjectIds=33560&selectedFilterIds=0_57786&selectedFilterIds=3_2021&selectedFilterIds=30611_1341910&selectedFilterIds=33560_1540282&selectedFilterIds=33560_1540283&selectedFilterIds=57831_1688487&selectedFilterIds=57939_1692961&selectedFilterIds=57939_1692963&selectedFilterIds=57939_1692990&selectedFilterIds=57939_1693000&selectedFilterIds=57939_1693164&selectedFilterIds=57939_1693165&selectedFilterIds=57939_1693247&selectedFilterIds=57939_1693269&selectedFilterIds=57939_1693295&selectedFilterIds=57939_1693296&selectedFilterIds=57939_1693297&selectedFilterIds=57939_1693331&selectedFilterIds=57939_1693332&selectedFilterIds=57939_1693354&selectedFilterIds=57939_1693355&selectedFilterIds=57939_1693371&selectedFilterIds=57939_1693373&selectedFilterIds=57939_1693377&selectedFilterIds=57939_1693389&selectedFilterIds=57939_1693390&selectedFilterIds=57939_1693407&selectedFilterIds=57939_1693408&selectedFilterIds=57939_1693690&selectedFilterIds=57939_1693716&selectedFilterIds=57939_1693717&selectedFilterIds=57939_1693733&selectedFilterIds=57939_1693740&selectedFilterIds=57939_1693787&selectedFilterIds=57939_1693798&selectedFilterIds=57939_1693833&selectedFilterIds=57939_1693834&selectedFilterIds=57939_1693835&selectedFilterIds=57939_1694690&selectedFilterIds=57939_1694922&selectedFilterIds=57939_1699115&selectedFilterIds=57939_1699129&selectedFilterIds=57939_1699130&selectedFilterIds=57939_1699132&selectedFilterIds=57939_1700494&selectedFilterIds=57939_1788712&selectedFilterIds=57956_1694823&selectedFilterIds=57956_1694824&selectedFilterIds=57956_1694825&selectedFilterIds=57956_1694826&selectedFilterIds=57956_1694827&selectedFilterIds=57956_1694828&selectedFilterIds=57956_1694829&selectedFilterIds=57956_1694830&selectedFilterIds=57956_1694831&selectedFilterIds=57956_1694832&selectedFilterIds=57956_1694833&selectedFilterIds=57956_1694834&selectedFilterIds=57956_1694835&selectedFilterIds=57956_1694836&selectedFilterIds=57956_1694837&selectedFilterIds=57956_1694838&selectedFilterIds=57956_1694839&selectedFilterIds=57956_1694840&selectedFilterIds=57956_1694841&selectedFilterIds=57956_1694842&selectedFilterIds=57956_1694843&selectedFilterIds=57956_1694844&selectedFilterIds=57956_1694845&selectedFilterIds=57956_1694846&selectedFilterIds=57956_1694847&selectedFilterIds=57956_1694848&selectedFilterIds=57956_1694849&selectedFilterIds=57956_1694850&selectedFilterIds=57956_1694851&selectedFilterIds=57956_1694852&selectedFilterIds=57956_1694853&selectedFilterIds=57956_1694854&selectedFilterIds=57956_1694855&selectedFilterIds=57956_1694856&selectedFilterIds=57956_1694857&selectedFilterIds=57956_1694858&selectedFilterIds=57956_1694859&selectedFilterIds=57956_1694860&selectedFilterIds=57956_1694861&selectedFilterIds=57956_1694862&selectedFilterIds=57956_1694863&selectedFilterIds=57956_1694864&selectedFilterIds=57956_1694865&selectedFilterIds=57956_1694866&selectedFilterIds=57956_1694867&selectedFilterIds=57956_1694868&selectedFilterIds=57956_1694869&selectedFilterIds=57956_1694870&selectedFilterIds=57956_1694871"
)

func GetPriceAverage() (price []types.Stat, err error) {
	res, err := processFedStatResponse(http.Get(INDICATOR_PRICE_AVERAGE_PATH))
	if err != nil {
		return nil, err
	}

	for _, r := range res.Results {
		v := r.(map[string]interface{})

		for year := 2017; year <= currentYear; year++ {
			for monthNumber, month := range MONTH_CODES {
				for field, val := range v {
					if strings.Contains(field, "dim"+strconv.Itoa(year)+"_"+month) {
						price = append(price, types.Stat{
							Timestamp: buildTimestamp(year, monthNumber).String(),
							Name:      v["dim57939"].(string),
							Unit:      v["dim57956"].(string),
							Value:     val.(string),
							Country:   COUNTRY,
							Section:   "priceAverage",
						})
						continue
					}
				}
			}
		}
	}

	return price, nil
}
