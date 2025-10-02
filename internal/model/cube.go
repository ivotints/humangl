package model

import "github.com/go-gl/gl/v4.1-core/gl"

// Cube represents a cube VAO
type Cube struct {
	vao uint32
	vbo uint32
	ebo uint32
}

// NewCube creates a new cube with specified dimensions
func NewCube() *Cube {
	vertices := []float32{
		// Front face
		-0.5, -0.5, 0.5, 1.0, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 0.0, 0.0,

		// Back face
		-0.5, -0.5, -0.5, 0.0, 1.0, 0.0,
		0.5, -0.5, -0.5, 0.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 0.0, 1.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0,

		// Top face
		-0.5, 0.5, -0.5, 0.0, 0.0, 1.0,
		0.5, 0.5, -0.5, 0.0, 0.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 0.0, 1.0,

		// Bottom face
		-0.5, -0.5, -0.5, 1.0, 1.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 1.0, 0.0,
		-0.5, -0.5, 0.5, 1.0, 1.0, 0.0,

		// Right face
		0.5, -0.5, -0.5, 0.0, 1.0, 1.0,
		0.5, 0.5, -0.5, 0.0, 1.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 1.0, 1.0,

		// Left face
		-0.5, -0.5, -0.5, 1.0, 0.0, 1.0,
		-0.5, 0.5, -0.5, 1.0, 0.0, 1.0,
		-0.5, 0.5, 0.5, 1.0, 0.0, 1.0,
		-0.5, -0.5, 0.5, 1.0, 0.0, 1.0,
	}

	indices := []uint32{
		0, 1, 2, 2, 3, 0,
		4, 5, 6, 6, 7, 4,
		8, 9, 10, 10, 11, 8,
		12, 13, 14, 14, 15, 12,
		16, 17, 18, 18, 19, 16,
		20, 21, 22, 22, 23, 20,
	}

	var vao, vbo, ebo uint32
	gl.GenVertexArrays(1, &vao)
	gl.GenBuffers(1, &vbo)
	gl.GenBuffers(1, &ebo)

	gl.BindVertexArray(vao)

	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(indices)*4, gl.Ptr(indices), gl.STATIC_DRAW)

	// Position attribute
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)
	// Color attribute
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, 6*4, gl.PtrOffset(3*4))
	gl.EnableVertexAttribArray(1)

	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)

	return &Cube{vao: vao, vbo: vbo, ebo: ebo}
}

// Draw renders the cube
func (c *Cube) Draw() {
	gl.BindVertexArray(c.vao)
	gl.DrawElements(gl.TRIANGLES, 36, gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}
