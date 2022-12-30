package types

import (
	"log"
	"math/rand"

	"gorm.io/gorm"
)

type Stat struct {
	gorm.Model
	Timestamp string `json:"timestamp"`
	Location  string `json:"location"`
	Name      string `json:"name"`
	Value     string `json:"value"`
	Unit      string `json:"unit"`
	Country   string `json:"country" gorm:"not null"`
	Section   string `json:"section" gorm:"not null"`
}
type Stats []Stat

type Metadata struct {
	Count uint32 `json:"count"`
	Time  string `json:"time"`
}

type MetadataRaw struct {
	Section string `json:"section"`
	Count   uint32 `json:"count"`
	Time    string `json:"time"`
}

type SettingsItem struct {
	Timestamp string `json:"timestamp"`
	Location  string `json:"location"`
	Section   string `json:"section"`
	Period    string `json:"period"`
}
type SettingsItems []SettingsItem

type StatResponse struct {
	Timestamp string `json:"timestamp,omitempty"`
	Location  string `json:"location,omitempty"`
	Name      string `json:"name,omitempty"`
	Value     string `json:"value,omitempty"`
	Unit      string `json:"unit,omitempty"`
}
type StatResponses []StatResponse

type Parser struct {
	COUNTRY  string
	SECTIONS []string
	Get      map[string](func() ([]Stat, error))
	// LoadAll      func() error
	FetchSection func(section string) (value []Stat, err error)
}

func (p Parser) LoadAll() (err error) {
	for section := range p.Get {
		_, err := p.Get[section]()
		if err != nil {
			log.Println("Error with", section, err)
		}
	}
	return err
}

func (p Parser) AvailableSections() (sections []string) {
	for section := range p.Get {
		sections = append(sections, section)
	}
	return
}

func (s StatResponses) PeakRandom(count int) (result []StatResponse) {
	for i := 0; i < count && i < len(s)/2; i++ {
		index := rand.Intn(len(s))
		result = append(result, s[index])
	}
	return
}

func (s Stats) PeakRandom(count int) (result []Stat) {
	for i := 0; i < count && i < len(s)/2; i++ {
		index := rand.Intn(len(s))
		result = append(result, s[index])
	}
	return
}
