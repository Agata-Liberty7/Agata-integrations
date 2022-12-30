package destat

import (
	"log"
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/common"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"github.com/tealeg/xlsx/v3"
)

var (
	INDICATOR_VRP_PATH = "https://www-genesis.destatis.de/genesis/online?operation=ergebnistabelleDownload&levelindex=2&levelid=1657118366442&option=xlsx"
	VRP_BASE_YEAR      = 2021
)

func GetVRP() (vrp []types.Stat, err error) {
	body, err := common.DownloadFile(INDICATOR_POPULATION_PATH)
	if err != nil {
		log.Println("Faild to parse, fallback to mock:", err)
		return VrpData, nil
	}

	x, err := xlsx.OpenBinary(body)
	if err != nil {
		log.Println("Faild to parse, fallback to mock:", err)
		return VrpData, nil
	}

	sheet := x.Sheets[0]

	for row := 5; row < 71; row++ {
		location := "Deutschland"

		valueCell, err := sheet.Cell(row, 1)
		if err != nil {
			continue
		}

		value := valueCell.String()
		if value == "" || value == "-" {
			continue
		}

		year := POPULATION_BASE_YEAR + (row - 5)
		timestamp := time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String()

		vrp = append(vrp, types.Stat{
			Timestamp: timestamp,
			Location:  location,
			Value:     value,
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		})
	}

	return vrp, nil
}
