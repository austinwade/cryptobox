package renderer

import ("github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/shibukawa/nanovgo")

const (	bitcoinIconId       = 0xF15A
	sourcePath = "src/github.com/austinwade/cryptobox/"
	robotoRegularFileName = "Roboto-Regular.ttf"
	fontAwesomeFileName = "fontawesome-webfont.ttf"
)

var windowWidth, windowHeight int
var context *nanovgo.Context

func Init(width, height int ) {
	context, _ = nanovgo.NewContext(nanovgo.AntiAlias)

	createFonts(context)

	windowWidth, windowHeight = width, height
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

func Draw(window *glfw.Window, etherValue, bitcoinValue string) {
	wipeWindow(window)

	context.BeginFrame(windowWidth, windowHeight, 1)

	drawCurrencyValues(context, etherValue, bitcoinValue)

	context.EndFrame()
}

func wipeWindow(window *glfw.Window) {
	fbWidth, fbHeight := window.GetFramebufferSize()
	windowWidth, windowHeight = fbWidth, fbHeight

	gl.Viewport(0, 0, fbWidth, fbHeight)
	gl.ClearColor(0.22, 0.24, 0.24, 1)
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
	x, y := float32(100), float32(40)

	context.BeginPath()
	context.SetFontSize(36.0)
	context.SetFontFace("sans")

	context.SetTextAlign(nanovgo.AlignRight)

	context.SetFillColor(nanovgo.RGBA(255, 255, 255, 255))
	context.Text(x,y, btcUsd)

	context.SetFillColor(nanovgo.RGBA(255, 255, 255, 255))
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






