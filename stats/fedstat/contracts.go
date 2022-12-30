package fedstat

import "github.com/Agata-Liberty7/Agata-integrations/stats/types"

var (
	INDICATOR_CONTRACTS_PATH = "https://www.fedstat.ru/indicator/dataGrid.do?id=43269&lineObjectIds=0&lineObjectIds=33560&lineObjectIds=33560&lineObjectIds=57831&lineObjectIds=58274&lineObjectIds=58333&lineObjectIds=30611&columnObjectIds=3&selectedFilterIds=0_43269&selectedFilterIds=3_2011&selectedFilterIds=3_2012&selectedFilterIds=3_2013&selectedFilterIds=3_2014&selectedFilterIds=3_2015&selectedFilterIds=3_2016&selectedFilterIds=3_2017&selectedFilterIds=3_2018&selectedFilterIds=3_2019&selectedFilterIds=3_2020&selectedFilterIds=30611_950475&selectedFilterIds=33560_1558883&selectedFilterIds=57831_1688487&selectedFilterIds=57831_1688488&selectedFilterIds=57831_1688489&selectedFilterIds=57831_1688490&selectedFilterIds=57831_1688491&selectedFilterIds=57831_1688492&selectedFilterIds=57831_1688493&selectedFilterIds=57831_1688494&selectedFilterIds=57831_1688495&selectedFilterIds=57831_1688496&selectedFilterIds=57831_1688497&selectedFilterIds=57831_1688498&selectedFilterIds=57831_1688499&selectedFilterIds=57831_1688500&selectedFilterIds=57831_1688501&selectedFilterIds=57831_1688502&selectedFilterIds=57831_1688503&selectedFilterIds=57831_1688504&selectedFilterIds=57831_1688505&selectedFilterIds=57831_1688506&selectedFilterIds=57831_1688507&selectedFilterIds=57831_1688508&selectedFilterIds=57831_1688509&selectedFilterIds=57831_1688510&selectedFilterIds=57831_1688511&selectedFilterIds=57831_1688513&selectedFilterIds=57831_1688514&selectedFilterIds=57831_1688515&selectedFilterIds=57831_1688516&selectedFilterIds=57831_1688517&selectedFilterIds=57831_1688518&selectedFilterIds=57831_1688519&selectedFilterIds=57831_1688521&selectedFilterIds=57831_1688522&selectedFilterIds=57831_1688523&selectedFilterIds=57831_1688524&selectedFilterIds=57831_1688525&selectedFilterIds=57831_1688526&selectedFilterIds=57831_1688527&selectedFilterIds=57831_1688528&selectedFilterIds=57831_1688529&selectedFilterIds=57831_1688530&selectedFilterIds=57831_1688531&selectedFilterIds=57831_1688532&selectedFilterIds=57831_1688533&selectedFilterIds=57831_1688534&selectedFilterIds=57831_1688535&selectedFilterIds=57831_1688536&selectedFilterIds=57831_1688537&selectedFilterIds=57831_1688538&selectedFilterIds=57831_1688539&selectedFilterIds=57831_1688540&selectedFilterIds=57831_1688541&selectedFilterIds=57831_1688542&selectedFilterIds=57831_1688543&selectedFilterIds=57831_1688545&selectedFilterIds=57831_1688546&selectedFilterIds=57831_1688547&selectedFilterIds=57831_1688548&selectedFilterIds=57831_1688549&selectedFilterIds=57831_1688550&selectedFilterIds=57831_1688551&selectedFilterIds=57831_1688552&selectedFilterIds=57831_1688553&selectedFilterIds=57831_1688554&selectedFilterIds=57831_1688555&selectedFilterIds=57831_1688556&selectedFilterIds=57831_1688557&selectedFilterIds=57831_1688559&selectedFilterIds=57831_1688560&selectedFilterIds=57831_1688561&selectedFilterIds=57831_1688562&selectedFilterIds=57831_1688563&selectedFilterIds=57831_1688564&selectedFilterIds=57831_1688565&selectedFilterIds=57831_1688566&selectedFilterIds=57831_1688568&selectedFilterIds=57831_1688571&selectedFilterIds=57831_1688573&selectedFilterIds=57831_1688574&selectedFilterIds=57831_1688575&selectedFilterIds=57831_1688576&selectedFilterIds=57831_1688577&selectedFilterIds=57831_1688578&selectedFilterIds=57831_1688579&selectedFilterIds=57831_1688581&selectedFilterIds=57831_1688582&selectedFilterIds=57831_1688583&selectedFilterIds=57831_1688584&selectedFilterIds=57831_1688585&selectedFilterIds=57831_1688586&selectedFilterIds=57831_1688587&selectedFilterIds=57831_1692937&selectedFilterIds=57831_1692938&selectedFilterIds=57831_1692939&selectedFilterIds=57831_1692940&selectedFilterIds=57831_1695534&selectedFilterIds=57831_1697988&selectedFilterIds=57831_1707677&selectedFilterIds=57831_1751201&selectedFilterIds=57831_1795276&selectedFilterIds=57831_1795277&selectedFilterIds=58274_1744472&selectedFilterIds=58333_1710011&selectedFilterIds=58333_1710012&selectedFilterIds=58333_1710014&selectedFilterIds=58333_1710016&selectedFilterIds=58333_1710018&selectedFilterIds=58333_1789097&selectedFilterIds=58333_1789098&selectedFilterIds=58333_1789099&selectedFilterIds=58333_1790445&selectedFilterIds=58333_1790446&selectedFilterIds=58333_1790447&selectedFilterIds=58333_1790448&selectedFilterIds=58333_1790449&selectedFilterIds=58333_1790450&selectedFilterIds=58333_1790451&selectedFilterIds=58333_1790452&selectedFilterIds=58333_1790453&selectedFilterIds=58333_1790454&selectedFilterIds=58333_1790455&selectedFilterIds=58333_1790456&selectedFilterIds=58333_1836028"
)

type FedStatContracts struct {
	Year  string `json:"year"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Value string `json:"value"`
}

func GetContracts() (contracts []types.Stat, err error) {
	// res, err := processFedStatResponse(http.Get(INDICATOR_CONTRACTS_PATH))
	// if err != nil {
	// 	return nil, err
	// }

	// for _, r := range res.Results {
	// 	v := r.(map[string]interface{})

	// 	for field, val := range v {

	// 		for year := 2012; year <= currentYear; year++ {
	// 			if strings.Contains(field, "dim"+strconv.Itoa(year)) {
	// 				contracts = append(contracts, FedStatContracts{
	// 					Year:  strconv.Itoa(year),
	// 					Title: v["dim57831"].(string),
	// 					Type:  v["dim58333"].(string),
	// 					Value: val.(string),
	// 				})
	// 			}
	// 		}
	// 	}
	// }

	return contracts, nil
}
