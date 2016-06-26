package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"runtime"
	"github.com/amortaza/go-bellina"
)

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

func tick() {

	bl.Root()
	{
		bl.Pos(64,64)
		bl.Dim(800,600)
		bl.Color(.3,.5,.5)

		bl.On("hover", func(i interface{}){
			e := i.(*mouse_hover.Event)

			if e.IsInEvent {
			}
		})

	}
	bl.End()
}

func uninit() {
}

func init_() {
	runtime.LockOSThread()
}

func main() {
	json := getJSONFromApi()

	ethusd, btcusd := getCryptoValues(json)

	bl.Start( 1024, 768, "Bellina v0.2", init_, tick, uninit )

	fmt.Println("btc: " + btcusd)
	fmt.Println("eth: " + ethusd)
}
