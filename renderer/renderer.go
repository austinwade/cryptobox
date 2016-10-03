package renderer

import ("github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/shibukawa/nanovgo"
	"github.com/austinwade/cryptobox/currency"
	"strconv"
)

var windowWidth int
var windowHeight int
var context *nanovgo.Context
var marqueePosition = 2000.0

func Init(width, height int) {
	context, _ = nanovgo.NewContext(nanovgo.AntiAlias)

	createFont(context)

	windowWidth, windowHeight = width, height
}

func Draw(window *glfw.Window, marketStats currency.Market) {
	wipeWindow(window)

	context.BeginFrame(windowWidth, windowHeight, 1)

	drawStats(context, marketStats)

	marqueePosition -= 0.2

	context.EndFrame()
}

func createFont(context *nanovgo.Context) {
	sourcePath := "src/github.com/austinwade/cryptobox/"
	robotoRegularFileName := "Roboto-Medium.ttf"

	textFont := context.CreateFont("sans", sourcePath + robotoRegularFileName)

	if textFont < 0 {
		panic("Could not find font: " + robotoRegularFileName)
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

func drawStats(context *nanovgo.Context, marketStats currency.Market) {
	initFont()

	white := nanovgo.RGB(255, 255, 255)

	keys := [3]string{"BTC", "ETH", "XMR"}

	queuePosition := float32(0.0)
	for _, key := range keys {
		x := float32(marqueePosition) + (queuePosition * 500.0)

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

func drawValue(value float64, x float32, color nanovgo.Color) {
	// Truncate, leaving only 4 decimal places
	//float, _ := strconv.ParseFloat(value, 64)

	valueStr := strconv.FormatFloat(value, 'G', 7, 64)

	valueStr = "$" + valueStr

	drawText(valueStr, x, color)
}

func drawPercentChange(percent float64, x float32) {
	// Truncate, leaving only 4 decimal places
	//float, _ := strconv.ParseFloat(percent, 32)

	// Make the percent out of 100 instead of 1
	percent = percent * 100

	percentStr := strconv.FormatFloat(percent, 'f', 2, 32)

	if (percent >= 0) {
		percentStr = "+" + percentStr + "%"
		green := nanovgo.RGB(23, 151, 85)
		drawText(percentStr, x, green)
	} else {
		percentStr = "-" + percentStr + "%"
		red := nanovgo.RGB(217, 71, 85)
		drawText(percentStr, x, red)
	}
}