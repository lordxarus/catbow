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
	opts.Spread = 3
	opts.Frequency = .1
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

	// increase the offset to avoid collisions
	for range 20 {
		rb.colorizeRune('r')
	}
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
	assert.NotContains(t, out, ansi.Esc+"[0m")
	out = rb.CleanupStr()
	assert.Equal(t, out, ansi.Esc+"[0m")
}

func TestRainbowAlgorithm(t *testing.T) {
	rb := setupRainbow()
	rb.Opts.Seed = 1
	rb.Opts.Spread = 3.0

	defer fmt.Println(rb.CleanupStr())

	assert.Equal(t, 2, rb.offset)
	rgb := rb.calculateRainbow(rb.offset)
	assert.Equal(t, RgbColor{127, 236, 17}, rgb)

	rb.offset += 10
	rgb = rb.calculateRainbow(rb.offset)
	assert.Equal(t, RgbColor{233, 132, 14}, rgb)
}

func TestOffsetProgression(t *testing.T) {
	/*
		Explanation of the lolcat offset:
	*/
	rb := setupRainbow()
	defer fmt.Println(rb.CleanupStr())

	// offset gets initialized to seed (1 if using setupRainbow() test setup)
	// in lolcat the input is loaded into memory and split on lines. before the first
	// iteration the offset is incremented so we start with seed + 1

	assert.Equal(t, 2.0, rb.offset)
	rb.colorizeRune('a')
	assert.Equal(t, 2.3333333333333335, rb.offset)
	rb.colorizeRune('s')
	assert.Equal(t, 2.666666666666667, rb.offset)
	rb.colorizeRune('d')
	assert.Equal(t, 3.0000000000000004, rb.offset)
	rb.colorizeRune('f')
	assert.Equal(t, 3.333333333333334, rb.offset)
	rb.colorizeRune('\n')

	// rb.offset is reset to what it was when entering the
	// line loop in this case it gets set back to 2.0

	// then incremented again before the next iteration of the line loop
	assert.Equal(t, 3.0, rb.offset)

	// EOF

}

// logs from echo 'asdf\n' | lolcat

/*
buf.lines.each @os bef: 1
buf.lines.each @os after: 2
{spread: 3.0, freq: 0.1, seed: 1, animate: false, duration: 12, speed: 20.0, invert: false, truecolor: false, force: false, version: false, help: false, seed_given: true, os: 1}
calling rainbow with offset of:2.0
calling rainbow with offset of:2.3333333333333335
calling rainbow with offset of:2.6666666666666665
calling rainbow with offset of:3.0
calling rainbow with offset of:3.333333333333333

** newline reached **
buf.lines.each @os bef: 2
buf.lines.each @os after: 3
{spread: 3.0, freq: 0.1, seed: 1, animate: false, duration: 12, speed: 20.0, invert: false, truecolor: false, force: false, version: false, help: false, seed_given: true, os: 1}
calling rainbow with offset of:3.0 */
