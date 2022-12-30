package fedstat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/Agata-Liberty7/Agata-integrations/stats/types"
)

var (
	COUNTRY = "RU"

	parser = types.Parser{
		COUNTRY: COUNTRY,
		Get: map[string](func() ([]types.Stat, error)){
			"vvp":           GetVVP,
			"vrp":           GetVRP,
			"currency":      GetCurrency,
			"population":    GetPopulation,
			"retail":        GetRetail,
			"priceProducer": GetPriceProducer,
			"priceConsumer": GetPriceConsumer,
			"priceAverage":  GetPriceAverage,
			"contracts":     GetContracts,
		},
	}

	currentYear  = time.Now().Year()
	MONTH_CODES  = []string{"1540283", "1540282", "1540236", "1540229", "1540235", "1540234", "1540233", "1540228", "1540276", "1540273", "1540272", "1540230"}
	MONTH_TITLES = []string{"январь", "февраль", "март", "апрель", "май", "июнь", "июль", "август", "сентябрь", "октябрь", "ноябрь", "декабрь"}
)

func New() *types.Parser {
	return &parser
}

type FesStatRequest struct {
	id                string
	lineObjectIds     []string
	columnObjectIds   []string
	selectedFilterIDs []string
}

type FedStatResponse struct {
	Results []interface{} `json:"results"`
	Count   string        `json:"__count"`
}

func getFedStatRequest(params FesStatRequest) (req *http.Request, err error) {
	payload := url.Values{}
	for _, value := range params.lineObjectIds {
		payload.Add("lineObjectIds", value)
	}
	for _, value := range params.columnObjectIds {
		payload.Add("columnObjectIds", value)
	}
	for _, value := range params.selectedFilterIDs {
		payload.Add("selectedFilterIds", value)
	}

	req, err = http.NewRequest(
		"GET",
		"https://www.fedstat.ru/indicator/dataGrid.do?id="+params.id,
		bytes.NewBufferString(payload.Encode()),
	)

	return req, err
}

func processFedStatResponse(resp *http.Response, parentErr error) (res *FedStatResponse, err error) {
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

func buildTimestamp(year int, month int) time.Time {
	return time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
}
