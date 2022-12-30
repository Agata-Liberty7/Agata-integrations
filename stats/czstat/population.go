package czstat

import (
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/common"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"github.com/tealeg/xlsx/v3"
)

var (
	INDICATOR_POPULATION_PATH = "https://www.czso.cz/documents/10180/142756992/1300682100.xlsx/7d338880-9c02-4dcd-9b5a-04a6b9b7c2bf?version=1.3"
	POPULATION_BASE_YEAR      = 2011
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

	location := "Česká republika"

	for col := 2; col <= 11; col++ {
		valueCell, err := sheet.Cell(2, col)
		if err != nil {
			continue
		}

		value := valueCell.String()
		if value == "" || value == "-" {
			continue
		}

		year := POPULATION_BASE_YEAR + (col - 2)
		timestamp := time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String()

		population = append(population, types.Stat{
			Timestamp: timestamp,
			Location:  location,
			Value:     value,
			Country:   COUNTRY,
			Section:   "population",
		})
	}

	return population, nil
}
