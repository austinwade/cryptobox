package renderer

import ("github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/shibukawa/nanovgo"
	"container/list"
)

var windowWidth int
var windowHeight int
var context *nanovgo.Context

func Init(width, height int ) {
	context, _ = nanovgo.NewContext(nanovgo.AntiAlias)

	createFont(context)

	windowWidth, windowHeight = width, height
}

func createFont(context *nanovgo.Context) {
	sourcePath := "src/github.com/austinwade/cryptobox/"
	robotoRegularFileName := "Roboto-Regular.ttf"

	textFont := context.CreateFont("sans", sourcePath + robotoRegularFileName)

	if textFont < 0 {
		panic("Could not find font: " + robotoRegularFileName)
	}
}

func Draw(window *glfw.Window, currencyValues list) {
	wipeWindow(window)

	context.BeginFrame(windowWidth, windowHeight, 1)

	queuePosition := 0
	for value := currencyValues.Front(); value != nil; value = value.Next() {

		drawValue(context, queuePosition, value)
		queuePosition++
	}

	context.EndFrame()
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

func cpToUTF8(cp int) string {
	return string([]rune{rune(cp)})
}

func drawValue(context *nanovgo.Context, index int, value string) {
	x, y := float32(100), float32(40)

	context.BeginPath()
	context.SetFontSize(50.0)
	context.SetFontFace("sans")

	context.SetFontBlur(1.0)
	context.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
	context.Text(x, y, btcUsd)

	context.SetFontBlur(0.0)
	context.SetFillColor(nanovgo.RGBA(255, 255, 255, 250))
	context.Text(x, y, btcUsd)

	context.SetFontBlur(1.0)
	context.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
	context.Text(x + 100, y, ethUsd)

	context.SetFontBlur(0.0)
	context.SetFillColor(nanovgo.RGBA(255, 255, 255, 250))
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






