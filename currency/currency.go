package currency

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/Jeffail/gabs"
	//"strconv"
	"strconv"
)

const apiUrl = "https://poloniex.com/public?command=returnTicker"

type stats struct{
	usDollarValue string
	percentChange float32
}

type MarketStats map[string] stats

var MarketStats MarketStats

func init() {
	UpdateCoinStats()
}

func UpdateCoinStats() {
	rawApiJson := []byte(callPoloniexApi())

	jsonParsed, _ := gabs.ParseJSON(rawApiJson)

	coinStatsMap := getStatsMap(jsonParsed)

	MarketStats = coinStatsMap
}

func callPoloniexApi() string {
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	json := string(body[:])

	return json
}

func getStatsMap(jsonParsed *gabs.Container) (statsMap MarketStats) {
	btcStats := getStats(jsonParsed, "USDT_BTC")
	ethStats := getStats(jsonParsed, "USDT_ETH")
	xmrStats := getStats(jsonParsed, "USDT_XMR")

	statsMap["BTC"] = btcStats
	statsMap["ETH"] = ethStats
	statsMap["XMR"] = xmrStats

	return statsMap
}

func getStats(jsonParsed *gabs.Container, market string) (currencyStats stats) {
	currencyStats = stats {
		usDollarValue: jsonParsed.Path(market + ".last").Data().(string),
		percentChange: getPercentChange(jsonParsed, market),
	}

	return currencyStats
}

func getPercentChange(jsonParsed *gabs.Container, market string) (value float64) {
	// Parsing as float32 using "gabs" does not work correctly, so we must go
	// from string, to float32, back to string
	rawString, _ := jsonParsed.Path(market + ".percentChange").Data().(string)

	float, _ := strconv.ParseFloat(rawString, 32)
	value = strconv.FormatFloat(float, 'f', 4, 32)

	return value
}

