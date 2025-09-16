package catbow

import (
	"fmt"
	"testing"

	"github.com/lordxarus/catbow/catbow/encoder/ansi"
	"github.com/stretchr/testify/assert"
)

func setupRainbow() *rainbowStrategy {
	opts := NewRainbowOptions()
	opts.Seed = 1
	return NewRainbowStrategy(opts)
}

func TestColorGeneration(t *testing.T) {

	rb := setupRainbow()
	outR := rb.colorizeRune('r')
	outB := rb.colorizeRune('b')

	assert.Contains(t, outR, ansi.Esc+"[38;2")
	assert.Contains(t, outR, "mr")

	assert.Contains(t, outB, ansi.Esc+"[38;2")
	assert.Contains(t, outB, "mb")

	assert.NotEqual(t, outR, rb.colorizeRune('r'))
}

func TestNoColorGeneration(t *testing.T) {
	rb := setupRainbow()
	rb.Opts.NoColor = true

	assert.Equal(t, string('r'), rb.colorizeRune('r'))

}

func TestColorReset(t *testing.T) {
	rb := setupRainbow()
	out := rb.colorizeRune('a')
	assert.NotContains(t, out, ansi.Reset)
	out = rb.Cleanup()
	assert.Contains(t, out, ansi.Reset)
}

func TestRainbowAlgorithm(t *testing.T) {
	rb := setupRainbow()
	rb.Opts.Seed = 1
	rb.Opts.Spread = 3.0

	defer fmt.Println(rb.Cleanup())

	rgb := rb.calculateRainbow()
	assert.Equal(t, 0, rb.offset)
	assert.Equal(t, RgbColor{127, 236, 17}, rgb)

	rb.offset += 10
	rgb = rb.calculateRainbow()
	assert.Equal(t, rb.offset, 0)
	assert.Equal(t, RgbColor{233, 132, 14}, rgb)
}
