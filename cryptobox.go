package main

import (
	"github.com/goxjs/gl"
	"github.com/goxjs/glfw"
	"github.com/austinwade/cryptobox/renderer"
	"runtime"
	"github.com/austinwade/cryptobox/bitcoin"
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
	err := glfw.Init(gl.ContextWatcher)

	if err != nil {
		panic(err)
	}

	window, _ := glfw.CreateWindow(windowWidth, windowHeight, "Cryptobox", nil, nil)

	window.SetKeyCallback(key)
	window.MakeContextCurrent()

	renderer.Init(windowWidth, windowHeight)

	glfw.SwapInterval(0)

	loop(window)
}


func loop(window *glfw.Window) {

	for !window.ShouldClose() {

		etherValue, bitcoinValue := bitcoin.GetCurrencyValues()

		renderer.Draw(window, etherValue, bitcoinValue)

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