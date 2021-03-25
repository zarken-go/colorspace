package colorspace

type Illuminant uint8

const (
	IlluminantATwo   Illuminant = iota + 1 // A 2°
	IlluminantATen                         // A 10°
	IlluminantCTwo                         // C 2°
	IlluminantCTen                         // C 10°
	IlluminantD50Two                       // D50 2°
	IlluminantD50Ten                       // D50 10°
	IlluminantD55Two                       // D55 2°
	IlluminantD55Ten                       // D55 10°
	IlluminantD65Two                       // D65 2°
	IlluminantD65Ten                       // D65 10°
	IlluminantD72Two                       // D72 2°
	IlluminantD72Ten                       // D72 10°

	// https://en.wikipedia.org/wiki/Illuminant_D65#Definition
	IlluminantDefault = IlluminantD65Two
)

func (Illuminant Illuminant) xyzValues() (float64, float64, float64) {
	switch Illuminant {
	case IlluminantATwo:
		return 1.0985, 1, .35585
	case IlluminantATen:
		return 1.11144, 1, .352
	case IlluminantCTwo:
		return .98074, 1, 1.18232
	case IlluminantCTen:
		return .97285, 1, 1.16145
	case IlluminantD50Two:
		return .96422, 1, .82521
	case IlluminantD50Ten:
		return .9672, 1, .81427
	case IlluminantD55Two:
		return .95682, 1, .92149
	case IlluminantD55Ten:
		return .95799, 1, .90926
	case IlluminantD65Two:
		return .95047, 1, 1.08883
	case IlluminantD65Ten:
		return .94811, 1, 1.07304
	case IlluminantD72Two:
		return .94972, 1, 1.22638
	case IlluminantD72Ten:
		return .94416, 1, 1.20641
	}

	return 0, 0, 0
}
