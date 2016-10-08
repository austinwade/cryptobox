package main

import (
	"github.com/goxjs/glfw"
	//"runtime"
	"github.com/austinwade/cryptobox/currency"
	"github.com/austinwade/cryptobox/renderer"
	"time"
)

const (
	windowWidth = 1950
	windowHeight = 70
)

func init() {
	//runtime.LockOSThread()
}

func main() {
	window := renderer.InitializeWindow()

	loop(window)
}

func loop(window *glfw.Window) {

	marketStats := currency.MarketStats
	statsLastUpdated := time.Now()

	for !window.ShouldClose() {

		if hasOneMinutePassed(statsLastUpdated) {
			currency.UpdateMarketStats()
			statsLastUpdated = time.Now()
			marketStats = currency.MarketStats
		}

		renderer.Draw(window, marketStats)

		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func hasOneMinutePassed(timeToTest time.Time) (bool) {
	oneMinuteLater := timeToTest.Add(time.Minute)

	if (oneMinuteLater.Before(time.Now())) {
		return true
	}

	return false
}