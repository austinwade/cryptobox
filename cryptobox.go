package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	json := getJSONFromApi()

	// todo, parse better
	ethbtc := json[49:56]
	ethusd := json[100:105]

	ethValue, err := strconv.Atoi(ethbtc)
	if (err != nil) {return}

	ethToBtc, err := strconv.Atoi(ethusd)
	if (err != nil) {return}

	btcValue := (ethValue / ethToBtc)

	btcusd, err := strconv.Itoa(btcValue)

	fmt.Printf("\n")
	fmt.Printf("btc: " + btcusd + "\n")
	fmt.Printf("eth: " + ethusd + "\n")
}

func getJSONFromApi() string {
	var apiUrl string = "https://api.etherscan.io/api?module=stats&action=ethprice"

    resp, err := http.Get(apiUrl)

	if err != nil {fmt.Printf("Whoops, somthing went wrong.")}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	json := string(body[:])
	return json
}
