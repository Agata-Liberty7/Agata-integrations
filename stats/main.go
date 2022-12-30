package stats

import (
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/stats/azstat"
	"github.com/Agata-Liberty7/Agata-integrations/stats/belstat"
	"github.com/Agata-Liberty7/Agata-integrations/stats/czstat"
	"github.com/Agata-Liberty7/Agata-integrations/stats/destat"
	"github.com/Agata-Liberty7/Agata-integrations/stats/eurostat"
	"github.com/Agata-Liberty7/Agata-integrations/stats/fedstat"
	"github.com/Agata-Liberty7/Agata-integrations/stats/kazstat"
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
	"github.com/Agata-Liberty7/Agata-integrations/stats/usastat"
)

// type Parser

// type StatResponse struct {
// 	Timestamp string `json:"timestamp,omitempty"`
// 	Location  string `json:"location,omitempty"`
// 	Name      string `json:"name,omitempty"`
// 	Value     string `json:"value,omitempty"`
// 	Unit      string `json:"unit,omitempty"`
// }
// type StatResponses []StatResponse

// type Stat struct {
// 	gorm.Model
// 	Timestamp string `json:"timestamp"`
// 	Location  string `json:"location"`
// 	Name      string `json:"name"`
// 	Value     string `json:"value"`
// 	Unit      string `json:"unit"`
// 	Country   string `json:"country" gorm:"not null"`
// 	Section   string `json:"section" gorm:"not null"`
// }
// type Stats []Stat

// type Metadata struct {
// 	Count uint32 `json:"count"`
// 	Time  string `json:"time"`
// }

// type MetadataRaw struct {
// 	Section string `json:"section"`
// 	Count   uint32 `json:"count"`
// 	Time    string `json:"time"`
// }

// type SettingsItem struct {
// 	Timestamp string `json:"timestamp"`
// 	Location  string `json:"location"`
// 	Section   string `json:"section"`
// 	Period    string `json:"period"`
// }
// type SettingsItems []SettingsItem

var (
	startTime = time.Now().Format("02.01 15:04")
	MetaDB    = types.SettingsItems{
		types.SettingsItem{Location: "RU", Timestamp: startTime, Period: "month", Section: "vvp"},
		types.SettingsItem{Location: "RU", Timestamp: startTime, Period: "month", Section: "vrp"},
		types.SettingsItem{Location: "RU", Timestamp: startTime, Period: "month", Section: "currency"},
		types.SettingsItem{Location: "RU", Timestamp: startTime, Period: "month", Section: "population"},
		types.SettingsItem{Location: "RU", Timestamp: startTime, Period: "month", Section: "retail"},
		types.SettingsItem{Location: "RU", Timestamp: startTime, Period: "month", Section: "priceProducer"},
		types.SettingsItem{Location: "RU", Timestamp: startTime, Period: "month", Section: "priceConsumer"},
		types.SettingsItem{Location: "RU", Timestamp: startTime, Period: "month", Section: "priceAverage"},

		types.SettingsItem{Location: "KZ", Timestamp: startTime, Period: "month", Section: "vvp"},
		types.SettingsItem{Location: "KZ", Timestamp: startTime, Period: "month", Section: "vrp"},
		types.SettingsItem{Location: "KZ", Timestamp: startTime, Period: "month", Section: "currency"},
		types.SettingsItem{Location: "KZ", Timestamp: startTime, Period: "month", Section: "population"},
		types.SettingsItem{Location: "KZ", Timestamp: startTime, Period: "month", Section: "retail"},
		types.SettingsItem{Location: "KZ", Timestamp: startTime, Period: "month", Section: "priceProducer"},
		types.SettingsItem{Location: "KZ", Timestamp: startTime, Period: "month", Section: "priceConsumer"},
		types.SettingsItem{Location: "KZ", Timestamp: startTime, Period: "month", Section: "priceAverage"},

		types.SettingsItem{Location: "BY", Timestamp: startTime, Period: "month", Section: "vvp"},
		types.SettingsItem{Location: "BY", Timestamp: startTime, Period: "month", Section: "vrp"},
		types.SettingsItem{Location: "BY", Timestamp: startTime, Period: "month", Section: "currency"},
		types.SettingsItem{Location: "BY", Timestamp: startTime, Period: "month", Section: "population"},
		types.SettingsItem{Location: "BY", Timestamp: startTime, Period: "month", Section: "retail"},
		types.SettingsItem{Location: "BY", Timestamp: startTime, Period: "month", Section: "priceProducer"},
		types.SettingsItem{Location: "BY", Timestamp: startTime, Period: "month", Section: "priceConsumer"},
		types.SettingsItem{Location: "BY", Timestamp: startTime, Period: "month", Section: "priceAverage"},

		types.SettingsItem{Location: "UA", Timestamp: startTime, Period: "month", Section: "vvp"},
		types.SettingsItem{Location: "UA", Timestamp: startTime, Period: "month", Section: "vrp"},
		types.SettingsItem{Location: "UA", Timestamp: startTime, Period: "month", Section: "currency"},
		types.SettingsItem{Location: "UA", Timestamp: startTime, Period: "month", Section: "population"},
		types.SettingsItem{Location: "UA", Timestamp: startTime, Period: "month", Section: "retail"},
		types.SettingsItem{Location: "UA", Timestamp: startTime, Period: "month", Section: "priceProducer"},
		types.SettingsItem{Location: "UA", Timestamp: startTime, Period: "month", Section: "priceConsumer"},
		types.SettingsItem{Location: "UA", Timestamp: startTime, Period: "month", Section: "priceAverage"},

		types.SettingsItem{Location: "EU", Timestamp: startTime, Period: "month", Section: "vrp"},
		types.SettingsItem{Location: "EU", Timestamp: startTime, Period: "month", Section: "population"},
		types.SettingsItem{Location: "EU", Timestamp: startTime, Period: "month", Section: "currency"},
	}

	SECTIONS = []string{
		"vvp",
		"vrp",
		"currency",
		"population",
		"retail",
		"priceProducer",
		"priceConsumer",
		"priceAverage",
		// "contracts",
	}
)

func getParser(country string) (value *types.Parser, err error) {
	switch country {
	case "ru":
		return fedstat.New(), nil
	case "eu":
		return eurostat.New(), nil
	case "kz":
		return kazstat.New(), nil
	case "az":
		return azstat.New(), nil
	case "us":
		return usastat.New(), nil
	case "de":
		return destat.New(), nil
	case "cz":
		return czstat.New(), nil
	case "be":
		return belstat.New(), nil

	default:
		return nil, nil
	}
}
