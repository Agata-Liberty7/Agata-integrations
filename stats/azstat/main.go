package azstat

import (
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	COUNTRY = "AZ"
	parser  = types.Parser{
		COUNTRY: COUNTRY,
		Get: map[string](func() ([]types.Stat, error)){
			// "vvp":           GetVVP,
			// "currency":   GetCurrency,
			"vrp":        GetVRP,
			"population": GetPopulation,
			"retail":     GetRetail,
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
