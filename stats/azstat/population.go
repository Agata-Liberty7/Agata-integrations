package azstat

import (
	"bytes"
	"strconv"
	"strings"
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/common"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"github.com/extrame/xls"
)

var (
	INDICATOR_POPULATION_PATH = "https://www.stat.gov.az/source/demoqraphy/en/001_15en.xls"
	POPULATION_BASE_YEAR      = 2009
)

func GetPopulation() (population []types.Stat, err error) {
	body, err := common.DownloadFile(INDICATOR_POPULATION_PATH)
	if err != nil {
		return nil, err
	}

	reader := bytes.NewReader(body)
	wb, err := xls.OpenReader(reader, "utf-8")
	if err != nil {
		return nil, err
	}

	sheet := wb.GetSheet(0)

	for row := 4; row < 135; row++ {
		location := sheet.Row(row).Col(1)
		if location == "  including:" || location == "" {
			continue
		}

		location = strings.TrimSpace(location)

		for col := 3; col <= 4; col++ {

			// TODO: some values returns from xls formated as date
			// witch fails to parse to int
			value := sheet.Row(row).Col(col)
			intValue, _ := strconv.ParseFloat(value, 64)
			value = strconv.Itoa(int(intValue * 1000))

			if value == "0" || value == "-" {
				continue
			}

			year := POPULATION_BASE_YEAR
			if col == 4 {
				year = 2021
			}

			population = append(population, types.Stat{
				Timestamp: time.Date(year, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String(),
				Location:  location,
				Value:     value,
				Country:   COUNTRY,
				Section:   "population",
			})
		}
	}

	return population, nil
}
