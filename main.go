package main

import (
    "log"
    "runtime"

    "github.com/go-gl/gl/v4.1-core/gl"
    "github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
    // GLFW must be called on the main OS thread
    runtime.LockOSThread()
}

func main() {
    if err := glfw.Init(); err != nil {
        log.Fatalln("failed to initialize glfw:", err)
    }
    defer glfw.Terminate()

    // Request OpenGL 4.1 core profile
    glfw.WindowHint(glfw.ContextVersionMajor, 4)
    glfw.WindowHint(glfw.ContextVersionMinor, 1)
    glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)

    win, err := glfw.CreateWindow(800, 600, "HumanGL test", nil, nil)
    if err != nil {
        log.Fatalln("failed to create window:", err)
    }
    win.MakeContextCurrent()

    if err := gl.Init(); err != nil {
        log.Fatalln("failed to initialize gl:", err)
    }

    version := gl.GoStr(gl.GetString(gl.VERSION))
    log.Println("OpenGL version", version)

    for !win.ShouldClose() {
        gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
        win.SwapBuffers()
        glfw.PollEvents()
    }
}
