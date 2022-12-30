package kazstat

import (
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/common"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"github.com/tealeg/xlsx/v3"
)

var (
	INDICATOR_VRP_PATH = "https://stat.gov.kz/api/getFile/?docId=ESTAT429111&lang=ru"
	VRP_BASE_YEAR      = 2009
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

	sheet := x.Sheets[4]

	for row := 3; row < 22; row++ {
		locationCell, err := sheet.Cell(row, 0)
		if err != nil {
			continue
		}

		location := locationCell.String()

		if row == 3 {
			location = "Казахстан"
		}

		for col := 1; col < 13; col++ {
			valueCell, err := sheet.Cell(row, col)
			if err != nil {
				continue
			}

			value := valueCell.String()
			if value == "" || value == "-" {
				continue
			}

			year := VRP_BASE_YEAR + (col - 1)

			vrp = append(vrp, types.Stat{
				Timestamp: time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String(),
				Location:  location,
				Value:     value,
				Unit:      "млн. тенге",
				Country:   COUNTRY,
				Section:   "vrp",
			})
		}
	}

	return vrp, nil
}
