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
	INDICATOR_VRP_PATH = "https://www.stat.gov.az/source/system_nat_accounts/en/034en.xls"
	VRP_BASE_YEAR      = 2003
)

func GetVRP() (vrp []types.Stat, err error) {
	body, err := common.DownloadFile(INDICATOR_VRP_PATH)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(body)
	wb, err := xls.OpenReader(reader, "UTF-8")
	if err != nil {
		return nil, err
	}

	sheet := wb.GetSheet(0)

	for row := 4; row < 104; row++ {
		location := sheet.Row(row).Col(1)
		if location == "  including:" {
			continue
		}

		location = strings.TrimSpace(location)

		for col := 2; col < 24; col++ {
			value := sheet.Row(row).Col(col)
			if value == "" || value == "-" {
				continue
			}

			year := VRP_BASE_YEAR + (col - 2)

			vrp = append(vrp, types.Stat{
				Timestamp: time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String(),
				Location:  location,
				Value:     value,
				Unit:      "thsd. manats",
				Country:   COUNTRY,
				Section:   "vrp",
			})
		}
	}

	return vrp, nil
}
