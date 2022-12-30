package eurostat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	COUNTRY = "EU"
	parser  = types.Parser{
		COUNTRY: COUNTRY,
		Get: map[string](func() ([]types.Stat, error)){
			// "vvp":           GetVVP,
			"vrp":        GetVRP,
			"currency":   GetCurrency,
			"population": GetPopulation,
			// "retail":        GetRetail,
			// "priceProducer": GetPriceProducer,
			// "priceConsumer": GetPriceConsumer,
			// "priceAverage":  GetPriceAverage,
			// "contracts":     GetContracts,
		},
	}
)

func New() *types.Parser {
	return &parser
}

type EuroStatDimension struct {
	Label    string `json:"label"`
	Category struct {
		Index map[string]int    `json:"index"`
		Label map[string]string `json:"label"`
	} `json:"category"`
}

type EuroStatResponse struct {
	Label     string            `json:"label"`
	Value     map[string]string `json:"value"`
	Dimension struct {
		Geo      EuroStatDimension `json:"geo"`
		Currency EuroStatDimension `json:"currency"`
	} `json:"dimension"`
}

func processEuroStatResponse(resp *http.Response, parentErr error) (res *EuroStatResponse, err error) {
	defer resp.Body.Close()

	if err != nil {
		return nil, parentErr
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &res)
	return
}
