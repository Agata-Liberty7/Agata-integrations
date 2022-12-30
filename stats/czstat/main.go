package czstat

import (
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	COUNTRY = "CZ"
	parser  = types.Parser{
		COUNTRY: COUNTRY,
		Get: map[string](func() ([]types.Stat, error)){
			// "vvp":           GetVVP,
			"vrp":        GetVRP,
			"population": GetPopulation,
			// "currency":   GetCurrency,
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
