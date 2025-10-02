package matrix

// MatrixStack manages a stack of matrices
type MatrixStack struct {
	stack []Matrix4
}

// NewMatrixStack creates a new matrix stack with identity matrix
func NewMatrixStack() *MatrixStack {
	return &MatrixStack{
		stack: []Matrix4{NewIdentity()},
	}
}

// Current returns the top matrix in the stack
func (s *MatrixStack) Current() Matrix4 {
	return s.stack[len(s.stack)-1]
}

// Push copies the top matrix and pushes it onto the stack
func (s *MatrixStack) Push() {
	s.stack = append(s.stack, s.Current())
}

// Pop removes the top matrix from the stack
func (s *MatrixStack) Pop() {
	if len(s.stack) > 1 {
		s.stack = s.stack[:len(s.stack)-1]
	}
}

// Apply applies a transformation to the top matrix
func (s *MatrixStack) Apply(m Matrix4) {
	current := s.Current()
	s.stack[len(s.stack)-1] = current.Multiply(m)
}

// Translate applies a translation to the top matrix
func (s *MatrixStack) Translate(x, y, z float32) {
	s.Apply(Translate(x, y, z))
}

// Scale applies a scaling to the top matrix
func (s *MatrixStack) Scale(x, y, z float32) {
	s.Apply(Scale(x, y, z))
}

// RotateX applies a rotation around X axis to the top matrix
func (s *MatrixStack) RotateX(angle float32) {
	s.Apply(RotateX(angle))
}

// RotateY applies a rotation around Y axis to the top matrix
func (s *MatrixStack) RotateY(angle float32) {
	s.Apply(RotateY(angle))
}

// RotateZ applies a rotation around Z axis to the top matrix
func (s *MatrixStack) RotateZ(angle float32) {
	s.Apply(RotateZ(angle))
}
