package vector3

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
	"math/rand"

	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/mathex"
	"github.com/EliCDavis/vector/vector2"
)

// Vector contains 3 components
type Vector[T vector.Number] struct {
	X T
	Y T
	Z T
}

type (
	Float64 = Vector[float64]
	Float32 = Vector[float32]
	Int     = Vector[int]
	Int64   = Vector[int64]
	Int32   = Vector[int32]
	Int16   = Vector[int16]
	Int8    = Vector[int8]
)

// New creates a new vector with corresponding 3 components
func New[T vector.Number](x T, y T, z T) Vector[T] {
	return Vector[T]{
		X: x,
		Y: y,
		Z: z,
	}
}

// Fill creates a vector where each component is equal to v
func Fill[T vector.Number](v T) Vector[T] {
	return Vector[T]{
		X: v,
		Y: v,
		Z: v,
	}
}

// Right is (1, 0, 0)
func Right[T vector.Number]() Vector[T] {
	return New[T](1, 0, 0)
}

// Left is (-1, 0, 0)
func Left[T vector.Number]() Vector[T] {
	return New[T](-1, 0, 0)
}

// Forward is (0, 0, 1)
func Forward[T vector.Number]() Vector[T] {
	return New[T](0, 0, 1)
}

// Backwards is (0, 0, -1)
func Backwards[T vector.Number]() Vector[T] {
	return New[T](0, 0, -1)
}

// Up is (0, 1, 0)
func Up[T vector.Number]() Vector[T] {
	return New[T](0, 1, 0)
}

// Down is (0, -1, 0)
func Down[T vector.Number]() Vector[T] {
	return New[T](0, -1, 0)
}

// Zero is (0, 0, 0)
func Zero[T vector.Number]() Vector[T] {
	return New[T](0, 0, 0)
}

// One is (1, 1, 1)
func One[T vector.Number]() Vector[T] {
	return New[T](1, 1, 1)
}

func FromColor(c color.Color) Float64 {
	r, g, b, _ := c.RGBA()
	return New(float64(r)/0xffff, float64(g)/0xffff, float64(b)/0xffff)
}

// Average sums all vector3's components together and divides each
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
	return Vector[T]{
		X: T((float64(b.X-a.X) * t) + float64(a.X)),
		Y: T((float64(b.Y-a.Y) * t) + float64(a.Y)),
		Z: T((float64(b.Z-a.Z) * t) + float64(a.Z)),
	}
}

func Min[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		T(math.Min(float64(a.X), float64(b.X))),
		T(math.Min(float64(a.Y), float64(b.Y))),
		T(math.Min(float64(a.Z), float64(b.Z))),
	)
}

func Max[T vector.Number](a, b Vector[T]) Vector[T] {
	return New(
		T(math.Max(float64(a.X), float64(b.X))),
		T(math.Max(float64(a.Y), float64(b.Y))),
		T(math.Max(float64(a.Z), float64(b.Z))),
	)
}

func MaxX[T vector.Number](a, b Vector[T]) T {
	return T(math.Max(float64(a.X), float64(b.X)))
}

func MaxY[T vector.Number](a, b Vector[T]) T {
	return T(math.Max(float64(a.Y), float64(b.Y)))
}

func MaxZ[T vector.Number](a, b Vector[T]) T {
	return T(math.Max(float64(a.Z), float64(b.Z)))
}

func MinX[T vector.Number](a, b Vector[T]) T {
	return T(math.Min(float64(a.X), float64(b.X)))
}

func MinY[T vector.Number](a, b Vector[T]) T {
	return T(math.Min(float64(a.Y), float64(b.Y)))
}

func MinZ[T vector.Number](a, b Vector[T]) T {
	return T(math.Min(float64(a.Z), float64(b.Z)))
}

func Midpoint[T vector.Number](a, b Vector[T]) Vector[T] {
	// center = (b - a)0.5 + a
	// center = b0.5 - a0.5 + a
	// center = b0.5 + a0.5
	// center = 0.5(b + a)
	return Vector[T]{
		X: T(float64(a.X+b.X) * 0.5),
		Y: T(float64(a.Y+b.Y) * 0.5),
		Z: T(float64(a.Z+b.Z) * 0.5),
	}
}

// Builds a vector from the data found from the passed in array to the best of
// it's ability. If the length of the array is smaller than the vector itself,
// only those values will be used to build the vector, and the remaining vector
// components will remain the default value of the vector's data type (some
// version of 0).
func FromArray[T vector.Number](data []T) Vector[T] {
	v := Vector[T]{}

	if len(data) > 0 {
		v.X = data[0]
	}

	if len(data) > 1 {
		v.Y = data[1]
	}

	if len(data) > 2 {
		v.Z = data[2]
	}

	return v
}

func (v Vector[T]) ToArr() []T {
	return []T{v.X, v.Y, v.Z}
}

func (v Vector[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}{
		X: float64(v.X),
		Y: float64(v.Y),
		Z: float64(v.Z),
	})
}

func (v *Vector[T]) UnmarshalJSON(data []byte) error {
	aux := &struct {
		X float64 `json:"x"`
		Y float64 `json:"y"`
		Z float64 `json:"z"`
	}{
		X: 0,
		Y: 0,
		Z: 0,
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	v.X = T(aux.X)
	v.Y = T(aux.Y)
	v.Z = T(aux.Z)
	return nil
}

func (v Vector[T]) ContainsNaN() bool {
	if math.IsNaN(float64(v.X)) {
		return true
	}

	if math.IsNaN(float64(v.Y)) {
		return true
	}

	if math.IsNaN(float64(v.Z)) {
		return true
	}

	return false
}

func (v Vector[T]) Format(format string) string {
	return fmt.Sprintf(format, v.X, v.Y, v.Z)
}

func (v Vector[T]) MinComponent() T {
	return T(math.Min(float64(v.X), math.Min(float64(v.Y), float64(v.Z))))
}

func (v Vector[T]) MaxComponent() T {
	return T(math.Max(float64(v.X), math.Max(float64(v.Y), float64(v.Z))))
}

func (v Vector[T]) ToInt() Vector[int] {
	return Vector[int]{
		X: int(v.X),
		Y: int(v.Y),
		Z: int(v.Z),
	}
}

func (v Vector[T]) ToFloat64() Vector[float64] {
	return Vector[float64]{
		X: float64(v.X),
		Y: float64(v.Y),
		Z: float64(v.Z),
	}
}

func (v Vector[T]) ToFloat32() Vector[float32] {
	return Vector[float32]{
		X: float32(v.X),
		Y: float32(v.Y),
		Z: float32(v.Z),
	}
}

func (v Vector[T]) ToInt64() Vector[int64] {
	return Vector[int64]{
		X: int64(v.X),
		Y: int64(v.Y),
		Z: int64(v.Z),
	}
}

/*
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
		}
	}
*/
func (v Vector[T]) XZY() Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Z,
		Z: v.Y,
	}
}

func (v Vector[T]) ZXY() Vector[T] {
	return Vector[T]{
		X: v.Z,
		Y: v.X,
		Z: v.Y,
	}
}

func (v Vector[T]) ZYX() Vector[T] {
	return Vector[T]{
		X: v.Z,
		Y: v.Y,
		Z: v.X,
	}
}

func (v Vector[T]) YXZ() Vector[T] {
	return Vector[T]{
		X: v.Y,
		Y: v.X,
		Z: v.Z,
	}
}

func (v Vector[T]) YZX() Vector[T] {
	return Vector[T]{
		X: v.Y,
		Y: v.Z,
		Z: v.X,
	}
}

// XY returns vector2 with the x and y components
func (v Vector[T]) XY() vector2.Vector[T] {
	return vector2.New(v.X, v.Y)
}

// XZ returns vector2 with the x and z components
func (v Vector[T]) XZ() vector2.Vector[T] {
	return vector2.New(v.X, v.Z)
}

// YZ returns vector2 with the y and z components
func (v Vector[T]) YZ() vector2.Vector[T] {
	return vector2.New(v.Y, v.Z)
}

// YX returns vector2 with the y and x components
func (v Vector[T]) YX() vector2.Vector[T] {
	return vector2.New(v.Y, v.X)
}

// ZX returns vector2 with the z and x components
func (v Vector[T]) ZX() vector2.Vector[T] {
	return vector2.New(v.Z, v.X)
}

// ZY returns vector2 with the z and y components
func (v Vector[T]) ZY() vector2.Vector[T] {
	return vector2.New(v.Z, v.Y)
}

// Midpoint returns the midpoint between this vector and the vector passed in.
func (v Vector[T]) Midpoint(o Vector[T]) Vector[T] {
	return Vector[T]{
		X: T(float64(o.X+v.X) * 0.5),
		Y: T(float64(o.Y+v.Y) * 0.5),
		Z: T(float64(o.Z+v.Z) * 0.5),
	}
}

// Perpendicular finds a vector that meets this vector at a right angle.
// https://stackoverflow.com/a/11132720/4974261
func (v Vector[T]) Perpendicular() Vector[T] {
	var c Vector[T]
	if v.Y != 0 || v.Z != 0 {
		c = Right[T]()
	} else {
		c = Up[T]()
	}
	return v.Cross(c)
}

// Round takes each component of the vector and rounds it to the nearest whole
// number
func (v Vector[T]) Round() Vector[T] {
	return New(
		mathex.Round(v.X),
		mathex.Round(v.Y),
		mathex.Round(v.Z),
	)
}

// RoundToInt takes each component of the vector and rounds it to the nearest
// whole number, and then casts it to a int
func (v Vector[T]) RoundToInt() Vector[int] {
	return New(
		int(mathex.Round(v.X)),
		int(mathex.Round(v.Y)),
		int(mathex.Round(v.Z)),
	)
}

// Floor applies the floor math operation to each component of the vector
func (v Vector[T]) Floor() Vector[T] {
	return New(
		mathex.Floor(v.X),
		mathex.Floor(v.Y),
		mathex.Floor(v.Z),
	)
}

// FloorToInt applies the floor math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) FloorToInt() Vector[int] {
	return New(
		int(mathex.Floor(v.X)),
		int(mathex.Floor(v.Y)),
		int(mathex.Floor(v.Z)),
	)
}

// Ceil applies the ceil math operation to each component of the vector
func (v Vector[T]) Ceil() Vector[T] {
	return New(
		mathex.Ceil(v.X),
		mathex.Ceil(v.Y),
		mathex.Ceil(v.Z),
	)
}

// CeilToInt applies the ceil math operation to each component of the vector,
// and then casts it to a int
func (v Vector[T]) CeilToInt() Vector[int] {
	return New(
		int(mathex.Ceil(v.X)),
		int(mathex.Ceil(v.Y)),
		int(mathex.Ceil(v.Z)),
	)
}

// Sqrt applies the math.Sqrt to each component of the vector
func (v Vector[T]) Sqrt() Vector[T] {
	return New(
		T(math.Sqrt(float64(v.X))),
		T(math.Sqrt(float64(v.Y))),
		T(math.Sqrt(float64(v.Z))),
	)
}

// Abs applies the Abs math operation to each component of the vector
func (v Vector[T]) Abs() Vector[T] {
	return New(
		T(math.Abs(float64(v.X))),
		T(math.Abs(float64(v.Y))),
		T(math.Abs(float64(v.Z))),
	)
}

func (v Vector[T]) Clamp(min, max T) Vector[T] {
	return Vector[T]{
		X: mathex.Clamp(v.X, min, max),
		Y: mathex.Clamp(v.Y, min, max),
		Z: mathex.Clamp(v.Z, min, max),
	}
}

// Add takes each component of our vector and adds them to the vector passed
// in, returning a resulting vector
func (v Vector[T]) Add(other Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X + other.X,
		Y: v.Y + other.Y,
		Z: v.Z + other.Z,
	}
}

func (v Vector[T]) Sub(other Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X - other.X,
		Y: v.Y - other.Y,
		Z: v.Z - other.Z,
	}
}

func (v Vector[T]) Dot(other Vector[T]) float64 {
	return float64((v.X * other.X) + (v.Y * other.Y) + (v.Z * other.Z))
}

func (v Vector[T]) Cross(other Vector[T]) Vector[T] {
	return Vector[T]{
		X: (v.Y * other.Z) - (v.Z * other.Y),
		Y: (v.Z * other.X) - (v.X * other.Z),
		Z: (v.X * other.Y) - (v.Y * other.X),
	}
}

func (v Vector[T]) Normalized() Vector[T] {
	return v.DivByConstant(v.Length())
}

// Rand returns a vector with each component being a random value between [0.0, 1.0)
func Rand(r *rand.Rand) Vector[float64] {
	return Vector[float64]{
		X: r.Float64(),
		Y: r.Float64(),
		Z: r.Float64(),
	}
}

// RandRange returns a vector where each component is a random value that falls
// within the values of min and max
func RandRange[T vector.Number](r *rand.Rand, min, max T) Vector[T] {
	dist := float64(max - min)
	return Vector[T]{
		X: T(r.Float64()*dist) + min,
		Y: T(r.Float64()*dist) + min,
		Z: T(r.Float64()*dist) + min,
	}
}

// RandInUnitSphere returns a randomly sampled point in or on the unit
func RandInUnitSphere(r *rand.Rand) Vector[float64] {
	for {
		p := RandRange(r, -1., 1.)
		if p.LengthSquared() < 1 {
			return p
		}
	}
}

// RandNormal returns a random normal
func RandNormal(r *rand.Rand) Vector[float64] {
	return Vector[float64]{
		X: -1. + (r.Float64() * 2.),
		Y: -1. + (r.Float64() * 2.),
		Z: -1. + (r.Float64() * 2.),
	}.Normalized()
}

func (v Vector[T]) Scale(t float64) Vector[T] {
	return Vector[T]{
		X: T(float64(v.X) * t),
		Y: T(float64(v.Y) * t),
		Z: T(float64(v.Z) * t),
	}
}

func (v Vector[T]) Reflect(normal Vector[T]) Vector[T] {
	return v.Sub(normal.Scale(2. * v.Dot(normal)))
}

func (v Vector[T]) Refract(normal Vector[T], etaiOverEtat float64) Vector[T] {
	cosTheta := math.Min(v.Scale(-1).Dot(normal), 1.0)
	perpendicular := v.Add(normal.Scale(cosTheta)).Scale(etaiOverEtat)
	parallel := normal.Scale(-math.Sqrt(math.Abs(1.0 - perpendicular.LengthSquared())))
	return perpendicular.Add(parallel)
}

// MultByVector is component wise multiplication, also known as Hadamard product.
func (v Vector[T]) MultByVector(o Vector[T]) Vector[T] {
	return Vector[T]{
		X: v.X * o.X,
		Y: v.Y * o.Y,
		Z: v.Z * o.Z,
	}
}

func (v Vector[T]) DivByConstant(t float64) Vector[T] {
	return Vector[T]{
		X: T(float64(v.X) / t),
		Y: T(float64(v.Y) / t),
		Z: T(float64(v.Z) / t),
	}
}

func (v Vector[T]) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

func (v Vector[T]) LengthSquared() float64 {
	return float64((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

func (v Vector[T]) DistanceSquared(other Vector[T]) float64 {
	xDist := other.X - v.X
	yDist := other.Y - v.Y
	zDist := other.Z - v.Z
	return float64((xDist * xDist) + (yDist * yDist) + (zDist * zDist))
}

func (v Vector[T]) Distance(other Vector[T]) float64 {
	return math.Sqrt(v.DistanceSquared(other))
}

func (v Vector[T]) Angle(other Vector[T]) float64 {
	denominator := math.Sqrt(v.LengthSquared() * other.LengthSquared())
	if denominator < 1e-15 {
		return 0.
	}
	return math.Acos(mathex.Clamp(v.Dot(other)/denominator, -1., 1.))
}

func (v Vector[T]) NearZero() bool {
	const s = 1e-8
	return (math.Abs(float64(v.X)) < s) && (math.Abs(float64(v.Y)) < s) && (math.Abs(float64(v.Z)) < s)
}

func (v Vector[T]) Flip() Vector[T] {
	return Vector[T]{
		X: v.X * -1,
		Y: v.Y * -1,
		Z: v.Z * -1,
	}
}

func (v Vector[T]) FlipX() Vector[T] {
	return Vector[T]{
		X: v.X * -1,
		Y: v.Y,
		Z: v.Z,
	}
}

func (v Vector[T]) FlipY() Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y * -1,
		Z: v.Z,
	}
}

func (v Vector[T]) FlipZ() Vector[T] {
	return Vector[T]{
		X: v.X,
		Y: v.Y,
		Z: v.Z * -1,
	}
}

// Log returns the natural logarithm for each component
func (v Vector[T]) Log() Vector[T] {
	return Vector[T]{
		X: T(math.Log(float64(v.X))),
		Y: T(math.Log(float64(v.Y))),
		Z: T(math.Log(float64(v.Z))),
	}
}

// Log10 returns the decimal logarithm for each component.
func (v Vector[T]) Log10() Vector[T] {
	return Vector[T]{
		X: T(math.Log10(float64(v.X))),
		Y: T(math.Log10(float64(v.Y))),
		Z: T(math.Log10(float64(v.Z))),
	}
}

// Log2 returns the binary logarithm for each component
func (v Vector[T]) Log2() Vector[T] {
	return Vector[T]{
		X: T(math.Log2(float64(v.X))),
		Y: T(math.Log2(float64(v.Y))),
		Z: T(math.Log2(float64(v.Z))),
	}
}

// Exp2 returns 2**x, the base-2 exponential for each component
func (v Vector[T]) Exp2() Vector[T] {
	return Vector[T]{
		X: T(math.Exp2(float64(v.X))),
		Y: T(math.Exp2(float64(v.Y))),
		Z: T(math.Exp2(float64(v.Z))),
	}
}

// Exp returns e**x, the base-e exponential for each component
func (v Vector[T]) Exp() Vector[T] {
	return Vector[T]{
		X: T(math.Exp(float64(v.X))),
		Y: T(math.Exp(float64(v.Y))),
		Z: T(math.Exp(float64(v.Z))),
	}
}

// Expm1 returns e**x - 1, the base-e exponential for each component minus 1. It is more accurate than Exp(x) - 1 when the component is near zero
func (v Vector[T]) Expm1() Vector[T] {
	return Vector[T]{
		X: T(math.Expm1(float64(v.X))),
		Y: T(math.Expm1(float64(v.Y))),
		Z: T(math.Expm1(float64(v.Z))),
	}
}
