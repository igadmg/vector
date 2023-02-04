package vector2

import (
	"math"
	"math/rand"

	"github.com/EliCDavis/vector"
)

type Vector[T vector.Number] struct {
	x T
	y T
}

type (
	Float64 = Vector[float64]
	Float32 = Vector[float32]
	Int     = Vector[int]
	Int64   = Vector[int64]
)

func New[T vector.Number](x T, y T) Vector[T] {
	return Vector[T]{
		x: x,
		y: y,
	}
}

func Zero[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: 0,
		y: 0,
	}
}

func Up[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: 0,
		y: 1,
	}
}

func Down[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: 0,
		y: -1,
	}
}

func Left[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: -1,
		y: 0,
	}
}

func Right[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: 1,
		y: 0,
	}
}

func One[T vector.Number]() Vector[T] {
	return Vector[T]{
		x: 1,
		y: 1,
	}
}

// Lerp linearly interpolates between a and b by t
func Lerp[T vector.Number](a, b Vector[T], t float64) Vector[T] {
	return b.Sub(a).Scale(t).Add(a)
}

func Min[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		T(math.Min(float64(a.x), float64(b.x))),
		T(math.Min(float64(a.y), float64(b.y))),
	)
}

func Max[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		T(math.Max(float64(a.x), float64(b.x))),
		T(math.Max(float64(a.y), float64(b.y))),
	)
}

func Rand() Vector[float64] {
	return Vector[float64]{
		x: rand.Float64(),
		y: rand.Float64(),
	}
}

// Sqrt applies the math.Sqrt to each component of the vector
func (v Vector[T]) Sqrt() Vector[T] {
	return New(
		T(math.Sqrt(float64(v.x))),
		T(math.Sqrt(float64(v.y))),
	)
}

func (v Vector[T]) ToInt() Vector[int] {
	return Vector[int]{
		x: int(v.x),
		y: int(v.y),
	}
}

func (v Vector[T]) ToFloat64() Vector[float64] {
	return Vector[float64]{
		x: float64(v.x),
		y: float64(v.y),
	}
}

func (v Vector[T]) Clamp(min, max T) Vector[T] {
	return Vector[T]{
		x: T(math.Max(math.Min(float64(v.x), float64(max)), float64(min))),
		y: T(math.Max(math.Min(float64(v.y), float64(max)), float64(min))),
	}
}

func (v Vector[T]) ToFloat32() Vector[float32] {
	return Vector[float32]{
		x: float32(v.x),
		y: float32(v.y),
	}
}

func (v Vector[T]) ToInt64() Vector[int64] {
	return Vector[int64]{
		x: int64(v.x),
		y: int64(v.y),
	}
}

func (v Vector[T]) X() T {
	return v.x
}

// SetX changes the x component of the vector
func (v Vector[T]) SetX(newX T) Vector[T] {
	return Vector[T]{
		x: newX,
		y: v.y,
	}
}

func (v Vector[T]) Y() T {
	return v.y
}

// SetY changes the y component of the vector
func (v Vector[T]) SetY(newY T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: newY,
	}
}

// Midpoint returns the midpoint between this vector and the vector passed in.
func (v Vector[T]) Midpoint(o Vector[T]) Vector[T] {
	return o.Add(v).Scale(0.5)
}

func (v Vector[T]) Dot(other Vector[T]) float64 {
	return float64(v.x*other.x) + float64(v.y*other.y)
}

// Perpendicular creates a vector perpendicular to the one passed in with the
// same magnitude
func (v Vector[T]) Perpendicular() Vector[T] {
	return Vector[T]{
		x: v.y,
		y: -v.x,
	}
}

// Add returns a vector that is the result of two vectors added together
func (v Vector[T]) Add(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x + other.x,
		y: v.y + other.y,
	}
}

func (v Vector[T]) Sub(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x - other.x,
		y: v.y - other.y,
	}
}

func (v Vector[T]) Length() float64 {
	return math.Sqrt(float64(v.x*v.x) + float64(v.y*v.y))
}

func (v Vector[T]) LengthSquared() float64 {
	return float64(v.x*v.x) + float64(v.y*v.y)
}

func (v Vector[T]) Normalized() Vector[T] {
	return v.DivByConstant(v.Length())
}

func (v Vector[T]) Scale(t float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) * t),
		y: T(float64(v.y) * t),
	}
}

func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x * o.x,
		y: v.y * o.y,
	}
}

func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return v.Scale(1.0 / t)
}

func (v Vector[T]) DistanceSquared(other Vector[T]) float64 {
	xDist := other.x - v.x
	yDist := other.y - v.y
	return float64((xDist * xDist) + (yDist * yDist))
}

// Distance is the euclidean distance between two points
func (v Vector[T]) Distance(other Vector[T]) float64 {
	return math.Sqrt(v.DistanceSquared(other))
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector[T]) Round() Vector[T] {
	return Vector[T]{
		x: T(math.Round(float64(v.x))),
		y: T(math.Round(float64(v.y))),
	}
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Vector[T]) RoundToInt() Vector[int] {
	return New(
		int(math.Round(float64(v.x))),
		int(math.Round(float64(v.y))),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Vector[T]) Ceil() Vector[T] {
	return Vector[T]{
		x: T(math.Ceil(float64(v.x))),
		y: T(math.Ceil(float64(v.y))),
	}
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) CeilToInt() Vector[int] {
	return New(
		int(math.Ceil(float64(v.x))),
		int(math.Ceil(float64(v.y))),
	)
}

func (v Vector[T]) Floor() Vector[T] {
	return Vector[T]{
		x: T(math.Floor(float64(v.x))),
		y: T(math.Floor(float64(v.y))),
	}
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) FloorToInt() Vector[int] {
	return New(
		int(math.Floor(float64(v.x))),
		int(math.Floor(float64(v.y))),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Vector[T]) Abs() Vector[T] {
	return Vector[T]{
		x: T(math.Abs(float64(v.x))),
		y: T(math.Abs(float64(v.y))),
	}
}

func (v Vector[T]) NearZero() bool {
	const s = 1e-8
	return (math.Abs(float64(v.X())) < s) && (math.Abs(float64(v.Y())) < s)
}
