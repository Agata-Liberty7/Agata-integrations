package fedstat

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	INDICATOR_CURRENCY_PATH = "https://www.fedstat.ru/indicator/dataGrid.do?id=42108&lineObjectIds=0&lineObjectIds=30611&lineObjectIds=42107&columnObjectIds=3&columnObjectIds=33560&selectedFilterIds=0_42108&selectedFilterIds=3_2002&selectedFilterIds=3_2003&selectedFilterIds=3_2004&selectedFilterIds=3_2005&selectedFilterIds=3_2006&selectedFilterIds=3_2007&selectedFilterIds=3_2008&selectedFilterIds=3_2009&selectedFilterIds=3_2010&selectedFilterIds=3_2011&selectedFilterIds=3_2012&selectedFilterIds=3_2013&selectedFilterIds=3_2014&selectedFilterIds=3_2015&selectedFilterIds=3_2016&selectedFilterIds=3_2017&selectedFilterIds=3_2018&selectedFilterIds=3_2019&selectedFilterIds=3_2020&selectedFilterIds=3_2021&selectedFilterIds=30611_950351&selectedFilterIds=33560_1540220&selectedFilterIds=33560_1540222&selectedFilterIds=33560_1540224&selectedFilterIds=33560_1540226&selectedFilterIds=33560_1540227&selectedFilterIds=33560_1540228&selectedFilterIds=33560_1540229&selectedFilterIds=33560_1540230&selectedFilterIds=33560_1540233&selectedFilterIds=33560_1540234&selectedFilterIds=33560_1540235&selectedFilterIds=33560_1540236&selectedFilterIds=33560_1540272&selectedFilterIds=33560_1540273&selectedFilterIds=33560_1540276&selectedFilterIds=33560_1540282&selectedFilterIds=33560_1540283&selectedFilterIds=33560_1540288&selectedFilterIds=33560_1558883&selectedFilterIds=42107_1613598&selectedFilterIds=42107_1613599"
)

func GetCurrency() (currency []types.Stat, err error) {
	res, err := processFedStatResponse(http.Get(INDICATOR_CURRENCY_PATH))
	if err != nil {
		return nil, err
	}

	for _, r := range res.Results {
		v := r.(map[string]interface{})

		for year := 2002; year <= currentYear; year++ {

			unit := "USD"
			if v["dim42107"] != "Доллар США по отношению к рублю" {
				unit = "EUR"
			}

			for monthNumber, month := range MONTH_CODES {
				for field, val := range v {
					if strings.Contains(field, "dim"+strconv.Itoa(year)+"_"+month) {
						currency = append(currency, types.Stat{
							Timestamp: buildTimestamp(year, monthNumber).String(),
							Name:      unit,
							Unit:      "RUB",
							Value:     val.(string),
							Country:   COUNTRY,
							Section:   "currency",
						})
					}
				}
			}
		}
	}

	return currency, nil
}
