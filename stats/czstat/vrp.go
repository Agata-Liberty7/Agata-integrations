package czstat

import (
	"strings"
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/common"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"github.com/tealeg/xlsx/v3"
)

var (
	INDICATOR_VRP_PATH = "https://apl.czso.cz/pll/rocenka/rocenka.presB?jmeno_tabulka=RD01-1&rok=2020&mylang=EN&ceny=bc&vystup=EXCEL&priznak=RD&typ=1&jak=4&dejarchiv=0"
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

	sheet := x.Sheets[0]

	for row := 6; row < 31; row++ {
		locationCell, err := sheet.Cell(row, 0)
		if err != nil {
			continue
		}

		location := locationCell.String()

		if strings.Contains(location, "NUTS") {
			continue
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

			nameCell, err := sheet.Cell(3, col)
			if err != nil {
				continue
			}
			name := nameCell.String()

			vrp = append(vrp, types.Stat{
				Timestamp: time.Date(VRP_BASE_YEAR, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String(),
				Location:  location,
				Value:     value,
				Name:      name,
				Unit:      "million CZK",
				Country:   COUNTRY,
				Section:   "vrp",
			})
		}
	}

	return vrp, nil
}
