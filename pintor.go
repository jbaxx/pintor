// Package pintor colorizes text printed to a terminal
package pintor

import (
	"fmt"
	"strconv"
)

const csi = "\x1b["

type color int

const (
	reset     color = 0
	bold            = 1
	italic          = 3
	underline       = 4
)

const (
	fgBlack color = iota + 30
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
)

const (
	bgBlack color = iota + 40
	bgRed
	bgGreen
	bgYellow
	bgBlue
	bgMagenta
	bgCyan
	bgWhite
)

// Values that can be used to create a Formatter object.
// Default leaves the current format as it is.
// Black, Red, Green, Yellow, Blue, Magenta, Cyan and White are the colors supported,
// both for foreground and background color.
// Bold, Italic and Underline are the modifiers supported.
const (
	Default = 0
	Black   = 1 << iota
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
	Bold
	Italic
	Underline
)

var fgUMap = map[uint]color{
	Black:   fgBlack,
	Red:     fgRed,
	Green:   fgGreen,
	Yellow:  fgYellow,
	Blue:    fgBlue,
	Magenta: fgMagenta,
	Cyan:    fgCyan,
	White:   fgWhite,
}

var bgUMap = map[uint]color{
	Black:   bgBlack,
	Red:     bgRed,
	Green:   bgGreen,
	Yellow:  bgYellow,
	Blue:    bgBlue,
	Magenta: bgMagenta,
	Cyan:    bgCyan,
	White:   bgWhite,
}

// Formatter format the text output based on the
// foreground, background and modifiers parameters.
type Formatter struct {
	foreground uint
	background uint
	modifiers  uint
}

// NewFormatter returns a new Formatter
func NewFormatter(foreground, background, modifiers uint) *Formatter {
	return &Formatter{
		foreground: foreground,
		background: background,
		modifiers:  modifiers,
	}
}

// Format applies the format defined at the Formatter.
// It receives a string and returns a string with the formatting applied.
func (f *Formatter) Format(texto string) string {
	sequence := f.compile()
	end := buildEscape([]color{reset})
	return fmt.Sprintf("%s%s%s", sequence, texto, end)
}

func (f *Formatter) compile() string {
	var p []color

	if g, ok := fgUMap[f.foreground]; ok {
		p = append(p, g)
	}

	if b, ok := bgUMap[f.background]; ok {
		p = append(p, b)
	}

	if f.modifiers&Bold != 0 {
		p = append(p, bold)
	}

	if f.modifiers&Italic != 0 {
		p = append(p, italic)
	}

	if f.modifiers&Underline != 0 {
		p = append(p, underline)
	}

	return buildEscape(p)
}

func buildEscape(colors []color) string {
	escape := csi
	for i, c := range colors {
		if i > 0 {
			escape += ";"
		}
		escape += strconv.Itoa(int(c))
	}
	escape += "m"
	return escape
}
