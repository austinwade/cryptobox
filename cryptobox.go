package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"github.com/amortaza/go-bellina"
)

func main() {
	json := getJSONFromApi()

	ethusd, btcusd := getCryptoValues(json)

	bl.Start( 1024, 768, "Bellina v0.2", init_, tick, uninit )

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

func tick() {

	bl.Root()
	{
		bl.Pos(64,64)
		bl.Dim(800,600)
		bl.Color(.3,.5,.5)
		bl.Flag(bl.Z_COLOR_SOLID | bl.Z_BORDER_ALL)

		bl.Font("arial", 6)
		bl.FontColor(1,1,1)
		bl.FontNudge(3,3)
		bl.Label("Hello world")

		bl.BorderThickness([]int32{2,2,2,2})
		bl.BorderColor(1,1,1)

		bl.On("hover", func(i interface{}){
			e := i.(*mouse_hover.Event)

			if e.IsInEvent {
			}
		})

		bl.Div()
		{
			bl.ID("red")
			bl.Pos(60, 60)
			bl.Dim(164,148)
			bl.Color(.1,0,.0)
			bl.BorderThickness([]int32{1,1,1,1})
			bl.BorderColor(1,1,1)
			bl.BorderTopsCanvas()

			bl.On("hover", func(i interface{}){
				e := i.(*mouse_hover.Event)

				if e.IsInEvent {
				}
			})

			button.ID("btn", func() {
				fmt.Println("wow")
			})
			{
				//button.Label("SHazzy")
				//button.OnClick(func() {
				//	fmt.Println("click")
				//})
			}
			//button.End()
		}
		bl.End()
	}
	bl.End()
}

func uninit() {
}

func init() {
	runtime.LockOSThread()
}
