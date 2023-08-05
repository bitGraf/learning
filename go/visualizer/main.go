package main

import (
	"fmt"
	"runtime"

	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	runtime.LockOSThread()
}

func key_callback(key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
	if (key == glfw.KeyR) && (action == glfw.Press) {
		// reset
	}
}

func main() {

	// Initialize GLFW
	var window Window
	if !window.InitGLFW(640, 480, key_callback) {
		fmt.Println("Error initializing GLFW")
		return
	}
	defer window.Close()

	// Create render data
	var renderer Renderer
	renderer.Init()

	gl.ClearColor(0.4, 0.2, 0.5, 1.0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		renderer.DrawRect()

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
