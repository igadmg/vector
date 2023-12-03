# Vector

![Coverage](https://img.shields.io/badge/Coverage-95.4%25-brightgreen)
[![Go Report Card](https://goreportcard.com/badge/github.com/EliCDavis/vector)](https://goreportcard.com/report/github.com/EliCDavis/vector)
[![GoDoc](https://godoc.org/github.com/EliCDavis/vector?status.svg)](http://godoc.org/github.com/EliCDavis/vector)

Collection of **generic, immutable** vector math functions I've written overtime for different hobby projects.

## API

| Function      | Vector2 | Vector3 | Vector4 | Description                                            |
|---------------|---------|---------|---------|--------------------------------------------------------|
| Abs           | ✅      | ✅     | ✅      | Returns a vector with each component's absolute value  |
| Add           | ✅      | ✅     | ✅      | Component Wise Addition                                |
| Angle         | ✅      | ✅     |         | Returns the angle between two vectors                  |
| Ceil          | ✅      | ✅     | ✅      | Ceils each vectors component to the nearest integer    |
| Clamp         | ✅      | ✅     | ✅      | Clamps each component between two values               |
| ContainsNaN   | ✅      | ✅     | ✅      | Returns true if any component of the vector is NaN     |
| Cross         |         | ✅     |          | Returns the cross product between two vectors          |
| Dot           | ✅      | ✅     | ✅      | Returns the dot product between two vectors            |
| Flip          | ✅      | ✅     | ✅      | Scales the vector by -1                                |
| FlipX         | ✅      | ✅     | ✅      | Returns a vector with the X component multiplied by -1 |
| FlipY         | ✅      | ✅     | ✅      | Returns a vector with the Y component multiplied by -1 |
| FlipZ         |         | ✅     | ✅      | Returns a vector with the Z component multiplied by -1 |
| FlipW         |         |         | ✅      | Returns a vector with the W component multiplied by -1 |
| Floor         | ✅      | ✅     | ✅      | Floors each vectors component                          |
| Format        | ✅      | ✅     | ✅      | Build a string with vector data                        |
| Length        | ✅      | ✅     | ✅      | Returns the length of the vector                       |
| LengthSquared | ✅      | ✅     | ✅      | Returns the squared length of the vector               |
| MaxComponent  | ✅      | ✅     | ✅      | Returns the vectors largest component                  |
| Midpoint      | ✅      | ✅     | ✅      | Finds the mid point between two vectors                |
| MinComponent  | ✅      | ✅     | ✅      | Returns the vectors smallest component                 |
| Normalized    | ✅      | ✅     | ✅      | Returns the normalized vector                          |
| Round         | ✅      | ✅     | ✅      | Rounds each vectors component to the nearest integer   |
| Scale         | ✅      | ✅     | ✅      | Scales the vector by some constant                     |
| Sqrt          | ✅      | ✅     | ✅      | Returns a vector with each component's square root     |
| Sub           | ✅      | ✅     | ✅      | Component Wise Subtraction                             |
| X             | ✅      | ✅     | ✅      | Returns the x component of the vector                  |
| Y             | ✅      | ✅     | ✅      | Returns the y component of the vector                  |
| Z             |         | ✅     | ✅      | Returns the z component of the vector                  |
| W             |         |         | ✅      | Returns the w component of the vector                  |


## Example

Below is an example on how to implement the different sign distance field functions in a generic fashion to work for both int, int64, float32, and float64.

```go
package sdf

import (
	"github.com/EliCDavis/vector"
	"github.com/EliCDavis/vector/vector3"
)

type Field[T vector.Number] func(v vector3.Vector[T]) float64

func Sphere[T vector.Number](pos vector3.Vector[T], radius float64) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		return v.Distance(pos) - radius
	}
}

func Box[T vector.Number](pos vector3.Vector[T], bounds vector3.Vector[T]) Field[T] {
	halfBounds := bounds.Scale(0.5)
	// It's best to watch the video to understand
	// https://www.youtube.com/watch?v=62-pRVZuS5c
	return func(v vector3.Vector[T]) float64 {
		q := v.Sub(pos).Abs().Sub(halfBounds)
		inside := math.Min(float64(q.MaxComponent()), 0)
		return vector3.Max(q, vector3.Zero[T]()).Length() + inside
	}
}

func Union[T vector.Number](a, b Field[T]) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		return math.Min(a(v), b(v))
	}
}

func Intersect[T vector.Number](a, b Field[T]) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		return math.Max(a(v), b(v))
	}
}

func Translate[T vector.Number](field Field[T], translation vector3.Vector[T]) Field[T] {
	return func(v vector3.Vector[T]) float64 {
		return field(v.Sub(translation))
	}
}
```
