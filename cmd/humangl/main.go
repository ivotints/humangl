package main

import (
	"fmt"
	"log"
	"math"
	"runtime"
	"time"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
	"github.com/ivotints/humangl/internal/matrix"
	"github.com/ivotints/humangl/internal/model"
	"github.com/ivotints/humangl/internal/renderer"
)

const (
	windowWidth  = 800
	windowHeight = 600
	windowTitle  = "HumanGL - Skeletal Animation"
)

// Camera state
var (
	cameraDistance = float32(5.0)
	cameraAngleX   = float32(0.0)
	cameraAngleY   = float32(0.0)
	lastMouseX     = float64(windowWidth / 2)
	lastMouseY     = float64(windowHeight / 2)
	firstMouse     = true
)

// Animation state
var (
	lastTime = time.Now()
)

// Shader source code
const vertexShaderSource = `
#version 410 core

layout (location = 0) in vec3 aPos;
layout (location = 1) in vec3 aColor;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

out vec3 FragColor;

void main() {
    gl_Position = projection * view * model * vec4(aPos, 1.0);
    FragColor = aColor;
}
` + "\x00"

const fragmentShaderSource = `
#version 410 core

in vec3 FragColor;
out vec4 color;

void main() {
    color = vec4(FragColor, 1.0);
}
` + "\x00"

func init() {
	// This is needed to arrange that main() runs on main thread.
	// See documentation for functions that are only allowed to be called from the main thread.
	runtime.LockOSThread()
}

func main() {
	// Initialize GLFW
	if err := glfw.Init(); err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	// Configure GLFW
	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	// Create window
	window, err := glfw.CreateWindow(windowWidth, windowHeight, windowTitle, nil, nil)
	if err != nil {
		log.Fatalln("failed to create window:", err)
	}

	window.MakeContextCurrent()

	// Initialize OpenGL
	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize OpenGL:", err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL version:", version)

	// Set up viewport
	gl.Viewport(0, 0, windowWidth, windowHeight)

	// Enable depth testing
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)

	// Set clear color
	gl.ClearColor(0.2, 0.3, 0.3, 1.0)

	// Create shader program
	shaderProgram, err := renderer.NewProgram(vertexShaderSource, fragmentShaderSource)
	if err != nil {
		log.Fatalln("failed to create shader program:", err)
	}
	defer shaderProgram.Delete()

	// Create cube model
	cube := model.NewCube()

	// Create human model
	human := renderer.NewHuman(cube)

	// Create matrix stack
	matrixStack := matrix.NewMatrixStack()

	// Set up input callbacks
	setupInputCallbacks(window, human)

	// Main render loop
	for !window.ShouldClose() {
		// Calculate delta time
		currentTime := time.Now()
		deltaTime := float32(currentTime.Sub(lastTime).Seconds())
		lastTime = currentTime

		// Update animation
		updateAnimation(human, deltaTime)

		// Clear buffers
		gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

		// Use shader program
		shaderProgram.Use()

		// Set up projection matrix
		projection := createPerspectiveMatrix(45.0, float32(windowWidth)/float32(windowHeight), 0.1, 100.0)
		shaderProgram.SetMatrix4("projection", projection)

		// Set up view matrix
		view := createViewMatrix()
		shaderProgram.SetMatrix4("view", view)

		// Reset matrix stack
		matrixStack = matrix.NewMatrixStack()

		// Draw human
		human.Draw(shaderProgram, matrixStack)

		// Swap buffers and poll events
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func setupInputCallbacks(window *glfw.Window, human *renderer.Human) {
	// Keyboard callback
	window.SetKeyCallback(func(w *glfw.Window, key glfw.Key, scancode int, action glfw.Action, mods glfw.ModifierKey) {
		if action == glfw.Press {
			switch key {
			case glfw.KeyEscape:
				w.SetShouldClose(true)
			case glfw.Key1:
				human.AnimationMode = 0 // Idle
				fmt.Println("Animation: Idle")
			case glfw.Key2:
				human.AnimationMode = 1 // Walk
				fmt.Println("Animation: Walk")
			case glfw.Key3:
				human.AnimationMode = 2 // Jump
				fmt.Println("Animation: Jump")
			}
		}
	})

	// Mouse movement callback for camera control
	window.SetCursorPosCallback(func(w *glfw.Window, xpos, ypos float64) {
		if firstMouse {
			lastMouseX = xpos
			lastMouseY = ypos
			firstMouse = false
		}

		xoffset := xpos - lastMouseX
		yoffset := lastMouseY - ypos // Reversed since y-coordinates go from bottom to top

		lastMouseX = xpos
		lastMouseY = ypos

		sensitivity := float32(0.005)
		cameraAngleY += float32(xoffset) * sensitivity
		cameraAngleX += float32(yoffset) * sensitivity

		// Constrain pitch
		if cameraAngleX > 1.5 {
			cameraAngleX = 1.5
		}
		if cameraAngleX < -1.5 {
			cameraAngleX = -1.5
		}
	})

	// Mouse scroll callback for zoom
	window.SetScrollCallback(func(w *glfw.Window, xoff, yoff float64) {
		cameraDistance -= float32(yoff) * 0.5
		if cameraDistance < 1.0 {
			cameraDistance = 1.0
		}
		if cameraDistance > 10.0 {
			cameraDistance = 10.0
		}
	})

	// Print controls
	fmt.Println("Controls:")
	fmt.Println("  1 - Idle animation")
	fmt.Println("  2 - Walk animation")
	fmt.Println("  3 - Jump animation")
	fmt.Println("  Mouse - Rotate camera")
	fmt.Println("  Scroll - Zoom in/out")
	fmt.Println("  ESC - Exit")
}

func updateAnimation(human *renderer.Human, deltaTime float32) {
	switch human.AnimationMode {
	case 1: // Walk
		human.WalkCycle += deltaTime * 3.0 // Speed of walk animation
	case 2: // Jump
		human.JumpHeight += deltaTime * 4.0 // Speed of jump animation
	}
}

func createPerspectiveMatrix(fovy, aspect, near, far float32) [16]float32 {
	fH := float32(math.Tan(float64(fovy/360.0*math.Pi))) * near
	fW := fH * aspect

	return [16]float32{
		near / fW, 0, 0, 0,
		0, near / fH, 0, 0,
		0, 0, -(far + near) / (far - near), -1,
		0, 0, -(2 * far * near) / (far - near), 0,
	}
}

func createViewMatrix() [16]float32 {
	// Calculate camera position based on spherical coordinates
	x := cameraDistance * float32(math.Cos(float64(cameraAngleX))) * float32(math.Cos(float64(cameraAngleY)))
	y := cameraDistance * float32(math.Sin(float64(cameraAngleX)))
	z := cameraDistance * float32(math.Cos(float64(cameraAngleX))) * float32(math.Sin(float64(cameraAngleY)))

	// Look at origin
	return lookAt([3]float32{x, y, z}, [3]float32{0, 0, 0}, [3]float32{0, 1, 0})
}

func lookAt(eye, center, up [3]float32) [16]float32 {
	// Calculate forward vector
	forward := [3]float32{
		center[0] - eye[0],
		center[1] - eye[1],
		center[2] - eye[2],
	}
	forward = normalize(forward)

	// Calculate right vector
	right := cross(forward, up)
	right = normalize(right)

	// Calculate up vector
	up = cross(right, forward)

	// Negate forward for right-handed coordinate system
	forward[0] = -forward[0]
	forward[1] = -forward[1]
	forward[2] = -forward[2]

	// Create view matrix
	return [16]float32{
		right[0], up[0], forward[0], 0,
		right[1], up[1], forward[1], 0,
		right[2], up[2], forward[2], 0,
		-dot(right, eye), -dot(up, eye), -dot(forward, eye), 1,
	}
}

func normalize(v [3]float32) [3]float32 {
	length := float32(math.Sqrt(float64(v[0]*v[0] + v[1]*v[1] + v[2]*v[2])))
	return [3]float32{v[0] / length, v[1] / length, v[2] / length}
}

func cross(a, b [3]float32) [3]float32 {
	return [3]float32{
		a[1]*b[2] - a[2]*b[1],
		a[2]*b[0] - a[0]*b[2],
		a[0]*b[1] - a[1]*b[0],
	}
}

func dot(a, b [3]float32) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}