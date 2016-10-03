package renderer

import ("github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/shibukawa/nanovgo"
	"github.com/austinwade/cryptobox/currency"
)

var windowWidth int
var windowHeight int
var context *nanovgo.Context

func Init(width, height int) {
	context, _ = nanovgo.NewContext(nanovgo.AntiAlias)

	createFont(context)

	windowWidth, windowHeight = width, height
}

func Draw(window *glfw.Window, marketStats currency.Market) {
	wipeWindow(window)

	context.BeginFrame(windowWidth, windowHeight, 1)

	drawStats(context, marketStats)

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
	x, y := float32(100), float32(50)

	context.BeginPath()
	context.SetFontSize(50.0)
	context.SetFontFace("sans")

	drawText(x, y, marketStats["BTC"].UsDollarValue)
}

func drawText(x float32, y float32, text string) {
	context.SetFontBlur(1.0)
	context.SetFillColor(nanovgo.RGBA(255, 255, 255, 255))
	context.Text(x, y, text)

	context.SetFontBlur(0.0)
	context.SetFillColor(nanovgo.RGBA(255, 255, 255, 250))
	context.Text(x, y, text)
}