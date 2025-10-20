//internal/renderer/renderer.go
package renderer

const VertexShaderSource = `#version 330 core
layout (location = 0) in vec3 aPos;
uniform mat4 uModel;
uniform mat4 uView;
uniform mat4 uProjection;
void main() {
	gl_Position = uProjection * uView * uModel * vec4(aPos, 1.0);
}` + "\x00"

const FragmentShaderSource = `#version 330 core
out vec4 FragColor;
uniform vec3 uColor;
void main() {
	FragColor = vec4(uColor, 1.0);
}` + "\x00"

