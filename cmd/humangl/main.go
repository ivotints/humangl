// cmd/humangl/main.go
package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v3.3-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ivotints/humangl/internal/renderer"
)

const (
	windowWidth  = 800
	windowHeight = 600
)

func init() {
	runtime.LockOSThread() // lock main() to run on main thread. OpenGL and GLFW should run only on main thread. Prevent Go to use gorutines and switch threads
}

func main() {
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err) // print error, add new line, exit(1)
	}
	defer glfw.Terminate() // call glfw.Terminate() when main() is finished

	// Configure GLFW
	glfw.WindowHint(glfw.Resizable, glfw.False)  // turn-off window resize
	glfw.WindowHint(glfw.ContextVersionMajor, 3)  // best for learning 3.3 version
	glfw.WindowHint(glfw.ContextVersionMinor, 3)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)  // modern functions of OpenGL
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)  // exclude old functionality

	// Create window
	window, err := glfw.CreateWindow(windowWidth, windowHeight, "humangl", nil, nil)
	if err != nil {
		log.Fatalln("failed to create window:", err)
	}

	window.MakeContextCurrent()  // makes `window` main window. calls like gl.Clear() will apply to it.

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize OpenGL:", err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version) // to prove that we use OpenGL higher than 3.0 -> 3.3

	// Set up viewport
	gl.Viewport(0, 0, windowWidth, windowHeight) // tells OpenGL window coordinates

	// Draw closer objects on top of deeper ones
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	shader, err := renderer.NewShader(renderer.VertexShaderSource, renderer.FragmentShaderSource)
	if err != nil {
		log.Fatalln("failed to create shader:", err)
	}
	defer shader.Delete()

	// Main render loop
	for !window.ShouldClose() {
		// Clear buffers
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		shader.Use()

		shader.SetVec3f("uColor", 1.0, 0.0, 0.0)

		// draw cube here

		// Swap buffers and poll events
		window.SwapBuffers()
		glfw.PollEvents()
	}
}
