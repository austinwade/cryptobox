package currency

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

const apiUrl = "https://poloniex.com/public?command=returnTicker"

var CoinStats string

func init() {
	UpdateCoinStats()
}

func UpdateCoinStats() {
	CoinStats = CoinStats + "a"
	//coinStatsJson := callPoloniexApi()
	//
	//formattedStatsString := getStatsString(coinStatsJson)
	//
	//CoinStats = formattedStatsString
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

func getStatsString(json string) (statsString string) {

	return "foo"
}

