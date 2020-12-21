# Pintor

Pintor is a small package to colorize text output, for use in your Command Line applications.
Pintor uses ANSI escape sequences to modify text.

## Installation

To install Pintor, 

1. Install Pintor

```sh
$ go get -u github.com/jbaxx/pintor
```

2. Import it in your code

```sh
import "github.com/jbaxx/pintor"
```

## Quick start
Start by creating a formatter object passing the foreground and background colors, and modifiers as parameters.
Then apply the format to a string, which will be properly ANSI escaped to print the desired colors.

```go
package main

import (
	"fmt"

	"github.com/jbaxx/pintor"
)

func main() {
	red := pintor.NewFormatter(pintor.Red, 0, 0)
	textColoredRed := red.Format("This text is red")
	fmt.Println(textColoredRed)

	blue := pintor.NewFormatter(pintor.Blue, 0, 0)
	fmt.Println(blue.Format("This text is blue"))

	fmt.Println(red.Format("This text is also red"))

	whiteMagentaBold := pintor.NewFormatter(pintor.White, pintor.Magenta, pintor.Bold)
	fmt.Println(whiteMagentaBold.Format("This text is bold, colored white with magenta foreground"))

	cyanBoldItalic := pintor.NewFormatter(pintor.Cyan, pintor.Default, pintor.Bold|pintor.Italic)
	fmt.Println(cyanBoldItalic.Format("This text is bold and italic, colored cyan, and uses the default background"))

}
```
