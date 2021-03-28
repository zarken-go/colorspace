package extractor

import (
	"errors"
	_ "image/jpeg"
	"image/png"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransparencyHandling(t *testing.T) {
	f, err := os.Open(`../assets/red-transparent-50.png`)
	if err != nil {
		t.Fatal(err)
	}
	Img, err := png.Decode(f)
	if err != nil {
		t.Fatal(err)
	}
	palette := FromImage(Img, 0xffffff) // Expected: 16744576
	assert.Equal(t, uint32(16744576), palette[0].Color)
	assert.Equal(t, 1, len(palette))

	palette = FromImage(Img, 0x000000) // Expected: 8323072
	assert.Equal(t, uint32(8323072), palette[0].Color)
	assert.Equal(t, 1, len(palette))
}

func TestPNGExtractingColors(t *testing.T) {
	p, err := FromPath(`../assets/test.png`, 0xffffff)
	if err != nil {
		t.Fatal(err)
	}
	mergeThree := p.MergeColors(3)
	assert.Equal(t, 3, len(mergeThree))
	assert.Equal(t, []uint32{14024704, 3407872, 7111569}, mergeThree.Colors())

	mergeOne := p.MergeColors(1)
	assert.Equal(t, 1, len(mergeOne))
	assert.Equal(t, []uint32{14024704}, mergeOne.Colors())
}

func TestJPEGExtractingColors(t *testing.T) {
	p, err := FromPath(`../assets/test.jpeg`, 0xffffff)
	if err != nil {
		t.Fatal(err)
	}
	mergeThree := p.MergeColors(3)
	assert.Equal(t, 3, len(mergeThree))
	assert.Equal(t, []uint32{15985689, 15216177, 2957328}, mergeThree.Colors())

	mergeOne := p.MergeColors(1)
	assert.Equal(t, 1, len(mergeOne))
	assert.Equal(t, []uint32{15985689}, mergeOne.Colors())
}

func TestFileNotFound(t *testing.T) {
	_, err := FromPath(`../not_found.jpeg`, 0xffffff)
	assert.EqualError(t, err, `open ../not_found.jpeg: The system cannot find the file specified.`)
}

type brokenReader struct{}

func (b brokenReader) Read(p []byte) (n int, err error) {
	return 0, errors.New(`failed to read data`)
}

func TestFromReaderErr(t *testing.T) {
	_, err := FromReader(&brokenReader{}, 0xffffff)
	assert.EqualError(t, err, `image: unknown format`)
}
