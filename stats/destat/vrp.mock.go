package destat

import (
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	timestamp = time.Date(2021, time.Month(1), 1, 0, 0, 0, 0, time.UTC).String()

	VrpData = []types.Stat{
		{
			Timestamp: timestamp,
			Location:  "Deutschland",
			Value:     "83237124",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Baden-Württemberg",
			Value:     "11124642",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Bayern",
			Value:     "13176989",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Berlin",
			Value:     "3677472",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Brandenburg",
			Value:     "2537868",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Bremen",
			Value:     "2537868",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Hamburg",
			Value:     "1853935",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Hessen",
			Value:     "6295017",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Mecklenburg-Vorpommern",
			Value:     "1611160",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Niedersachsen",
			Value:     "8027031",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Nordrhein-Westfalen",
			Value:     "17924591",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Rheinland-Pfalz",
			Value:     "4106485",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Saarland",
			Value:     "982348",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Sachsen",
			Value:     "4043002",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Sachsen-Anhalt",
			Value:     "2169253",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Schleswig-Holstein",
			Value:     "2922005",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Thüringen",
			Value:     "2108863",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
		{
			Timestamp: timestamp,
			Location:  "Sachsen",
			Value:     "4043002",
			Unit:      "million EUR",
			Country:   COUNTRY,
			Section:   "vrp",
		},
	}
)
