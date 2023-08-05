package main

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.6-core/gl"
)

type Renderer struct {
	vbo, vao uint32
	shader   uint32
}

func (this *Renderer) Init() {
	// create rect buffer
	rect := []float32{
		0, 0.5, 0,
		-0.5, -0.5, 0,
		0.5, -0.5, 0,
	}

	gl.GenBuffers(1, &this.vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, this.vbo)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(rect), gl.Ptr(rect), gl.STATIC_DRAW)

	gl.GenVertexArrays(1, &this.vao)
	gl.BindVertexArray(this.vao)
	gl.EnableVertexAttribArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, this.vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	// create simple shader
	vertSrc := `
		#version 410
		in vec3 vp;
		void main() {
			gl_Position = vec4(vp, 1.0);
		}
	` + "\x00" // null-terminate so C-bindings can handle it
	fragSrc := `
		#version 410
		out vec4 frag_colour;
		void main() {
			frag_colour = vec4(1, 1, 1, 1);
		}
	` + "\x00" // null-terminate so C-bindings can handle it
	this.shader = CreateShader(vertSrc, fragSrc)
}

func (this *Renderer) DrawRect() {
	gl.UseProgram(this.shader)

	gl.BindVertexArray(this.vao)
	gl.DrawArrays(gl.TRIANGLES, 0, 3)
}

func CreateShader(vertSrc, fragSrc string) uint32 {
	vert := CompileShaderComponent(vertSrc, gl.VERTEX_SHADER)
	if vert == 0 {
		return 0
	}
	frag := CompileShaderComponent(fragSrc, gl.FRAGMENT_SHADER)
	if frag == 0 {
		return 0
	}

	prog := gl.CreateProgram()
	gl.AttachShader(prog, vert)
	gl.AttachShader(prog, frag)
	gl.LinkProgram(prog)

	// check for errors
	var status int32
	gl.GetProgramiv(prog, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(prog, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(prog, logLength, nil, gl.Str(log))

		gl.DeleteProgram(prog)
		gl.DeleteShader(vert)
		gl.DeleteShader(frag)

		fmt.Printf("failed to link: %v", log)

		return 0
	}

	return prog
}

func CompileShaderComponent(src string, shader_type uint32) uint32 {
	shader := gl.CreateShader(shader_type)

	csources, free := gl.Strs(src)
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	// check for errors
	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		fmt.Printf("failed to compile %v: %v", src, log)

		return 0
	}

	return shader
}
