package catbow

import (
	"fmt"
	"math"

	"github.com/lordxarus/catbow/catbow/encoder/ansi"
)

type rainbowOptions struct {
	// Rotates the rainbow
	Spread float64
	// Controls the horizontal width of each color band
	Frequency float64
	// An offset for the starting color allowing varied but deterministic output
	Seed int
	// Disables catbow, input will equal output
	NoColor    bool
	redShift   float64
	greenShift float64
	blueShift  float64
}

func NewRainbowOptions() *rainbowOptions {
	return &rainbowOptions{
		Spread:     3.0,
		Frequency:  0.1,
		Seed:       0,
		NoColor:    false,
		redShift:   0,
		greenShift: (2 * math.Pi) / 3,
		blueShift:  (4 * math.Pi) / 3,
	}
}

type rainbowStrategy struct {
	Opts   rainbowOptions
	offset float64
}

func (rb *rainbowStrategy) Cleanup() string {
	fmt.Printf("%", rb.offset)
	return ansi.Reset
}

func NewRainbowStrategy(opts *rainbowOptions) *rainbowStrategy {
	return &rainbowStrategy{
		Opts:   *opts,
		offset: 0,
	}
}

func (rb *rainbowStrategy) calculateRainbow(offset float64) RgbColor {
	freq := rb.Opts.Frequency

	seed := float64(rb.Opts.Seed)
	o := math.Round(offset / rb.Opts.Spread)

	red := math.Sin((freq*o)+rb.Opts.redShift+seed)*127 + 128
	green := math.Sin((freq*o)+rb.Opts.greenShift+seed)*127 + 128
	blue := math.Sin((freq*o)+rb.Opts.blueShift+seed)*127 + 128

	return RgbColor{
		r: uint8(math.Floor(red)),
		g: uint8(math.Floor(green)),
		b: uint8(math.Floor(blue)),
	}
}

func (rb *rainbowStrategy) colorizeRune(r rune) string {
	/* TODO: Refactor match into a call to calculateRainbow
	and a call to the injected ColorFormatter which does what
	the fmt.Sprintf() call does but allows us to be agnostic as
	to what we're outputting to. Essentially this becomes the
	API for Colorizers to call

	*/
	if rb.Opts.NoColor {
		return string(r)
	}

	rgb := rb.calculateRainbow(rb.offset / rb.Opts.Spread)

	rb.offset += 1.0 / rb.Opts.Spread

	return fmt.Sprintf(
		ansi.Esc+"[38;2;%d;%d;%dm%c",
		rgb.r,
		rgb.g,
		rgb.b,
		r)
}
