package renderer

import ("github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/shibukawa/nanovgo"
	"github.com/austinwade/cryptobox/currency"
	"strconv"
	"strings"
)

var windowWidth = 1950
var windowHeight = 70

var context *nanovgo.Context
var marqueePositionOne = 2000.0
var marqueePositionTwo = 4625.0

var blowup bool
var premult bool

func Init(width, height int) {
	context, _ = nanovgo.NewContext(nanovgo.AntiAlias)

	createFont(context)
}

func InitializeWindow() (*glfw.Window) {
	err := glfw.Init(gl.ContextWatcher)

	if err != nil {
		panic(err)
	}

	window, _ := glfw.CreateWindow(windowWidth, windowHeight, "Cryptobox", nil, nil)

	window.SetKeyCallback(key)
	window.MakeContextCurrent()

	Init(windowWidth, windowHeight)

	glfw.SwapInterval(0)

	return window
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

func Draw(window *glfw.Window, marketStats currency.Market) {
	wipeWindow(window)

	context.BeginFrame(windowWidth, windowHeight, 1)

	drawStats(marketStats, marqueePositionOne)
	drawStats(marketStats, marqueePositionTwo)

	marqueePositionOne -= 2
	marqueePositionTwo -= 2

	if (marqueePositionOne < -3250) {
		marqueePositionOne = 2000.0
	}

	if (marqueePositionTwo < -3250) {
		marqueePositionTwo = 2000.0
	}

	context.EndFrame()
}

func createFont(context *nanovgo.Context) {
	sourcePath := "/Users/austin/Code/goWorkspace/src/github.com/austinwade/cryptobox/"
	fontFileName := "Roboto-Medium.ttf"

	textFont := context.CreateFont("sans", sourcePath + fontFileName)

	if textFont < 0 {
		panic("Could not find font: " + fontFileName)
	}
}

func wipeWindow(window *glfw.Window) {
	fbWidth, fbHeight := window.GetFramebufferSize()
	windowWidth, windowHeight = fbWidth, fbHeight

	gl.Viewport(0, 0, fbWidth, fbHeight)
	gl.ClearColor(0, 0, 0, 1)
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT | gl.STENCIL_BUFFER_BIT)
	gl.Enable(gl.BLEND)
	gl.BlendFunc(gl.SRC_ALPHA, gl.ONE_MINUS_SRC_ALPHA)
	gl.Enable(gl.CULL_FACE)
	gl.Disable(gl.DEPTH_TEST)
}

func drawStats(marketStats currency.Market, marqueePosition float64) {
	initFont()

	white := nanovgo.RGB(255, 255, 255)

	keys := [5]string{"BTC", "ETH", "XMR", "DSH", "LTC"}

	queuePosition := float32(0.0)
	for _, key := range keys {

		x := float32(marqueePosition) + (queuePosition * 525.0)

		drawText(key, x, white)
		drawValue(marketStats[key].UsDollarValue, x + 100, white)
		drawPercentChange(marketStats[key].PercentChange, x + 325)

		queuePosition += 1.0
	}
}

func initFont() {
	context.BeginPath()
	context.SetFontSize(50.0)
	context.SetFontFace("sans")
}

func drawText(text string, x float32, color nanovgo.Color) {
	y := float32(50)

	context.SetFontBlur(1.0)
	context.SetFillColor(color)
	context.Text(x, y, text)

	context.SetFontBlur(0.0)
	context.SetFillColor(color)
	context.Text(x, y, text)
}

func drawValue(value string, x float32, color nanovgo.Color) {
	// Truncate, leaving only 4 decimal places
	float, _ := strconv.ParseFloat(value, 64)

	value = getSevenDigitValueStr(float)

	value = "$" + value

	drawText(value, x, color)
}

func getSevenDigitValueStr(value float64) (valueStr string) {
	valueStr = strconv.FormatFloat(value, 'g', 7, 64)

	if (!strings.Contains(valueStr, ".")) {
		valueStr += "."
	}

	totalDigits := len(valueStr)
	if (totalDigits != 8) {
		for i := 0; i < (8-totalDigits); i++  {
			valueStr += "0"
		}
	}

	return valueStr
}

func drawPercentChange(percent string, x float32) {
	// Truncate, leaving only 4 decimal places
	float, _ := strconv.ParseFloat(percent, 32)

	// Make the percent out of 100 instead of 1
	float = float * 100

	percent = strconv.FormatFloat(float, 'f', 2, 32)

	if (float >= 0) {
		percent = "+" + percent + "%"
		green := nanovgo.RGB(23, 151, 85)
		drawText(percent, x, green)
	} else {
		percent = percent + "%"
		red := nanovgo.RGB(217, 71, 85)
		drawText(percent, x, red)
	}
}