package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	json := getJSONFromApi()

	ethusd, btcusd := getCryptoValues(json)

	fmt.Println("btc: " + btcusd)
	fmt.Println("eth: " + ethusd)
}

func getJSONFromApi() string {
	var apiUrl string = "https://api.etherscan.io/api?module=stats&action=ethprice"

    resp, _ := http.Get(apiUrl)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	json := string(body[:])
	return json
}

func getCryptoValues(json string) (string, string) {
	// todo, parse better
	ethbtc := json[49:56]
	ethusd := json[100:105]

	ethValue, _ := strconv.ParseFloat(ethusd, 32)

	ethToBtc, _ := strconv.ParseFloat(ethbtc, 32)

	btcValue := (ethValue / ethToBtc)

	btcusd := strconv.FormatFloat(btcValue, 'f', 2, 32)

	return ethusd, btcusd
}