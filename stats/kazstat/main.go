package kazstat

import (
	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	COUNTRY = "KZ"
	parser  = types.Parser{
		COUNTRY: COUNTRY,
		Get: map[string](func() ([]types.Stat, error)){
			// "vvp":           GetVVP,
			"vrp":        GetVRP,
			"population": GetPopulation,
			"retail":     GetRetail,
			// "currency":   GetCurrency,
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
