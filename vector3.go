package vector

import (
	"math"
)

// Vector3 contains 3 components
type Vector3 struct {
	x float64
	y float64
	z float64
}

// NewVector3 creates a new vector with corresponding 3 components
func NewVector3(x float64, y float64, z float64) Vector3 {
	return Vector3{
		x: x,
		y: y,
		z: z,
	}
}

// Vector3Right is (1, 0, 0)
func Vector3Right() Vector3 {
	return NewVector3(1, 0, 0)
}

// Vector3Zero is (0, 0, 0)
func Vector3Zero() Vector3 {
	return NewVector3(0, 0, 0)
}

// Vector3One is (1, 1, 1)
func Vector3One() Vector3 {
	return NewVector3(1, 1, 1)
}

// Vector3Up is (0, 1, 0)
func Vector3Up() Vector3 {
	return NewVector3(0, 1, 0)
}

// AverageVector3 sums all vector3's components together and divides each
// component by the number of vectors added
func AverageVector3(vectors []Vector3) Vector3 {
	var center Vector3
	for _, v := range vectors {
		center = center.Add(v)
	}
	return center.DivByConstant(float64(len(vectors)))
}

// X returns the x component
func (v Vector3) X() float64 {
	return v.x
}

// Y returns the y component
func (v Vector3) Y() float64 {
	return v.y
}

// Z returns the z component
func (v Vector3) Z() float64 {
	return v.z
}

// Perpendicular finds a vector that meets this vector at a right angle.
// https://stackoverflow.com/a/11132720/4974261
func (v Vector3) Perpendicular() Vector3 {
	var c Vector3
	if v.Y() != 0 || v.Z() != 0 {
		c = Vector3Right()
	} else {
		c = Vector3Up()
	}
	return v.Cross(c)
}

// Add takes each component of our vector and adds them to the vector passed
// in, returning a resulting vector
func (v Vector3) Add(other Vector3) Vector3 {
	return Vector3{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
	}
}

func (v Vector3) Sub(other Vector3) Vector3 {
	return Vector3{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
	}
}

func (v Vector3) Dot(other Vector3) float64 {
	return (v.x * other.x) + (v.y * other.y) + (v.z * other.z)
}

func (v Vector3) Cross(other Vector3) Vector3 {
	return Vector3{
		x: (v.y * other.z) - (v.z * other.y),
		y: (v.z * other.x) - (v.x * other.z),
		z: (v.x * other.y) - (v.y * other.x),
	}
}

func (v Vector3) Normalized() Vector3 {
	return v.DivByConstant(v.Length())
}

func (v Vector3) MultByConstant(t float64) Vector3 {
	return Vector3{
		x: v.x * t,
		y: v.y * t,
		z: v.z * t,
	}
}

func (v Vector3) MultByVector(o Vector3) Vector3 {
	return Vector3{
		x: v.x * o.x,
		y: v.y * o.y,
		z: v.z * o.z,
	}
}

func (v Vector3) DivByConstant(t float64) Vector3 {
	return v.MultByConstant(1.0 / t)
}

func (v Vector3) Length() float64 {
	return math.Sqrt((v.x * v.x) + (v.y * v.y) + (v.z * v.z))
}

func (v Vector3) SquaredLength() float64 {
	return (v.x * v.x) + (v.y * v.y) + (v.z * v.z)
}

func (v Vector3) Distance(other Vector3) float64 {
	return math.Sqrt(math.Pow(other.x-v.x, 2.0) + math.Pow(other.y-v.y, 2.0) + math.Pow(other.z-v.z, 2.0))
}
