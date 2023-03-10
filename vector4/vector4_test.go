package vector4_test

import (
	"encoding/json"
	"testing"

	"github.com/EliCDavis/vector/vector4"
	"github.com/stretchr/testify/assert"
)

func TestDefaults(t *testing.T) {
	tests := map[string]struct {
		got  vector4.Float64
		want vector4.Float64
	}{
		"zero": {got: vector4.Zero[float64](), want: vector4.New(0., 0., 0., 0.)},
		"one":  {got: vector4.One[float64](), want: vector4.New(1., 1., 1., 1.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), tc.got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), tc.got.W(), 0.000001)
		})
	}
}

func TestVectorOperations(t *testing.T) {
	start := vector4.New(1.2, -2.4, 3.7, 4.9)

	tests := map[string]struct {
		want vector4.Float64
		got  vector4.Float64
	}{
		"x":   {want: start.SetX(4), got: vector4.New(4., -2.4, 3.7, 4.9)},
		"y":   {want: start.SetY(4), got: vector4.New(1.2, 4., 3.7, 4.9)},
		"z":   {want: start.SetZ(4), got: vector4.New(1.2, -2.4, 4., 4.9)},
		"w":   {want: start.SetW(4), got: vector4.New(1.2, -2.4, 3.7, 4.)},
		"add": {want: start.Add(vector4.New(1., -2., 3., 4.)), got: vector4.New(2.2, -4.4, 6.7, 8.9)},
		"sub": {want: start.Sub(vector4.New(1., -2., 3., 4.)), got: vector4.New(0.2, -0.4, 0.7, 0.9)},
		"div": {want: start.DivByConstant(2), got: vector4.New(0.6, -1.2, 1.85, 2.45)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.InDelta(t, tc.want.X(), tc.got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), tc.got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), tc.got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), tc.got.W(), 0.000001)
		})
	}
}

func TestScaleVecFloat(t *testing.T) {
	tests := map[string]struct {
		vec    vector4.Float64
		scalar float64
		want   vector4.Float64
	}{
		"1, 2, 3, 4 *  2 =  2,  4,  6,  8": {vec: vector4.New(1., 2., 3., 4.), scalar: 2, want: vector4.New(2., 4., 6., 8.)},
		"1, 2, 3, 4 *  0 =  0,  0,  0,  0": {vec: vector4.New(1., 2., 3., 4.), scalar: 0, want: vector4.New(0., 0., 0., 0.)},
		"1, 2, 3, 4 * -2 = -2, -4, -6, -8": {vec: vector4.New(1., 2., 3., 4.), scalar: -2, want: vector4.New(-2., -4., -6., -8.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := tc.vec.Scale(tc.scalar)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), got.W(), 0.000001)
		})
	}
}

func TestJSON(t *testing.T) {
	in := vector4.New(1.2, 2.3, 3.4, 5.6)
	out := vector4.New(0., 0., 0., 0.)

	marshalledData, marshallErr := json.Marshal(in)
	unmarshallErr := json.Unmarshal(marshalledData, &out)

	assert.NoError(t, marshallErr)
	assert.NoError(t, unmarshallErr)
	assert.Equal(t, "{\"x\":1.2,\"y\":2.3,\"z\":3.4,\"w\":5.6}", string(marshalledData))
	assert.Equal(t, 1.2, out.X())
	assert.Equal(t, 2.3, out.Y())
	assert.Equal(t, 3.4, out.Z())
	assert.Equal(t, 5.6, out.W())
}

func TestAverage(t *testing.T) {
	// ASSIGN =================================================================
	vals := []vector4.Float64{
		vector4.New(1., 2., 3., 4.),
		vector4.New(1., 2., 3., 4.),
		vector4.New(1., 2., 3., 4.),
	}

	// ACT ====================================================================
	avg := vector4.Average(vals)

	// ASSERT =================================================================
	assert.InDelta(t, 1., avg.X(), 0.000001)
	assert.InDelta(t, 2., avg.Y(), 0.000001)
	assert.InDelta(t, 3., avg.Z(), 0.000001)
	assert.InDelta(t, 4., avg.W(), 0.000001)
}

func TestLerp(t *testing.T) {
	tests := map[string]struct {
		left  vector4.Float64
		right vector4.Float64
		t     float64
		want  vector4.Float64
	}{
		"(0, 0, 0, 0) =(0)=> (0, 0, 0, 0) = (0, 0, 0, 0)": {
			left:  vector4.New(0., 0., 0., 0.),
			right: vector4.New(0., 0., 0., 0.),
			t:     0,
			want:  vector4.New(0., 0., 0., 0.),
		},
		"(0, 0, 0, 0) =(0.5)=> (1, 2, 3, 4) = (0.5, 1, 1.5, 2.0)": {
			left:  vector4.New(0., 0., 0., 0.),
			right: vector4.New(1., 2., 3., 4.),
			t:     0.5,
			want:  vector4.New(0.5, 1., 1.5, 2.),
		},
		"(0, 0, 0, 0) =(1)=> (1, 2, 3, 4) = (1, 2, 3, 4)": {
			left:  vector4.New(0., 0., 0., 0.),
			right: vector4.New(1., 2., 3., 4.),
			t:     1,
			want:  vector4.New(1., 2., 3., 4.),
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector4.Lerp(tc.left, tc.right, tc.t)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), got.W(), 0.000001)
		})
	}
}

func TestMin(t *testing.T) {
	tests := map[string]struct {
		left  vector4.Float64
		right vector4.Float64
		want  vector4.Float64
	}{
		"(1, 2, 3, 4) m (4, 3, 2, 1) = (1, 2, 2, 1)": {left: vector4.New(1., 2., 3., 4.), right: vector4.New(4., 3., 2., 1.), want: vector4.New(1., 2., 2., 1.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector4.Min(tc.left, tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), got.W(), 0.000001)
		})
	}
}

func TestMax(t *testing.T) {
	tests := map[string]struct {
		left  vector4.Float64
		right vector4.Float64
		want  vector4.Float64
	}{
		"(1, 2, 3, 4) m (4, 3, 2, 1) = (1, 2, 2, 1)": {left: vector4.New(1., 2., 3., 4.), right: vector4.New(4., 3., 2., 1.), want: vector4.New(4., 3., 3., 4.)},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := vector4.Max(tc.left, tc.right)

			assert.InDelta(t, tc.want.X(), got.X(), 0.000001)
			assert.InDelta(t, tc.want.Y(), got.Y(), 0.000001)
			assert.InDelta(t, tc.want.Z(), got.Z(), 0.000001)
			assert.InDelta(t, tc.want.W(), got.W(), 0.000001)
		})
	}
}

func TestToInt(t *testing.T) {
	in := vector4.New(1.2, 2.3, 3.4, 5.6)
	out := in.ToInt()
	assert.Equal(t, 1, out.X())
	assert.Equal(t, 2, out.Y())
	assert.Equal(t, 3, out.Z())
	assert.Equal(t, 5, out.W())
}

func TestToInt64(t *testing.T) {
	in := vector4.New(1.2, 2.3, 3.4, 5.6)
	out := in.ToInt64()
	assert.Equal(t, int64(1), out.X())
	assert.Equal(t, int64(2), out.Y())
	assert.Equal(t, int64(3), out.Z())
	assert.Equal(t, int64(5), out.W())
}

func TestToFloat32(t *testing.T) {
	in := vector4.New(1.2, 2.3, 3.4, 5.6)
	out := in.ToFloat32()
	assert.Equal(t, float32(1.2), out.X())
	assert.Equal(t, float32(2.3), out.Y())
	assert.Equal(t, float32(3.4), out.Z())
	assert.Equal(t, float32(5.6), out.W())
}

func TestToFloat64(t *testing.T) {
	in := vector4.New(1, 2, 3, 5)
	out := in.ToFloat64()
	assert.Equal(t, float64(1), out.X())
	assert.Equal(t, float64(2), out.Y())
	assert.Equal(t, float64(3), out.Z())
	assert.Equal(t, float64(5), out.W())
}
