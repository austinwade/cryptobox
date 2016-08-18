package bitcoin

import (
	"strconv"
	"fmt"
	"io/ioutil"
)

const apiUrl = "http://api.etherscan.io/api?module=stats&action=ethprice"

func getCurrencyValues() (string, string) {
	apiJson := getApiJson()

	etherUsdValue := apiJson[100:105]
	etherToBitcoin := apiJson[49:56]

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

