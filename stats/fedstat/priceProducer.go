package fedstat

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	INDICATOR_PRICE_PRODUCER_PATH = "https://www.fedstat.ru/indicator/dataGrid.do?id=57606&lineObjectIds=0&lineObjectIds=30611&lineObjectIds=58601&lineObjectIds=58601&lineObjectIds=57831&lineObjectIds=57939&lineObjectIds=57956&columnObjectIds=3&columnObjectIds=33560&selectedFilterIds=0_57606&selectedFilterIds=3_2017&selectedFilterIds=3_2018&selectedFilterIds=3_2019&selectedFilterIds=3_2020&selectedFilterIds=3_2021&selectedFilterIds=30611_950351&selectedFilterIds=33560_1540228&selectedFilterIds=33560_1540229&selectedFilterIds=33560_1540230&selectedFilterIds=33560_1540233&selectedFilterIds=33560_1540234&selectedFilterIds=33560_1540235&selectedFilterIds=33560_1540236&selectedFilterIds=33560_1540272&selectedFilterIds=33560_1540273&selectedFilterIds=33560_1540276&selectedFilterIds=33560_1540282&selectedFilterIds=33560_1540283&selectedFilterIds=57831_1688487&selectedFilterIds=57939_1692983&selectedFilterIds=57939_1693006&selectedFilterIds=57939_1693320&selectedFilterIds=57939_1693328&selectedFilterIds=57939_1693344&selectedFilterIds=57939_1693348&selectedFilterIds=57939_1693382&selectedFilterIds=57939_1693657&selectedFilterIds=57939_1693664&selectedFilterIds=57939_1693668&selectedFilterIds=57939_1693688&selectedFilterIds=57939_1693717&selectedFilterIds=57939_1693828&selectedFilterIds=57939_1693843&selectedFilterIds=57939_1693944&selectedFilterIds=57939_1694031&selectedFilterIds=57939_1694032&selectedFilterIds=57939_1694073&selectedFilterIds=57939_1694076&selectedFilterIds=57939_1694082&selectedFilterIds=57939_1694087&selectedFilterIds=57939_1694135&selectedFilterIds=57939_1694306&selectedFilterIds=57939_1694312&selectedFilterIds=57939_1694797&selectedFilterIds=57939_1757622&selectedFilterIds=57939_1757637&selectedFilterIds=57939_1757645&selectedFilterIds=57939_1757665&selectedFilterIds=57939_1757761&selectedFilterIds=57939_1757885&selectedFilterIds=57939_1757888&selectedFilterIds=57939_1757901&selectedFilterIds=57939_1757908&selectedFilterIds=57939_1757914&selectedFilterIds=57956_1694823&selectedFilterIds=57956_1694824&selectedFilterIds=57956_1694825&selectedFilterIds=57956_1694826&selectedFilterIds=57956_1694828&selectedFilterIds=57956_1694829&selectedFilterIds=57956_1694831&selectedFilterIds=57956_1694832&selectedFilterIds=57956_1694834&selectedFilterIds=57956_1694835&selectedFilterIds=57956_1694836&selectedFilterIds=57956_1694837&selectedFilterIds=57956_1694838&selectedFilterIds=57956_1694841&selectedFilterIds=57956_1694842&selectedFilterIds=57956_1694846&selectedFilterIds=57956_1694847&selectedFilterIds=57956_1694851&selectedFilterIds=57956_1694857&selectedFilterIds=57956_1694860&selectedFilterIds=57956_1694865&selectedFilterIds=57956_1694866&selectedFilterIds=57956_1704138&selectedFilterIds=57956_1744082&selectedFilterIds=57956_1744083&selectedFilterIds=57956_1744084&selectedFilterIds=57956_1744085&selectedFilterIds=57956_1750461&selectedFilterIds=57956_1752495&selectedFilterIds=57956_1757950&selectedFilterIds=57956_1757951&selectedFilterIds=57956_1757952&selectedFilterIds=57956_1757953&selectedFilterIds=57956_1757954&selectedFilterIds=57956_1757955&selectedFilterIds=57956_1757956&selectedFilterIds=57956_1757957&selectedFilterIds=58601_1744136"
)

func GetPriceProducer() (price []types.Stat, err error) {
	res, err := processFedStatResponse(http.Get(INDICATOR_PRICE_PRODUCER_PATH))
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
							Section:   "priceProducer",
						})
						continue
					}
				}
			}
		}
	}

	return price, nil
}
