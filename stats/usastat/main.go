package usastat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	COUNTRY = "US"
	parser  = types.Parser{
		COUNTRY: COUNTRY,
		Get: map[string](func() ([]types.Stat, error)){
			"vrp":        GetVRP,
			"population": GetPopulation,
			// "vvp":           GetVVP,
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

type UsaStatDataItem struct {
	State      string `json:"State"`
	Year       int    `json:"ID Year"`
	Population int    `json:"Population"`
}

type UsaStatResponse struct {
	Data []UsaStatDataItem `json:"data"`
}

func processUsaStatResponse(resp *http.Response, parentErr error) (res *UsaStatResponse, err error) {
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
