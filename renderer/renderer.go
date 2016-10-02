package renderer

import ("github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/shibukawa/nanovgo"
)

var windowWidth int
var windowHeight int
var context *nanovgo.Context

func Init(width, height int) {
	context, _ = nanovgo.NewContext(nanovgo.AntiAlias)

	createFont(context)

	windowWidth, windowHeight = width, height
}

func Draw(window *glfw.Window, coinStats string) {
	wipeWindow(window)

	context.BeginFrame(windowWidth, windowHeight, 1)

	drawStats(context, coinStats)

	context.EndFrame()
}

func createFont(context *nanovgo.Context) {
	sourcePath := "src/github.com/austinwade/cryptobox/"
	robotoRegularFileName := "Roboto-Regular.ttf"

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

func drawStats(context *nanovgo.Context, coinStats string) {
	x, y := float32(100), float32(45)

	context.BeginPath()
	context.SetFontSize(50.0)
	context.SetFontFace("sans")

	context.SetFontBlur(1.0)
	context.SetFillColor(nanovgo.RGBA(0, 0, 0, 255))
	context.Text(x, y, coinStats)

	context.SetFontBlur(0.0)
	context.SetFillColor(nanovgo.RGBA(255, 255, 255, 250))
	context.Text(x, y, coinStats)
}