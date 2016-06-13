package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {
	json := getJSONFromApi()
	
	fmt.Printf(json)

}

func getJSONFromApi() string {
	var etherscanApiUrl string = "https://api.etherscan.io/api?module=stats&action=ethprice"

    resp, err := http.Get(etherscanApiUrl)

	if err != nil {fmt.Printf("Whoops, somthing went wrong.")}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	json := string(body[:])
	return json
}
