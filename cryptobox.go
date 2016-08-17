package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/shibukawa/nanovgo"
)

var blowup bool
var premult bool

const (
	sourcePath = "src/github.com/austinwade/cryptobox/"

	apiUrl = "http://api.etherscan.io/api?module=stats&action=ethprice"

	windowWidth = 300
	windowHeight = 400

	robotoRegularFileName = "Roboto-Regular.ttf"
	fontAwesomeFileName = "fontawesome-webfont.ttf"

	bitcoinIconId       = 0xF15A
)

func main() {
	err := glfw.Init(gl.ContextWatcher)

	if err != nil {
		panic(err)
	}

	window := buildWindow()
	context := buildContext()

	render(window, context)
}

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

func buildWindow() *glfw.Window {
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "Cryptobox", nil, nil)
	if err != nil {
		panic(err)
	}

	window.SetKeyCallback(key)
	window.MakeContextCurrent()

	return window
}

func buildContext() *nanovgo.Context {
	context, err := nanovgo.NewContext(nanovgo.AntiAlias)
	if err != nil {
		panic(err)
	}

	createFonts(context)

	return context
}

func createFonts(context *nanovgo.Context) {
	textFont := context.CreateFont("sans", sourcePath + robotoRegularFileName)
	iconFont := context.CreateFont("icon", sourcePath + fontAwesomeFileName)

	if textFont < 0 {
		panic("Could not find font: " + robotoRegularFileName)
	}

	if iconFont < 0 {
		panic("Could not find font: " + fontAwesomeFileName)
	}
}

func render(window *glfw.Window, context *nanovgo.Context) {

	for !window.ShouldClose() {
		wipeWindow()

		context.BeginFrame(windowWidth, windowHeight, 1)

		context.BeginPath()

		etherValue, bitcoinValue := getCurrencyValues()

		//drawCurrencyIcons(ctx)

		drawCurrencyValues(context, etherValue, bitcoinValue)

		context.EndFrame()

		gl.Enable(gl.DEPTH_TEST)
		window.SwapBuffers()
		glfw.PollEvents()
		glfw.SwapInterval(0)
	}
}

func wipeWindow() {
	fbWidth, fbHeight := window.GetFramebufferSize()

	gl.Viewport(0, 0, fbWidth, fbHeight)
	gl.ClearColor(1, 1, 1, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)
}

func cpToUTF8(cp int) string {
	return string([]rune{rune(cp)})
}

func drawCurrencyValues(context *nanovgo.Context, btcUsd string, ethUsd string) {
	x, y := float32(100), float32(100)

	context.SetFontSize(36.0)
	context.SetFontFace("sans")

	context.SetTextAlign(nanovgo.AlignRight)

	context.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
	context.Text(x,y, btcUsd)

	context.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
	context.Text(x + 100, y, ethUsd)
}

func drawCurrencyIcons(context *nanovgo.Context) {
	x, y := float32(50), float32(100)

	context.SetFontSize(36.0)
	context.SetFontFace("icon")

	context.SetTextAlign(nanovgo.AlignLeft)

	context.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
	context.Text(x,y, cpToUTF8(bitcoinIconId))

	context.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
	context.Text(x+100,y, "e")
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