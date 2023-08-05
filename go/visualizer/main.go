package main

import (
	"flag"
	"fmt"
	"math"
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
	// set these by command-line args
	colormap_name := flag.String("cmap", "parula", "Name of colormap to use. To get a list of supported maps, use -list-cmaps flag")
	colormap_list := flag.Bool("list-cmaps", false, "Print list of supported maps")
	window_width := flag.Int("width", 640, "Window width")
	window_height := flag.Int("height", 480, "Window height")
	flag.Parse()

	others := flag.Args() // other args not parsed by flag
	if len(others) > 0 {
		fmt.Printf("Additional args:\n")
		for n := 0; n < len(others); n++ {
			fmt.Printf(" [%v]\n", others[n])
		}
	}

	if *colormap_list {
		print_map_types()
		return
	}

	map_type := map_type_from_string(colormap_name)

	fmt.Printf("Creating window [%vx%v], using %v colormap\n", *window_width, *window_height, map_type.ToString())
	//colormap_info()

	// Initialize GLFW
	var window Window
	if !window.InitGLFW(*window_width, *window_height, key_callback) {
		fmt.Println("Error initializing GLFW")
		return
	}
	// Create render data
	var renderer Renderer
	renderer.Init()

	defer func() {
		renderer.Shutdown()
		window.Close()
	}()

	DrawColormap := func(width, height int) {
		N := 99
		dx := float32(1.0) / float32(N+1)
		dy := float32(0.05)
		ypos := 1.0 - dy

		for n := 0; n <= N; n++ {
			xpos := dx * float32(n)

			color := colormap(float64(xpos), map_type)

			renderer.DrawRect(xpos, ypos, dx, dy, float32(color[0]), float32(color[1]), float32(color[2]))
		}
	}
	DrawHeatBar := func(u []float64, width, height int, min, max float64) {
		N := len(u) - 1
		dx := float32(1.0) / float32(N+1)
		ypos := float32(1.0 / 2.0)

		for n := 0; n <= N; n++ {
			xpos := dx * float32(n)

			T := u[n]
			f := (T - min) / (max - min)

			color := colormap(f, map_type)

			renderer.DrawRect(xpos, ypos, dx, dx, float32(color[0]), float32(color[1]), float32(color[2]))
		}
	}
	u := []float64{
		300, 350, 350, 350, 350, 350, 350, 350, 350, 400,
	}

	var t = 0.0

	gl.ClearColor(0.4, 0.2, 0.5, 1.0)
	for !window.ShouldClose() {
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		t = glfw.GetTime()
		xpos := float32(0.5*math.Cos(t) + 0.5)

		r := float32(0.5*math.Cos(2*t) + 0.5)

		renderer.DrawRect(xpos, 0, 0.1, 0.1, r, 0.5, 0.5)

		DrawColormap(window.width, window.height)
		DrawHeatBar(u, window.width, window.height, 300, 400)

		glfw.PollEvents()
		window.SwapBuffers()
	}
}
