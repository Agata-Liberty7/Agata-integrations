package fedstat

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

type FedStatPriceConsumer struct {
	Year  string   `json:"year"`
	Title string   `json:"title"`
	Type  string   `json:"type"`
	Value []string `json:"value"`
}

func GetPriceConsumer() (price []types.Stat, err error) {
	getFedStatRequest(INDICATOR_PRICE_CONSUMER)
	// req, err := getFedStatRequest(INDICATOR_PRICE_CONSUMER)
	// if err != nil {
	// 	return nil, err
	// }

	// res, err := processFedStatResponse(http.DefaultClient.Do(req))
	res, err := processFedStatResponse(http.Get("https://www.fedstat.ru/indicator/dataGrid.do?id=31448&lineObjectIds=0&lineObjectIds=58273&lineObjectIds=57831&lineObjectIds=30611&columnObjectIds=3&columnObjectIds=33560&selectedFilterIds=0_31448&selectedFilterIds=3_2000&selectedFilterIds=3_2001&selectedFilterIds=3_2002&selectedFilterIds=3_2003&selectedFilterIds=3_2004&selectedFilterIds=3_2005&selectedFilterIds=3_2006&selectedFilterIds=3_2007&selectedFilterIds=3_2008&selectedFilterIds=3_2009&selectedFilterIds=3_2010&selectedFilterIds=3_2011&selectedFilterIds=3_2012&selectedFilterIds=3_2013&selectedFilterIds=3_2014&selectedFilterIds=3_2015&selectedFilterIds=3_2016&selectedFilterIds=3_2017&selectedFilterIds=3_2018&selectedFilterIds=3_2019&selectedFilterIds=3_2020&selectedFilterIds=3_2021&selectedFilterIds=30611_950351&selectedFilterIds=33560_1540228&selectedFilterIds=33560_1540229&selectedFilterIds=33560_1540230&selectedFilterIds=33560_1540233&selectedFilterIds=33560_1540234&selectedFilterIds=33560_1540235&selectedFilterIds=33560_1540236&selectedFilterIds=33560_1540272&selectedFilterIds=33560_1540273&selectedFilterIds=33560_1540276&selectedFilterIds=33560_1540282&selectedFilterIds=33560_1540283&selectedFilterIds=57831_1688487&selectedFilterIds=58273_1709713&selectedFilterIds=58273_1709731&selectedFilterIds=58273_1709748&selectedFilterIds=58273_1709753&selectedFilterIds=58273_1709756&selectedFilterIds=58273_1709791&selectedFilterIds=58273_1743458&selectedFilterIds=58273_1743466&selectedFilterIds=58273_1743476&selectedFilterIds=58273_1754708&selectedFilterIds=58273_1754712&selectedFilterIds=58273_1754719&selectedFilterIds=58273_1754735&selectedFilterIds=58273_1754744&selectedFilterIds=58273_1754759&selectedFilterIds=58273_1755075&selectedFilterIds=58273_1755103&selectedFilterIds=58273_1755177"))

	if err != nil {
		return nil, err
	}

	for _, r := range res.Results {
		v := r.(map[string]interface{})

		for year := 2000; year <= currentYear; year++ {
			for monthNumber, month := range MONTH_CODES {
				for field, val := range v {
					if strings.Contains(field, "dim"+strconv.Itoa(year)+"_"+month) {
						price = append(price, types.Stat{
							Timestamp: buildTimestamp(year, monthNumber).String(),
							Location:  v["dim57831"].(string),
							Name:      v["dim58273"].(string),
							Value:     val.(string),
							Unit:      "RUB",
							Country:   COUNTRY,
							Section:   "priceConsumer",
						})
						continue
					}
				}
			}
		}
	}

	return price, nil
}
