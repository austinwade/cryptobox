package currency

import (
	"strconv"
	"fmt"
	"net/http"
	"io/ioutil"
)

const apiUrl = "http://api.etherscan.io/api?module=stats&action=ethprice"

func GetCurrencyValues() (string, string) {
	apiJson := getApiJson()

	etherUsdValue, etherToBitcoin := parseJson(apiJson)

	etherToUsdFloat, _ := strconv.ParseFloat(etherUsdValue, 32)

	etherToBitcoinFloat, _ := strconv.ParseFloat(etherToBitcoin, 32)

	bitcoinUsdValue := (etherToUsdFloat / etherToBitcoinFloat)

	bitcoinValue := strconv.FormatFloat(bitcoinUsdValue, 'f', 2, 32)

	return etherUsdValue, bitcoinValue
}

func getApiJson() string {
	resp, err := http.Get(apiUrl)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	json := string(body[:])

	return json
}

func parseJson(json string) (eth string, btc string) {
	etherUsdValue := json[100:105]
	etherToBitcoin := json[49:56]

	return etherUsdValue, etherToBitcoin
}

