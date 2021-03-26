package ciede2000

import (
	"math"

	"github.com/zarken-go/colorspace"
)

const (
	pow25To7 = 6103515625.0 // math.Pow(25, 7)
)

func deg2Rad(deg float64) float64 {
	return deg * (math.Pi / 180.0)
}

func DeltaE(Lab1, Lab2 colorspace.Lab) float64 {
	c1 := math.Sqrt(math.Pow(Lab1.A, 2) + math.Pow(Lab1.B, 2))
	c2 := math.Sqrt(math.Pow(Lab2.A, 2) + math.Pow(Lab2.B, 2))

	barC := (c1 + c2) / 2

	G := .5 * (1 - math.Sqrt(math.Pow(barC, 7)/(math.Pow(barC, 7)+pow25To7)))

	a1p := (1.0 + G) * Lab1.A
	a2p := (1.0 + G) * Lab2.A

	deg180InRad := math.Pi
	deg360InRad := 2 * math.Pi

	CPrime1 := math.Sqrt(math.Pow(a1p, 2) + math.Pow(Lab1.B, 2))
	CPrime2 := math.Sqrt(math.Pow(a2p, 2) + math.Pow(Lab2.B, 2))

	var h1p, h2p float64
	if !(a1p == 0 && Lab1.B == 0) {
		h1p = math.Atan2(Lab1.B, a1p)
		if h1p < 0 {
			h1p += deg360InRad
		}
	}
	if !(a2p == 0 && Lab2.B == 0) {
		h2p = math.Atan2(Lab2.B, a2p)
		if h2p < 0 {
			h2p += deg360InRad
		}
	}

	LpDelta := Lab2.L - Lab1.L
	CpDelta := CPrime2 - CPrime1

	var deltahPrime float64
	if CPrime1*CPrime2 == 0 {
		deltahPrime = 0
	} else {
		deltahPrime = h2p - h1p
		if deltahPrime < -deg180InRad {
			deltahPrime += deg360InRad
		} else if deltahPrime > deg180InRad {
			deltahPrime -= deg360InRad
		}
	}

	HpDelta := 2 * math.Sqrt(CPrime1*CPrime2) * math.Sin(deltahPrime/2)

	barLPrime := (Lab1.L + Lab2.L) / 2.0
	barCPrime := (CPrime1 + CPrime2) / 2.0

	var barhPrime float64

	hPrimeSum := h1p + h2p
	if CPrime1*CPrime2 == 0.0 {
		barhPrime = hPrimeSum
	} else {
		if math.Abs(h1p-h2p) <= deg180InRad {
			barhPrime = hPrimeSum / 2.0
		} else {
			if hPrimeSum < deg360InRad {
				barhPrime = (hPrimeSum + deg360InRad) / 2.0
			} else {
				barhPrime = (hPrimeSum - deg360InRad) / 2.0
			}
		}
	}

	T := 1.0 - (0.17 * math.Cos(barhPrime-deg2Rad(30))) +
		(0.24 * math.Cos(2.0*barhPrime)) +
		(0.32 * math.Cos(3.0*barhPrime+deg2Rad(6))) -
		(0.20 * math.Cos(4.0*barhPrime-deg2Rad(63)))

	deltaTheta := deg2Rad(30) *
		math.Exp(-math.Pow((barhPrime-deg2Rad(275))/deg2Rad(25), 2.0))

	Rc := 2.0 * math.Sqrt(math.Pow(barCPrime, 7.0)/(math.Pow(barCPrime, 7.0)+pow25To7))
	Sl := 1.0 + ((0.015 * math.Pow(barLPrime-50, 2.0)) /
		math.Sqrt(20.0+math.Pow(barLPrime-50, 2.0)))
	Sc := 1.0 + 0.045*barCPrime
	Sh := 1.0 + 0.015*barCPrime*T

	Rt := -math.Sin(2.0*deltaTheta) * Rc

	return math.Sqrt(
		math.Pow(LpDelta/Sl, 2) +
			math.Pow(CpDelta/Sc, 2) +
			math.Pow(HpDelta/Sh, 2) +
			(Rt * (CpDelta / Sc) * (HpDelta / Sh)))
}
