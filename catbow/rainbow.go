package catbow

import (
	"fmt"
	"github.com/lordxarus/catbow/catbow/encoder/ansi"
	"math"
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
	Opts                rainbowOptions
	offset              float64
	currLineLength      float64
	prevLineStartOffset float64
	wasLastRuneNewLine  bool
	lastLineBuffer      [512]rune
}

// extract to ColorEncoder
func (rb *rainbowStrategy) CleanupStr() string {
	return ansi.Reset
}

func NewRainbowStrategy(opts *rainbowOptions) *rainbowStrategy {
	s := &rainbowStrategy{
		Opts: *opts,
	}
	s.offset = float64(opts.Seed) + 1.0
	// s.offset = &offset

	// var ll = 0.0
	s.currLineLength = 0.0

	// var llSo = float64(opts.Seed) + 1.0
	s.prevLineStartOffset = s.offset

	// b := false
	s.wasLastRuneNewLine = false

	return s
}

func (rb *rainbowStrategy) calculateRainbow(offset float64) RgbColor {
	/*

		Color math:

		freq: .05
		spread: 1.05
		seed: 1

		// first character, offset starts at seed + 1
		offset: 2
		scled_offset = 2 / 1.05 = 1.9047619047619047
		red = sin(.05 * scled_offset) + 0 * 127 + 128
		red =




	*/

	freq := rb.Opts.Frequency

	scled_offst := offset / rb.Opts.Spread

	red := math.Sin((freq*scled_offst)+rb.Opts.redShift)*127 + 128
	green := math.Sin((freq*scled_offst)+rb.Opts.greenShift)*127 + 128
	blue := math.Sin((freq*scled_offst)+rb.Opts.blueShift)*127 + 128

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

	// this is again to deal with the prefix nature of the lolcat code

	// since we're doing essentially a postfix operation instead of
	// lolcat's prefix increment we add 1 to the Seed to derive the
	// starting offset when creating the strategy
	if r == '\n' || r == '\r' {
		rb.wasLastRuneNewLine = true
	}

	// what about mutliple newlines in a row?
	if rb.wasLastRuneNewLine {
		rb.wasLastRuneNewLine = false
		rb.prevLineStartOffset += 1
		rb.offset = rb.prevLineStartOffset
	} else {
		/*
			A small deviation from lolcat. Offset
			accumulates the spread:

			off += (1 / Spread) instead of
			off = (charIndex / Spread).

			Importantly this allows for the offset to be testable
		*/

		rb.offset = rb.offset + (1 / rb.Opts.Spread)
	}

	rgb := rb.calculateRainbow(rb.offset)

	return fmt.Sprintf(
		ansi.Esc+"[38;2;%d;%d;%dm%c",
		rgb.r,
		rgb.g,
		rgb.b,
		r)
}
