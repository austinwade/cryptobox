package main

import (
	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"runtime"
	"github.com/austinwade/cryptobox/currency"
	"github.com/austinwade/cryptobox/renderer"
	"time"
	"fmt"
)

var blowup bool
var premult bool

const (
	windowWidth = 1950
	windowHeight = 70
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window := initializeWindow()

	loop(window)
}

func initializeWindow() (*glfw.Window) {
	err := glfw.Init(gl.ContextWatcher)

	if err != nil {
		panic(err)
	}

	window, _ := glfw.CreateWindow(windowWidth, windowHeight, "Cryptobox", nil, nil)

	window.SetKeyCallback(key)
	window.MakeContextCurrent()

	renderer.Init(windowWidth, windowHeight)

	glfw.SwapInterval(0)

	return window
}

func loop(window *glfw.Window) {

	currency.UpdateCoinStats()
	coinStats := currency.MarketStats
	statsLastUpdated := time.Now()

	for !window.ShouldClose() {

		if hasOneMinutePassed(statsLastUpdated) {
			currency.UpdateCoinStats()
			statsLastUpdated = time.Now()
			coinStats = currency.MarketStats
			fmt.Println(currency.MarketStats)
		}

		renderer.Draw(window, coinStats)

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

func key(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {

	if key == glfw.KeyEscape && action == glfw.Press {
		w.SetShouldClose(true)

	} else if key == glfw.KeySpace && action == glfw.Press {
		blowup = !blowup

	} else if key == glfw.KeyP && action == glfw.Press {
		premult = !premult
	}
}