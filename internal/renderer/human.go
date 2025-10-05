package renderer

import (
	"math"

	"github.com/ivotints/humangl/internal/matrix"
	"github.com/ivotints/humangl/internal/model"
)

// BodyPartSizes stores the size factors for each body part
type BodyPartSizes struct {
	HeadSize     float32
	TorsoWidth   float32
	TorsoHeight  float32 
	TorsoDepth   float32
	UpperArmSize float32
	ForearmSize  float32
	ThighSize    float32
	LowerLegSize float32
}

// DefaultSizes returns default body part sizes
func DefaultSizes() BodyPartSizes {
	return BodyPartSizes{
		HeadSize:     0.5,
		TorsoWidth:   1.0,
		TorsoHeight:  1.5,
		TorsoDepth:   0.5,
		UpperArmSize: 0.7,
		ForearmSize:  0.6,
		ThighSize:    0.8,
		LowerLegSize: 0.7,
	}
}

// Human represents a human model with hierarchical parts
type Human struct {
	Cube  *model.Cube
	Sizes BodyPartSizes
	// Animation state
	WalkCycle     float32
	JumpHeight    float32
	AnimationMode int // 0 = idle, 1 = walk, 2 = jump
}

// NewHuman creates a new human model
func NewHuman(cube *model.Cube) *Human {
	return &Human{
		Cube:          cube,
		Sizes:         DefaultSizes(),
		WalkCycle:     0,
		JumpHeight:    0,
		AnimationMode: 0,
	}
}

// Draw renders the human model using the shader and matrix stack
func (h *Human) Draw(shader *Program, stack *matrix.MatrixStack) {
	// Apply model transformation based on animation state
	stack.Push()
	defer stack.Pop()

	// Apply root transformation
	if h.AnimationMode == 2 {
		// Jump animation - move the entire model up and down
		jumpOffset := float32(0.5 * (1.0 + math.Sin(float64(h.JumpHeight))))
		stack.Translate(0, jumpOffset, 0)
	} else if h.AnimationMode == 1 {
		// Walk animation - slight up/down motion
		walkBounce := float32(0.05 * (1.0 + math.Sin(2*float64(h.WalkCycle))))
		stack.Translate(0, walkBounce, 0)
	}

	// Draw torso
	stack.Push()
	stack.Scale(h.Sizes.TorsoWidth, h.Sizes.TorsoHeight, h.Sizes.TorsoDepth)
	shader.SetMatrix4("model", stack.Current())
	h.Cube.Draw()
	stack.Pop()

	// Draw head
	stack.Push()
	stack.Translate(0, h.Sizes.TorsoHeight/2+h.Sizes.HeadSize/2, 0)
	stack.Scale(h.Sizes.HeadSize, h.Sizes.HeadSize, h.Sizes.HeadSize)
	shader.SetMatrix4("model", stack.Current())
	h.Cube.Draw()
	stack.Pop()

	// Draw right upper arm
	stack.Push()
	stack.Translate(-(h.Sizes.TorsoWidth/2 + h.Sizes.UpperArmSize/2), h.Sizes.TorsoHeight/4, 0)
	if h.AnimationMode == 1 {
		// Walking arm swing
		stack.RotateZ(float32(0.3 * math.Sin(float64(h.WalkCycle))))
	}
	stack.Scale(h.Sizes.UpperArmSize, h.Sizes.UpperArmSize/2, h.Sizes.UpperArmSize/2)
	shader.SetMatrix4("model", stack.Current())
	h.Cube.Draw()

	// Draw right forearm
	stack.Translate(0, -h.Sizes.UpperArmSize/2, 0)
	if h.AnimationMode == 1 {
		// Walking forearm movement
		stack.RotateZ(float32(0.3 * math.Sin(float64(h.WalkCycle))))
	}
	stack.Scale(1, h.Sizes.ForearmSize/h.Sizes.UpperArmSize, 1)
	shader.SetMatrix4("model", stack.Current())
	h.Cube.Draw()
	stack.Pop()

	// Draw left upper arm
	stack.Push()
	stack.Translate(h.Sizes.TorsoWidth/2+h.Sizes.UpperArmSize/2, h.Sizes.TorsoHeight/4, 0)
	if h.AnimationMode == 1 {
		// Walking arm swing - opposite phase
		stack.RotateZ(float32(-0.3 * math.Sin(float64(h.WalkCycle))))
	}
	stack.Scale(h.Sizes.UpperArmSize, h.Sizes.UpperArmSize/2, h.Sizes.UpperArmSize/2)
	shader.SetMatrix4("model", stack.Current())
	h.Cube.Draw()

	// Draw left forearm
	stack.Translate(0, -h.Sizes.UpperArmSize/2, 0)
	if h.AnimationMode == 1 {
		// Walking forearm movement - opposite phase
		stack.RotateZ(float32(-0.3 * math.Sin(float64(h.WalkCycle))))
	}
	stack.Scale(1, h.Sizes.ForearmSize/h.Sizes.UpperArmSize, 1)
	shader.SetMatrix4("model", stack.Current())
	h.Cube.Draw()
	stack.Pop()

	// Draw right thigh
	stack.Push()
	stack.Translate(-h.Sizes.TorsoWidth/4, -h.Sizes.TorsoHeight/2-h.Sizes.ThighSize/2, 0)
	if h.AnimationMode == 1 {
		// Walking leg swing
		stack.RotateZ(float32(-0.3 * math.Sin(float64(h.WalkCycle))))
	}
	stack.Scale(h.Sizes.ThighSize/2, h.Sizes.ThighSize, h.Sizes.ThighSize/2)
	shader.SetMatrix4("model", stack.Current())
	h.Cube.Draw()

	// Draw right lower leg
	stack.Translate(0, -h.Sizes.ThighSize/2-h.Sizes.LowerLegSize/2, 0)
	if h.AnimationMode == 1 {
		// Walking lower leg movement
		stack.RotateZ(float32(0.6 * math.Sin(float64(h.WalkCycle))))
	}
	stack.Scale(1, h.Sizes.LowerLegSize/h.Sizes.ThighSize, 1)
	shader.SetMatrix4("model", stack.Current())
	h.Cube.Draw()
	stack.Pop()

	// Draw left thigh
	stack.Push()
	stack.Translate(h.Sizes.TorsoWidth/4, -h.Sizes.TorsoHeight/2-h.Sizes.ThighSize/2, 0)
	if h.AnimationMode == 1 {
		// Walking leg swing - opposite phase to right leg
		stack.RotateZ(float32(0.3 * math.Sin(float64(h.WalkCycle))))
	}
	stack.Scale(h.Sizes.ThighSize/2, h.Sizes.ThighSize, h.Sizes.ThighSize/2)
	shader.SetMatrix4("model", stack.Current())
	h.Cube.Draw()

	// Draw left lower leg
	stack.Translate(0, -h.Sizes.ThighSize/2-h.Sizes.LowerLegSize/2, 0)
	if h.AnimationMode == 1 {
		// Walking lower leg movement - opposite phase to right leg
		stack.RotateZ(float32(-0.6 * math.Sin(float64(h.WalkCycle))))
	}
	stack.Scale(1, h.Sizes.LowerLegSize/h.Sizes.ThighSize, 1)
	shader.SetMatrix4("model", stack.Current())
	h.Cube.Draw()
	stack.Pop()
}
