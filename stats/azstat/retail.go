package azstat

import (
	"bytes"
	"strings"
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/common"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"github.com/extrame/xls"
)

var (
	INDICATOR_RETAIL_PATH = "https://www.stat.gov.az/source/trade/en/f_trade/xt008_1en.xls"
	RETAIL_BASE_YEAR      = 1994
)

func GetRetail() (vrp []types.Stat, err error) {
	body, err := common.DownloadFile(INDICATOR_RETAIL_PATH)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(body)
	wb, err := xls.OpenReader(reader, "utf-8")
	if err != nil {
		return nil, err
	}

	sheet := wb.GetSheet(0)

	for row := 4; row < 119; row++ {
		name := strings.TrimSpace(sheet.Row(row).Col(0))
		for col := 3; col <= 29; col++ {

			value := sheet.Row(row).Col(col)
			if value == "" || value == "-" || strings.Contains(value, "Z") {
				continue
			}

			year := RETAIL_BASE_YEAR + (col - 3)

			vrp = append(vrp, types.Stat{
				Timestamp: time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String(),
				// Location:  "",
				Value:   value,
				Name:    name,
				Unit:    "thousand US dollars",
				Country: COUNTRY,
				Section: "retail",
			})
		}
	}

	return vrp, nil
}
