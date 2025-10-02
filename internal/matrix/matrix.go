package matrix

import (
	"math"
)

// Matrix4 represents a 4x4 matrix
type Matrix4 [16]float32

// NewIdentity creates a new identity matrix
func NewIdentity() Matrix4 {
	return Matrix4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1,
	}
}

// Multiply multiplies this matrix with another matrix
func (m Matrix4) Multiply(other Matrix4) Matrix4 {
	var result Matrix4
	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {
			sum := float32(0)
			for i := 0; i < 4; i++ {
				sum += m[row*4+i] * other[i*4+col]
			}
			result[row*4+col] = sum
		}
	}
	return result
}

// Translate creates a translation matrix
func Translate(x, y, z float32) Matrix4 {
	m := NewIdentity()
	m[12] = x
	m[13] = y
	m[14] = z
	return m
}

// Scale creates a scaling matrix
func Scale(x, y, z float32) Matrix4 {
	m := NewIdentity()
	m[0] = x
	m[5] = y
	m[10] = z
	return m
}

// RotateX creates a rotation matrix around X axis
func RotateX(angle float32) Matrix4 {
	m := NewIdentity()
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	m[5] = c
	m[6] = s
	m[9] = -s
	m[10] = c
	return m
}

// RotateY creates a rotation matrix around Y axis
func RotateY(angle float32) Matrix4 {
	m := NewIdentity()
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	m[0] = c
	m[2] = -s
	m[8] = s
	m[10] = c
	return m
}

// RotateZ creates a rotation matrix around Z axis
func RotateZ(angle float32) Matrix4 {
	m := NewIdentity()
	c := float32(math.Cos(float64(angle)))
	s := float32(math.Sin(float64(angle)))
	m[0] = c
	m[1] = s
	m[4] = -s
	m[5] = c
	return m
}
