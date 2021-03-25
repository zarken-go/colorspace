package colorspace

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	assertDelta = 1e-12
)

func TestXyzToLabStep(t *testing.T) {
	assert.InDelta(t, 1, xyzToLabStep(1), 1e-12)
	assert.InDelta(t, 0.987921360533, xyzToLabStep(0.9642), assertDelta)
	assert.InDelta(t, 0.176866219667, xyzToLabStep(0.005), assertDelta)
	assert.InDelta(t, 0.138008904853, xyzToLabStep(0.00001), assertDelta)
}

func TestRgbToSrgbStep(t *testing.T) {
	assert.InDelta(t, 0.003035269835, rgbToSRGBStep(10), assertDelta)
	assert.InDelta(t, 0.577580440429, rgbToSRGBStep(200), assertDelta)
}

func TestRGBToSRGB(t *testing.T) {
	srgb := RGB{Red: 43, Green: 117, Blue: 150}.ToSRGB()
	assert.InDelta(t, 0.024157632448, srgb.Red, assertDelta)
	assert.InDelta(t, 0.177888415983, srgb.Green, assertDelta)
	assert.InDelta(t, 0.304987314069, srgb.Blue, assertDelta)
}

func TestSrgbToXyz(t *testing.T) {
	xyz := SRGB{Red: 0.024157632448, Green: 0.177888415983, Blue: 0.304987314069}.ToXYZ()
	assert.InDelta(t, 0.128603764617, xyz.X, assertDelta)
	assert.InDelta(t, 0.154367425188, xyz.Y, assertDelta)
	assert.InDelta(t, 0.311500632336, xyz.Z, assertDelta)
}

func TestIntToRGB(t *testing.T) {
	rgb := IntToRGB(2822550)
	assert.Equal(t, 43, rgb.Red)
	assert.Equal(t, 17, rgb.Green)
	assert.Equal(t, 150, rgb.Blue)
}

func TestIntToLab(t *testing.T) {
	Lab := IntToLab(2848150, IlluminantDefault)
	assert.InDelta(t, 46.226667406495, Lab.L, assertDelta)
	assert.InDelta(t, -11.528703049555, Lab.A, assertDelta)
	assert.InDelta(t, -24.496721810427, Lab.B, assertDelta)
}

func TestXYZ_ToLab(t *testing.T) {
	XYZ := XYZ{X: 0.56, Y: 0.45, Z: 0.13}
	Lab := XYZ.ToLab(IlluminantDefault)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 36.011760445426, Lab.A, assertDelta)
	assert.InDelta(t, 54.779683776742, Lab.B, assertDelta)
}

func TestIlluminants(t *testing.T) {
	XYZ := XYZ{X: 0.56, Y: 0.45, Z: 0.13}
	Lab := XYZ.ToLab(IlluminantATwo)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 16.267907022539, Lab.A, assertDelta)
	assert.InDelta(t, 10.288416924947, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantATen)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 14.711752307274, Lab.A, assertDelta)
	assert.InDelta(t, 9.769048437474, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantCTwo)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 31.654169076204, Lab.A, assertDelta)
	assert.InDelta(t, 57.447048903138, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantCTen)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 32.772544284235, Lab.A, assertDelta)
	assert.InDelta(t, 56.876556100880, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantD50Two)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 34.009743392026, Lab.A, assertDelta)
	assert.InDelta(t, 45.245506340821, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantD50Ten)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 33.580866558918, Lab.A, assertDelta)
	assert.InDelta(t, 44.763910895405, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantD55Two)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 35.082426076692, Lab.A, assertDelta)
	assert.InDelta(t, 49.146662543010, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantD55Ten)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 34.912091382115, Lab.A, assertDelta)
	assert.InDelta(t, 48.681939398346, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantD65Two)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 36.011760445426, Lab.A, assertDelta)
	assert.InDelta(t, 54.779683776741, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantD65Ten)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 36.359263453464, Lab.A, assertDelta)
	assert.InDelta(t, 54.298972220383, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantD72Two)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 36.122070901543, Lab.A, assertDelta)
	assert.InDelta(t, 58.608513799276, Lab.B, assertDelta)

	Lab = XYZ.ToLab(IlluminantD72Ten)
	assert.InDelta(t, 72.891894157652, Lab.L, assertDelta)
	assert.InDelta(t, 36.943477539671, Lab.A, assertDelta)
	assert.InDelta(t, 58.089095878826, Lab.B, assertDelta)

	Lab = XYZ.ToLab(Illuminant(0))
	assert.Equal(t, math.Inf(1), Lab.L)
	assert.True(t, math.IsNaN(Lab.A))
	assert.True(t, math.IsNaN(Lab.B))
}
