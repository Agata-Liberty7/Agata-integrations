package fedstat

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	INDICATOR_VVP_PATH = "https://www.fedstat.ru/indicator/dataGrid.do?id=59448&lineObjectIds=0&lineObjectIds=30611&lineObjectIds=33560&lineObjectIds=57831&lineObjectIds=57831&lineObjectIds=57940&columnObjectIds=3&selectedFilterIds=0_59448&selectedFilterIds=3_2016&selectedFilterIds=3_2017&selectedFilterIds=3_2018&selectedFilterIds=3_2019&selectedFilterIds=30611_950352&selectedFilterIds=33560_1558883&selectedFilterIds=57831_1688487&selectedFilterIds=57940_1692320&selectedFilterIds=57940_1692321&selectedFilterIds=57940_1692322&selectedFilterIds=57940_1692323&selectedFilterIds=57940_1692324&selectedFilterIds=57940_1692325&selectedFilterIds=57940_1692326&selectedFilterIds=57940_1692327&selectedFilterIds=57940_1692328&selectedFilterIds=57940_1692329&selectedFilterIds=57940_1692330&selectedFilterIds=57940_1692331&selectedFilterIds=57940_1692332&selectedFilterIds=57940_1692748&selectedFilterIds=57940_1692749&selectedFilterIds=57940_1692750&selectedFilterIds=57940_1692751&selectedFilterIds=57940_1692933&selectedFilterIds=57940_1695397&selectedFilterIds=57940_1695572"
)

func GetVVP() (vvp []types.Stat, err error) {
	res, err := processFedStatResponse(http.Get(INDICATOR_VVP_PATH))
	if err != nil {
		return nil, err
	}

	for _, r := range res.Results {
		v := r.(map[string]interface{})

		for field, val := range v {
			for year := 2000; year <= currentYear; year++ {
				if strings.Contains(field, "dim"+strconv.Itoa(year)) {
					vvp = append(vvp, types.Stat{
						Timestamp: buildTimestamp(year, 1).String(),
						Name:      v["dim57940"].(string),
						Value:     val.(string),
						Country:   COUNTRY,
						Section:   "vvp",
					})
				}
			}
		}
	}

	return vvp, nil
}
