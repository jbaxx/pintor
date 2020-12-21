package pintor

import (
	"fmt"
	"testing"
)

var endReset = "\x1b[0m"

func TestFormat(t *testing.T) {

	formatTests := []struct {
		name       string
		text       string
		escape     string
		foreground uint
		background uint
		modifiers  uint
	}{
		{
			name:   "testing basic line",
			text:   "this line",
			escape: "\x1b[m",
		},
		{name: "black foreground ", text: "this line", escape: "\x1b[30m", foreground: Black},
		{name: "red foreground   ", text: "this line", escape: "\x1b[31m", foreground: Red},
		{name: "green foreground ", text: "this line", escape: "\x1b[32m", foreground: Green},
		{name: "yellow foreground", text: "this line", escape: "\x1b[33m", foreground: Yellow},
		{name: "blue foreground  ", text: "this line", escape: "\x1b[34m", foreground: Blue},
		{name: "magent foreground", text: "this line", escape: "\x1b[35m", foreground: Magenta},
		{name: "cyan foreground  ", text: "this line", escape: "\x1b[36m", foreground: Cyan},
		{name: "white foreground ", text: "this line", escape: "\x1b[37m", foreground: White},
		{name: "black background ", text: "this line", escape: "\x1b[40m", background: Black},
		{name: "red background   ", text: "this line", escape: "\x1b[41m", background: Red},
		{name: "green background ", text: "this line", escape: "\x1b[42m", background: Green},
		{name: "yellow background", text: "this line", escape: "\x1b[43m", background: Yellow},
		{name: "blue background  ", text: "this line", escape: "\x1b[44m", background: Blue},
		{name: "magent background", text: "this line", escape: "\x1b[45m", background: Magenta},
		{name: "cyan background  ", text: "this line", escape: "\x1b[46m", background: Cyan},
		{name: "white background ", text: "this line", escape: "\x1b[47m", background: White},
		{name: "bold     ", text: "this line", escape: "\x1b[1m", modifiers: Bold},
		{name: "italic   ", text: "this line", escape: "\x1b[3m", modifiers: Italic},
		{name: "underline", text: "this line", escape: "\x1b[4m", modifiers: Underline},
		// Mixing foreground and background
		{name: "red fg black bg    ", text: "this line", escape: "\x1b[31;40m",
			foreground: Red, background: Black},
		{name: "green fg magenta bg", text: "this line", escape: "\x1b[32;45m",
			foreground: Green, background: Magenta},
		{name: "white fg cyan bg   ", text: "this line", escape: "\x1b[37;46m",
			foreground: White, background: Cyan},
		// Mixing foreground and background with multiple modifiers.
		// The implementation by default puts the modifiers in ascending order.
		{name: "fg + bg and Bold", text: "this line", escape: "\x1b[30;42;1m",
			foreground: Black, background: Green, modifiers: Bold},
		{name: "fg + bg and Italic", text: "this line", escape: "\x1b[30;42;3m",
			foreground: Black, background: Green, modifiers: Italic},
		{name: "fg + bg and Underline", text: "this line", escape: "\x1b[30;42;4m",
			foreground: Black, background: Green, modifiers: Underline},
		{name: "fg + bg and Bold + Italic", text: "this line", escape: "\x1b[30;42;1;3m",
			foreground: Black, background: Green, modifiers: Bold | Italic},
		{name: "fg + bg and Bold + Underline", text: "this line", escape: "\x1b[30;42;1;4m",
			foreground: Black, background: Green, modifiers: Bold | Underline},
		{name: "fg + bg and Italic + Underline", text: "this line", escape: "\x1b[30;42;3;4m",
			foreground: Black, background: Green, modifiers: Italic | Underline},
		{name: "fg + bg and Underline + Italic", text: "this line", escape: "\x1b[30;42;3;4m",
			foreground: Black, background: Green, modifiers: Underline | Italic},
		{name: "fg + bg and Bold + Italic + Underline", text: "this line", escape: "\x1b[30;42;1;3;4m",
			foreground: Black, background: Green, modifiers: Bold | Italic | Underline},
	}

	checkFormat := func(t *testing.T, fg, bg, mo uint, text, escape string) {
		t.Helper()
		got := NewFormatter(fg, bg, mo).Format(text)
		want := fmt.Sprintf("%s%s%s", escape, text, endReset)
		if got != want {
			t.Errorf("got %#v - want %#v", got, want)
		}
	}

	for _, f := range formatTests {
		t.Run(f.name, func(t *testing.T) {
			checkFormat(t, f.foreground, f.background, f.modifiers, f.text, f.escape)
		})
	}

}
