package main

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.5-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	width  = 512
	height = 512
)

func main() {
	runtime.LockOSThread()

	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		panic(err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4) // OR 2
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(width, height, "Window", nil, nil)
	if err != nil {
		panic(err)
	}
	window.MakeContextCurrent()
	defer glfw.Terminate()

	// Create OpenGL context
	if err := gl.Init(); err != nil {
		panic(err)
	}
	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	gl.ClearColor(0.4, 0.2, 0.5, 1.0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
