package kazstat

import (
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/common"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"github.com/tealeg/xlsx/v3"
)

var (
	INDICATOR_POPULATION_PATH = "https://stat.gov.kz/api/getFile/?docId=ESTAT334809&lang=ru"
	POPULATION_BASE_YEAR      = 2021
)

func GetPopulation() (population []types.Stat, err error) {
	body, err := common.DownloadFile(INDICATOR_POPULATION_PATH)
	if err != nil {
		return nil, err
	}

	x, err := xlsx.OpenBinary(body)
	if err != nil {
		return nil, err
	}

	sheet := x.Sheets[0]

	for row := 10; row < 27; row++ {
		locationCell, err := sheet.Cell(row, 0)
		if err != nil {
			continue
		}

		location := locationCell.String()

		if row == 10 {
			location = "Казахстан"
		}

		for col := 1; col < 24; col++ {
			valueCell, err := sheet.Cell(row, col)
			if err != nil {
				continue
			}

			value := valueCell.String()
			if value == "" || value == "-" {
				continue
			}

			year := POPULATION_BASE_YEAR + (col-1)/12
			month := (col-1)%12 + 1
			timestamp := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC).String()

			population = append(population, types.Stat{
				Timestamp: timestamp,
				Location:  location,
				Value:     value,
				Country:   COUNTRY,
				Section:   "population",
			})
		}
	}

	return population, nil
}
