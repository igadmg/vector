package vector4

import (
	"encoding/json"
	"image/color"
	"math"

	"github.com/EliCDavis/vector"
)

// Vector contains 4 components
type Vector[T vector.Number] struct {
	x T
	y T
	z T
	w T
}

type (
	Float64 = Vector[float64]
	Float32 = Vector[float32]
	Int     = Vector[int]
	Int64   = Vector[int64]
)

// New creates a new vector with corresponding 3 components
func New[T vector.Number](x, y, z, w T) Vector[T] {
	return Vector[T]{
		x: x,
		y: y,
		z: z,
		w: w,
	}
}

// Fill creates a vector where each component is equal to v
func Fill[T vector.Number](v T) Vector[T] {
	return Vector[T]{
		x: v,
		y: v,
		z: v,
		w: v,
	}
}

func FromColor(c color.Color) Float64 {
	r, g, b, a := c.RGBA()
	return New(float64(r)/0xffff, float64(g)/0xffff, float64(b)/0xffff, float64(a)/0xffff)
}

// Zero is (0, 0, 0)
func Zero[T vector.Number]() Vector[T] {
	return New[T](0, 0, 0, 0)
}

// One is (1, 1, 1)
func One[T vector.Number]() Vector[T] {
	return New[T](1, 1, 1, 1)
}

// Average sums all vector4's components together and divides each
// component by the number of vectors added
func Average[T vector.Number](vectors []Vector[T]) Vector[T] {
	var center Vector[T]
	for _, v := range vectors {
		center = center.Add(v)
	}
	return center.DivByConstant(float64(len(vectors)))
}

// Lerp linearly interpolates between a and b by t
func Lerp[T vector.Number](a, b Vector[T], t float64) Vector[T] {
	return b.Sub(a).Scale(t).Add(a)
}

func (v Vector[T]) Scale(t float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) * t),
		y: T(float64(v.y) * t),
		z: T(float64(v.z) * t),
		w: T(float64(v.w) * t),
	}
}

func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return Vector[T]{
		x: T(float64(v.x) / t),
		y: T(float64(v.y) / t),
		z: T(float64(v.z) / t),
		w: T(float64(v.w) / t),
	}
}

func Min[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		T(math.Min(float64(a.x), float64(b.x))),
		T(math.Min(float64(a.y), float64(b.y))),
		T(math.Min(float64(a.z), float64(b.z))),
		T(math.Min(float64(a.w), float64(b.w))),
	)
}

func Max[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		T(math.Max(float64(a.x), float64(b.x))),
		T(math.Max(float64(a.y), float64(b.y))),
		T(math.Max(float64(a.z), float64(b.z))),
		T(math.Max(float64(a.w), float64(b.w))),
	)
}

func Midpoint[T vector.Number](a, b Vector[T]) Vector[T] {
	// center = (b - a)0.5 + a
	// center = b0.5 - a0.5 + a
	// center = b0.5 + a0.5
	// center = 0.5(b + a)
	return b.Add(a).Scale(0.5)
}

// Builds a vector from the data found from the passed in array to the best of
// it's ability. If the length of the array is smaller than the vector itself,
// only those values will be used to build the vector, and the remaining vector
// components will remain the default value of the vector's data type (some
// version of 0).
func FromArray[T vector.Number](data []T) Vector[T] {
	v := Vector[T]{}

	if len(data) > 0 {
		v.x = data[0]
	}

	if len(data) > 1 {
		v.y = data[1]
	}

	if len(data) > 2 {
		v.z = data[2]
	}

	if len(data) > 3 {
		v.w = data[3]
	}

	return v
}

func (v Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
		W float64 `json:"w"`
	}{
		X: float64(v.x),
		Y: float64(v.y),
		Z: float64(v.z),
		W: float64(v.w),
	})
}

func (v *Vector[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
		W float64 `json:"w"`
	}{
		X: 0,
		Y: 0,
		Z: 0,
		W: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.x = T(aux.X)
	v.y = T(aux.Y)
	v.z = T(aux.Z)
	v.w = T(aux.W)
	return nil
}

func (v Vector[T]) ToInt() Vector[int] {
	return Vector[int]{
		x: int(v.x),
		y: int(v.y),
		z: int(v.z),
		w: int(v.w),
	}
}

func (v Vector[T]) ToFloat64() Vector[float64] {
	return Vector[float64]{
		x: float64(v.x),
		y: float64(v.y),
		z: float64(v.z),
		w: float64(v.w),
	}
}

func (v Vector[T]) ToFloat32() Vector[float32] {
	return Vector[float32]{
		x: float32(v.x),
		y: float32(v.y),
		z: float32(v.z),
		w: float32(v.w),
	}
}

func (v Vector[T]) ToInt64() Vector[int64] {
	return Vector[int64]{
		x: int64(v.x),
		y: int64(v.y),
		z: int64(v.z),
		w: int64(v.w),
	}
}

// X returns the x component
func (v Vector[T]) X() T {
	return v.x
}

// SetX changes the x component of the vector
func (v Vector[T]) SetX(newX T) Vector[T] {
	return Vector[T]{
		x: newX,
		y: v.y,
		z: v.z,
		w: v.w,
	}
}

// Y returns the y component
func (v Vector[T]) Y() T {
	return v.y
}

// SetY changes the y component of the vector
func (v Vector[T]) SetY(newY T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: newY,
		z: v.z,
		w: v.w,
	}
}

// Z returns the z component
func (v Vector[T]) Z() T {
	return v.z
}

// SetZ changes the z component of the vector
func (v Vector[T]) SetZ(newZ T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y,
		z: newZ,
		w: v.w,
	}
}

// W returns the w component
func (v Vector[T]) W() T {
	return v.w
}

// SetW changes the w component of the vector
func (v Vector[T]) SetW(newW T) Vector[T] {
	return Vector[T]{
		x: v.x,
		y: v.y,
		z: v.z,
		w: newW,
	}
}

// Add takes each component of our vector and adds them to the vector passed
// in, returning a resulting vector
func (v Vector[T]) Add(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x + other.x,
		y: v.y + other.y,
		z: v.z + other.z,
		w: v.w + other.w,
	}
}

func (v Vector[T]) Sub(other Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x - other.x,
		y: v.y - other.y,
		z: v.z - other.z,
		w: v.w - other.w,
	}
}

func (v Vector[T]) Dot(other Vector[T]) float64 {
	return float64((v.x * other.x) + (v.y * other.y) + (v.z * other.z) + (v.w * other.w))
}

// Sqrt applies the math.Sqrt to each component of the vector
func (v Vector[T]) Sqrt() Vector[T] {
	return New(
		T(math.Sqrt(float64(v.x))),
		T(math.Sqrt(float64(v.y))),
		T(math.Sqrt(float64(v.z))),
		T(math.Sqrt(float64(v.w))),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Vector[T]) Abs() Vector[T] {
	return New(
		T(math.Abs(float64(v.x))),
		T(math.Abs(float64(v.y))),
		T(math.Abs(float64(v.z))),
		T(math.Abs(float64(v.w))),
	)
}

func (v Vector[T]) Clamp(min, max T) Vector[T] {
	return Vector[T]{
		x: T(math.Max(math.Min(float64(v.x), float64(max)), float64(min))),
		y: T(math.Max(math.Min(float64(v.y), float64(max)), float64(min))),
		z: T(math.Max(math.Min(float64(v.z), float64(max)), float64(min))),
		w: T(math.Max(math.Min(float64(v.w), float64(max)), float64(min))),
	}
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector[T]) Round() Vector[T] {
	return New(
		T(math.Round(float64(v.x))),
		T(math.Round(float64(v.y))),
		T(math.Round(float64(v.z))),
		T(math.Round(float64(v.w))),
	)
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Vector[T]) RoundToInt() Vector[int] {
	return New(
		int(math.Round(float64(v.x))),
		int(math.Round(float64(v.y))),
		int(math.Round(float64(v.z))),
		int(math.Round(float64(v.w))),
	)
}

// Floor applies the floor math operation to each component of the vector
func (v Vector[T]) Floor() Vector[T] {
	return New(
		T(math.Floor(float64(v.x))),
		T(math.Floor(float64(v.y))),
		T(math.Floor(float64(v.z))),
		T(math.Floor(float64(v.w))),
	)
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) FloorToInt() Vector[int] {
	return New(
		int(math.Floor(float64(v.x))),
		int(math.Floor(float64(v.y))),
		int(math.Floor(float64(v.z))),
		int(math.Floor(float64(v.w))),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Vector[T]) Ceil() Vector[T] {
	return New(
		T(math.Ceil(float64(v.x))),
		T(math.Ceil(float64(v.y))),
		T(math.Ceil(float64(v.z))),
		T(math.Ceil(float64(v.w))),
	)
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) CeilToInt() Vector[int] {
	return New(
		int(math.Ceil(float64(v.x))),
		int(math.Ceil(float64(v.y))),
		int(math.Ceil(float64(v.z))),
		int(math.Ceil(float64(v.w))),
	)
}

// MultByVector is component wise multiplication, also known as Hadamard product.
func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		x: v.x * o.x,
		y: v.y * o.y,
		z: v.z * o.z,
		w: v.w * o.w,
	}
}
