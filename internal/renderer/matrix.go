package renderer

import "github.com/go-gl/mathgl/mgl32"

func CreateModelMatrix(position mgl32.Vec3, rotation mgl32.Vec3, scale mgl32.Vec3) mgl32.Mat4 {
	model := mgl32.Ident4()

	model = model.Mul4(mgl32.Translate3D(position.X(), position.Y(), position.Z()))

	model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(rotation.Z()), mgl32.Vec3{0, 0, 1}))
	model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(rotation.X()), mgl32.Vec3{1, 0, 0}))
	model = model.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(rotation.Y()), mgl32.Vec3{0, 1, 0}))

	model = model.Mul4(mgl32.Scale3D(scale.X(), scale.Y(), scale.Z()))

	return model
}

func CreateViewMatrix(cameraPos mgl32.Vec3) mgl32.Mat4 {
	return mgl32.LookAtV(
		cameraPos,
		mgl32.Vec3{0,0,0},
		mgl32.Vec3{0,1,0},
	)
}

func CreateProjectionMatrix(aspectRatio float32) mgl32.Mat4 {
	return mgl32.Perspective(
		mgl32.DegToRad(45.0),
		aspectRatio,
		0.1,
		100.0,
	)
	}
