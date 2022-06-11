// Package pintor colorizes text printed to a terminal.
// It applies format to the output using ANSI escape sequences.
package pintor

import (
	"fmt"
	"os"
	"strconv"
)

const csi = "\x1b["

var (
	// outTTY indicates if standard output is a TTY
	outTTY bool
)

func init() {
	outTTY = isOutputTTY(os.Stdout)
}

// isOutputTTY determines if output is a TTY,
// it's outcome is used by the outTTY variable
// to print with fomatting only when a TTY is present in standard output
func isOutputTTY(f *os.File) bool {
	out, _ := f.Stat()
	if (out.Mode() & os.ModeCharDevice) == 0 {
		return false
	}
	return true
}

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
//
// Default leaves the current format as it is.
// Eight colors supported: Black, Red, Green, Yellow, Blue, Magenta, Cyan and White.
// Colors apply for foreground and background color.
// Modifiers supported: Bold, Italic and Underline.
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

// Formatter formats the text output.
type Formatter struct {
	foreground uint
	background uint
	modifiers  uint
}

// NewFormatter returns a new Formatter that can apply the colors and modifiers specified as parameters.
// Predefined bits control the colors and modifiers.
// Look at the package constant to get the colors and modifiers names.
// Only accepts one color for foreground and one color for background.
// To apply multiple modifiers at once set the modifiers parameter to:
//
// 		Bold|Italic
//
// This for example will apply both Bold and Italic to the text.
func NewFormatter(foreground, background, modifiers uint) *Formatter {
	return &Formatter{
		foreground: foreground,
		background: background,
		modifiers:  modifiers,
	}
}

// Format applies the format defined by the Formatter.
// It receives the target string and returns the string with the formatting applied.
// Formatting is performed using ANSI escape sequences.
func (f *Formatter) Format(target string) string {
	if !outTTY {
		return target
	}
	sequence := f.compile()
	end := buildEscape([]color{reset})
	return fmt.Sprintf("%s%s%s", sequence, target, end)
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
