package usastat

import (
	"strings"
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/common"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"github.com/tealeg/xlsx/v3"
)

var (
	INDICATOR_VRP_PATH = "https://www.bea.gov/sites/default/files/2022-06/qgdpstate0622.xlsx"
	VRP_BASE_YEAR      = 2021
)

func GetVRP() (vrp []types.Stat, err error) {
	body, err := common.DownloadFile(INDICATOR_VRP_PATH)
	if err != nil {
		return nil, err
	}

	x, err := xlsx.OpenBinary(body)
	if err != nil {
		return nil, err
	}

	sheet := x.Sheets[3]

	for row := 5; row < 65; row++ {
		locationCell, err := sheet.Cell(row, 0)
		if err != nil {
			continue
		}

		location := locationCell.String()
		location = strings.TrimSpace(location)

		for col := 1; col < 6; col++ {
			valueCell, err := sheet.Cell(row, col)
			if err != nil {
				continue
			}

			value := valueCell.String()
			if value == "" || value == "-" {
				continue
			}

			year := VRP_BASE_YEAR + col/5
			month := ((col-1)%4)*3 + 1

			timestamp := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).String()

			vrp = append(vrp, types.Stat{
				Timestamp: timestamp,
				Location:  location,
				Value:     value,
				Unit:      "Millions of dollars",
				Country:   COUNTRY,
				Section:   "vrp",
			})
		}
	}

	return vrp, nil
}
