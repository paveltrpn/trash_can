package main

type vec2_t [2]float32
type vec3_t [3]float32

type vec_t interface {
	dot(a, b vec_t) float32
}

func (vec2_t) dot(a, b vec2_t) float32 {
	return a[0]*b[0] + a[1]*b[1]
}

func (vec3_t) dot(a, b vec3_t) float32 {
	return a[0]*b[0] + a[1]*b[1] + a[2]*b[2]
}

func vec_dot(a, b vec_t) float32 {
	return a.dot(a, b)
}

func main() {

	var (
		a      vec_t
		v1, v2 vec2_t
	)

	b := vec_dot(v1, v2)
}
