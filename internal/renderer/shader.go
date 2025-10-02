package renderer

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

// Program represents a shader program
type Program struct {
	ID uint32
}

// Compile compiles shader from source
func Compile(source string, shaderType uint32) (uint32, error) {
	shader := gl.CreateShader(shaderType)

	csources, free := gl.Strs(source + "\x00")
	gl.ShaderSource(shader, 1, csources, nil)
	free()
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(log))

		return 0, fmt.Errorf("failed to compile %v: %v", source, log)
	}

	return shader, nil
}

// NewProgram creates and links a shader program from vertex and fragment shaders
func NewProgram(vertexShaderSource, fragmentShaderSource string) (*Program, error) {
	vertexShader, err := Compile(vertexShaderSource, gl.VERTEX_SHADER)
	if err != nil {
		return nil, fmt.Errorf("failed to compile vertex shader: %v", err)
	}

	fragmentShader, err := Compile(fragmentShaderSource, gl.FRAGMENT_SHADER)
	if err != nil {
		return nil, fmt.Errorf("failed to compile fragment shader: %v", err)
	}

	program := gl.CreateProgram()

	gl.AttachShader(program, vertexShader)
	gl.AttachShader(program, fragmentShader)
	gl.LinkProgram(program)

	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)
		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		return nil, fmt.Errorf("failed to link program: %v", log)
	}

	gl.DeleteShader(vertexShader)
	gl.DeleteShader(fragmentShader)

	return &Program{ID: program}, nil
}

// Use activates the shader program
func (p *Program) Use() {
	gl.UseProgram(p.ID)
}

// Delete deletes the shader program
func (p *Program) Delete() {
	gl.DeleteProgram(p.ID)
}

// SetMatrix4 sets a mat4 uniform
func (p *Program) SetMatrix4(name string, value [16]float32) {
	location := gl.GetUniformLocation(p.ID, gl.Str(name+"\x00"))
	gl.UniformMatrix4fv(location, 1, false, &value[0])
}
