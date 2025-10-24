package renderer

import "github.com/go-gl/gl/v3.3-core/gl"

type Cube struct {
	VAO uint32
	VBO uint32
	EBO uint32
}

func NewCube() *Cube {
	cube := &Cube{}

	vertices :=  []float32{
		// Front
		-0.5, -0.5, 0.5, // 0
		0.5, -0.5, 0.5,  // 1
		0.5, 0.5, 0.5,   // 2
		-0.5, 0.5, 0.5,  // 3

		// Back
		-0.5, -0.5, -0.5, // 4
		0.5, -0.5, -0.5,  // 5
		0.5, 0.5, -0.5,   // 6
		-0.5, 0.5, -0.5,  // 7
	}

	indices := []uint32{
		// Front
		0, 1, 2,
		2, 3, 0,

		// Back
		4, 5, 6,
		6, 7, 4,

		// Left
		0, 3, 7,
		7, 4, 0,

		// Right
		1, 5, 6,
		6, 2, 1,

		// Top
		3, 2, 6,
		6, 7, 3,

		// Bottom
		0, 4, 5,
		5, 1, 0,
	}

	gl.GenVertexArrays(1, &cube.VAO)
	gl.GenBuffers(1, &cube.VBO)
	gl.GenBuffers(1, &cube.EBO)

	gl.BindVertexArray(cube.VAO)

	// VBO
	gl.BindBuffer(gl.ARRAY_BUFFER, cube.VBO)
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(vertices), gl.Ptr(vertices), gl.STATIC_DRAW)

	// EBO
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, cube.EBO)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indices), gl.Ptr(indices), gl.STATIC_DRAW)

	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 3*4, gl.PtrOffset(0))
	gl.EnableVertexAttribArray(0)

	gl.BindVertexArray(0)

	return cube
}

func (c *Cube) Draw() {
	gl.BindVertexArray(c.VAO)
	gl.DrawElements(gl.TRIANGLES, 36, gl.UNSIGNED_INT, nil)
	gl.BindVertexArray(0)
}

func (c *Cube) Delete() {
	gl.DeleteVertexArrays(1, &c.VAO)
	gl.DeleteBuffers(1, &c.VBO)
	gl.DeleteBuffers(1, &c.EBO)
}
