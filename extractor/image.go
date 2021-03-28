package extractor

import (
	"image"
	"io"
	"math"
	"os"

	"github.com/zarken-go/colorspace/palette"
)

func FromPath(path string, backgroundColor uint32) (palette.Palette, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer f.Close()

	return FromReader(f, backgroundColor)
}

func FromReader(r io.Reader, backgroundColor uint32) (palette.Palette, error) {
	i, _, err := image.Decode(r)
	if err != nil {
		return nil, err
	}
	return FromImage(i, backgroundColor), nil
}

func FromImage(i image.Image, backgroundColor uint32) palette.Palette {
	// base percent of color for each background component.
	bgRedP := float64((backgroundColor>>16)&0xFF) / 255
	bgGreenP := float64((backgroundColor>>8)&0xFF) / 255
	bgBlueP := float64(backgroundColor&0xFF) / 255

	colors := make(map[uint32]int)

	for x := i.Bounds().Min.X; x < i.Bounds().Max.X; x++ {
		for y := i.Bounds().Min.Y; y < i.Bounds().Max.Y; y++ {
			// rgb values have already been scaled by the alpha percent
			red, green, blue, alpha := i.At(x, y).RGBA()

			// percent of color for the alpha (1-alphaPercent) is the background's weight
			alphaPercent := float64(alpha) / 65535

			// percent of color for component at pixel (x, y)
			redPercent := float64(red) / 65535
			greenPercent := float64(green) / 65535
			bluePercent := float64(blue) / 65535

			colorRed := math.Round((redPercent + (bgRedP * (1 - alphaPercent))) * 255)
			colorGreen := math.Round((greenPercent + (bgGreenP * (1 - alphaPercent))) * 255)
			colorBlue := math.Round((bluePercent + (bgBlueP * (1 - alphaPercent))) * 255)

			color := uint32(colorRed*65536) +
				uint32(colorGreen*256) +
				uint32(colorBlue)
			if count, ok := colors[color]; ok {
				colors[color] = count + 1
			} else {
				colors[color] = 1
			}
		}
	}

	return palette.New(colors)
}
