package kazstat

import (
	"bytes"
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/common"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"github.com/extrame/xls"
)

var (
	INDICATOR_RETAIL_PATH = "https://stat.gov.kz/api/getFile/?docId=ESTAT101167&lang=ru"
	RETAIL_BASE_YEAR      = 2009
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

	bases := []int{5, 28, 50, 72, 94, 117, 140, 163, 187, 211, 234}
	// bases := []int{5, 28, 50, 72, 94, 117, 140, 163, 187, 211, 234, 260, 284}

	for baseNum, base := range bases {
		for row := base; row < base+19; row++ {
			location := sheet.Row(row).Col(0)

			value := sheet.Row(row).Col(1)
			if value == "" || value == "-" {
				continue
			}

			year := RETAIL_BASE_YEAR + baseNum

			vrp = append(vrp, types.Stat{
				Timestamp: time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String(),
				Location:  location,
				Name:      "Сеть магазинов розничной торговли по продаже потребительских товаров по областям",
				Value:     value,
				Country:   COUNTRY,
				Section:   "retail",
			})
		}
	}

	return vrp, nil
}
