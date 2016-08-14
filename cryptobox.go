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
	IconBITCOIN       = 0xF15A
)

func main() {
	json := getJSONFromApi()

	ethusd, btcusd := getCryptoValues(json)

	fmt.Println("btc: " + btcusd)
	fmt.Println("eth: " + ethusd)

	err := glfw.Init(gl.ContextWatcher)
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Samples, 4)

	window, err := glfw.CreateWindow(300, 200, "NanoVGo", nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetKeyCallback(key)
	window.MakeContextCurrent()

	ctx, err := nanovgo.NewContext(nanovgo.AntiAlias /*nanovgo.AntiAlias | nanovgo.StencilStrokes | nanovgo.Debug*/)
	defer ctx.Delete()

	textFont := ctx.CreateFont("sans", "github.com/austinwade/cryptobox/Roboto-Regular.ttf")
	iconFont := ctx.CreateFont("icon", "github.com/austinwade/cryptobox/fontawesome-webfont.ttf")

	if textFont < 0 || iconFont < 0 {
		panic("Could not find font")
	}

	if err != nil {
		panic(err)
	}

	glfw.SwapInterval(0)

	for !window.ShouldClose() {
		//t, _ := fps.UpdateGraph()

		fbWidth, fbHeight := window.GetFramebufferSize()
		winWidth, winHeight := window.GetSize()
		//mx, my := window.GetCursorPos()

		gl.Viewport(0, 0, fbWidth, fbHeight)
		gl.ClearColor(1, 1, 1, 1)
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		gl.Enable(gl.CULL_FACE)
		gl.Disable(gl.DEPTH_TEST)

		ctx.BeginFrame(winWidth, winHeight, 1)

		ctx.BeginPath()

		drawCurrencyIcons(ctx)

		drawCurrencyValues(ctx, btcusd, ethusd)

		ctx.EndFrame()

		gl.Enable(gl.DEPTH_TEST)
		window.SwapBuffers()
		glfw.PollEvents()
	}
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
	context.Text(x,y+50, ethUsd)
}

func drawCurrencyIcons(context *nanovgo.Context) {
	x, y := float32(50), float32(100)

	context.SetFontSize(36.0)
	context.SetFontFace("icon")

	context.SetTextAlign(nanovgo.AlignLeft)

	context.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
	context.Text(x,y, cpToUTF8(IconBITCOIN))

	context.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
	context.Text(x,y+50, "e")
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

func getJSONFromApi() string {
	var apiUrl string = "http://api.etherscan.io/api?module=stats&action=ethprice"

    	resp, err := http.Get(apiUrl)

	if err != nil {
		fmt.Println(err)
	}

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
