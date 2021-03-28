package palette

import (
	"math"
	"sort"

	"github.com/zarken-go/colorspace"
	"github.com/zarken-go/colorspace/ciede2000"
)

type Palette []Color
type Color struct {
	Color uint32
	Lab   colorspace.Lab
	Score float64
}

func New(m map[uint32]int) Palette {
	var p Palette

	for color, count := range m {
		lab := colorspace.UInt32ToLab(color, colorspace.IlluminantDefault)
		score := math.Sqrt(math.Pow(lab.A, 2)+math.Pow(lab.B, 2)) *
			(1 - lab.L/200) *
			math.Sqrt(float64(count))
		p = append(p, Color{
			Color: color,
			Lab:   lab,
			Score: score,
		})
	}

	sort.Slice(p, func(i, j int) bool {
		return p[i].Score > p[j].Score
	})

	return p
}

func (P Palette) MergeColors(limit int) Palette {
	var merged Palette
	maxDelta := 100.0 / float64(limit)
	mergedCount := 0

	for _, color := range P {
		hasBeenMerged := false

		for _, mergedColor := range merged {
			if ciede2000.DeltaE(color.Lab, mergedColor.Lab) < maxDelta {
				hasBeenMerged = true
				break
			}
		}

		if hasBeenMerged {
			continue
		}

		merged = append(merged, color)
		mergedCount++

		if mergedCount == limit {
			break
		}
	}

	return merged
}

func (P Palette) Colors() []uint32 {
	var colors []uint32
	for _, c := range P {
		colors = append(colors, c.Color)
	}
	return colors
}
