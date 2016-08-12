package main//

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

func main() {
	//json := getJSONFromApi()

	ethusd, btcusd := "1", "2"//getCryptoValues(json)

	fmt.Println("btc: " + btcusd)
	fmt.Println("eth: " + ethusd)

	err := glfw.Init(gl.ContextWatcher)
	if err != nil {
		panic(err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Samples, 4)

	window, err := glfw.CreateWindow(1000, 600, "NanoVGo", nil, nil)
	if err != nil {
		panic(err)
	}
	window.SetKeyCallback(key)
	window.MakeContextCurrent()

	ctx, err := nanovgo.NewContext(nanovgo.AntiAlias /*nanovgo.AntiAlias | nanovgo.StencilStrokes | nanovgo.Debug*/)
	defer ctx.Delete()

	font := ctx.CreateFont("sans", "github.com/austinwade/cryptobox/Roboto-Regular.ttf")

	if font < 0 {
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
		if premult {
			gl.ClearColor(0, 0, 0, 0)
		} else {
			gl.ClearColor(1, 1, 1, 1)
		}
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
		gl.Enable(gl.BLEND)
		gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
		gl.Enable(gl.CULL_FACE)
		gl.Disable(gl.DEPTH_TEST)

		ctx.BeginFrame(winWidth, winHeight, 1)

		x, y := float32(30), float32(20)

		ctx.BeginPath()

		ctx.SetFontSize(150.0)
		ctx.SetFontFace("sans")

		ctx.SetTextAlign(nanovgo.AlignLeft | nanovgo.AlignMiddle)

		ctx.SetFontBlur(2)
		ctx.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
		ctx.Text(x+50,y+75, "$ " + ethusd)

		ctx.SetFontBlur(0)
		ctx.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
		ctx.Text(x+50,y+75, "$ " + btcusd)

		ctx.EndFrame()

		gl.Enable(gl.DEPTH_TEST)
		window.SwapBuffers()
		glfw.PollEvents()
	}
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