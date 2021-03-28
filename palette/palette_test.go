package palette

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPalette(t *testing.T) {
	m := make(map[uint32]int)
	m[16711680] = 10 // #ff0000
	m[16715535] = 1  // #ff0f0f
	m[16715520] = 3  // #ff0f00
	m[255] = 2       // #0000ff

	p := New(m)
	assert.Equal(t, []uint32{16711680, 255, 16715520, 16715535}, p.Colors())

	merged := p.MergeColors(3)
	assert.Equal(t, []uint32{16711680, 255}, merged.Colors())

	mergedLimit := p.MergeColors(2)
	assert.Equal(t, []uint32{16711680, 255}, mergedLimit.Colors())
}
