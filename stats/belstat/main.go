package belstat

import (
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	COUNTRY = "BE"
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
