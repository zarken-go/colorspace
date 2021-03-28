package colorspace

import (
	"math"
)

type RGB struct {
	Red   uint8
	Green uint8
	Blue  uint8
}

type SRGB struct {
	Red   float64
	Green float64
	Blue  float64
}

type XYZ struct {
	X float64
	Y float64
	Z float64
}

type Lab struct {
	L float64
	A float64
	B float64
}

func IntToRGB(color int) RGB {
	return RGB{
		Red:   uint8(color>>16) & 0xff,
		Green: uint8(color>>8) & 0xff,
		Blue:  uint8(color) & 0xff,
	}
}

func UInt32ToRGB(color uint32) RGB {
	return RGB{
		Red:   uint8(color>>16) & 0xff,
		Green: uint8(color>>8) & 0xff,
		Blue:  uint8(color) & 0xff,
	}
}

func IntToLab(color int, observer Illuminant) Lab {
	return IntToRGB(color).
		ToSRGB().
		ToXYZ().
		ToLab(observer)
}

func UInt32ToLab(color uint32, observer Illuminant) Lab {
	return UInt32ToRGB(color).
		ToSRGB().
		ToXYZ().
		ToLab(observer)
}

func (RGB RGB) ToSRGB() SRGB {
	return SRGB{
		Red:   rgbToSRGBStep(RGB.Red),
		Green: rgbToSRGBStep(RGB.Green),
		Blue:  rgbToSRGBStep(RGB.Blue),
	}
}

func rgbToSRGBStep(value uint8) float64 {
	f := float64(value) / 255
	if f <= 0.03928 {
		return f / 12.92
	}
	return math.Pow((f+0.055)/1.055, 2.4)
}

func (SRGB SRGB) ToXYZ() XYZ {
	return XYZ{
		X: (0.4124564 * SRGB.Red) + (0.3575761 * SRGB.Green) + (0.1804375 * SRGB.Blue),
		Y: (0.2126729 * SRGB.Red) + (0.7151522 * SRGB.Green) + (0.0721750 * SRGB.Blue),
		Z: (0.0193339 * SRGB.Red) + (0.1191920 * SRGB.Green) + (0.9503041 * SRGB.Blue),
	}
}

func (XYZ XYZ) ToLab(observer Illuminant) Lab {
	xn, yn, zn := observer.xyzValues()

	return Lab{
		L: 116.0*xyzToLabStep(XYZ.Y/yn) - 16.0,
		A: 500.0 * (xyzToLabStep(XYZ.X/xn) - xyzToLabStep(XYZ.Y/yn)),
		B: 200.0 * (xyzToLabStep(XYZ.Y/yn) - xyzToLabStep(XYZ.Z/zn)),
	}
}

func xyzToLabStep(v float64) float64 {
	if v > 216.0/24389.0 {
		return math.Pow(v, 1.0/3.0)
	}
	return 841.0*v/108.0 + 4.0/29.0
}
