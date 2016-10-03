package currency

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"github.com/Jeffail/gabs"
)

const apiUrl = "https://poloniex.com/public?command=returnTicker"

type MarketProperties struct{
	UsDollarValue string
	PercentChange string
}

type Market map[string] MarketProperties

var MarketStats Market

func init() {
	UpdateMarketStats()
}

func UpdateMarketStats() {
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

func getStatsMap(jsonParsed *gabs.Container) (marketStats Market) {
	btcStats := getStats(jsonParsed, "USDT_BTC")
	ethStats := getStats(jsonParsed, "USDT_ETH")
	xmrStats := getStats(jsonParsed, "USDT_XMR")

	marketStats = Market{}

	marketStats["BTC"] = btcStats
	marketStats["ETH"] = ethStats
	marketStats["XMR"] = xmrStats

	return marketStats
}

func getStats(jsonParsed *gabs.Container, market string) (currencyStats MarketProperties) {
	usDollarValue, _ := jsonParsed.Search(market,"last").Data().(string)
	percentChange, _ := jsonParsed.Search(market, "percentChange").Data().(string)

	currencyStats = MarketProperties {
		UsDollarValue: usDollarValue,
		PercentChange: percentChange,
	}

	return currencyStats
}