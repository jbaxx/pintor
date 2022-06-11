package main

import (
	"fmt"
	"os"

	"github.com/jbaxx/pintor"
)

var colores = []uint{
	pintor.Default,
	pintor.Black,
	pintor.Red,
	pintor.Green,
	pintor.Yellow,
	pintor.Blue,
	pintor.Magenta,
	pintor.Cyan,
	pintor.White,
}

func PrintAllColors(text string, modifier uint) {
	for _, i := range colores {
		for _, j := range colores {
			a := pintor.NewFormatter(i, j, modifier)
			fmt.Printf("%s ", a.Format(text))
		}
		fmt.Printf("\n")
	}
}

func printa() {
	out, _ := os.Stdout.Stat()
	var lol os.FileMode
	lol = lol | os.ModeCharDevice
	if (out.Mode() & os.ModeCharDevice) == 0 {
		fmt.Printf("lol %#v\n", lol)
		fmt.Printf("%#v\n", out.Mode())
		fmt.Println("no terminal")
	}
	fmt.Printf("lol %#v\n", lol)
	fmt.Println("terminal")
	fmt.Printf("%#v\n\n", out.Mode())
}

func main() {
	printa()
	text := "Pizza!"

	fmt.Println(pintor.NewFormatter(0, 0, 0).Format("Normal:"))
	PrintAllColors(text, 0)
	fmt.Println()

	fmt.Println(pintor.NewFormatter(0, 0, pintor.Bold).Format("Bold:"))
	PrintAllColors(text, pintor.Bold)
	fmt.Println()

	fmt.Println(pintor.NewFormatter(0, 0, pintor.Italic).Format("Italic:"))
	PrintAllColors(text, pintor.Italic)
	fmt.Println()

	fmt.Println(pintor.NewFormatter(0, 0, pintor.Underline).Format("Underline:"))
	PrintAllColors(text, pintor.Underline)
	fmt.Println()

	fmt.Println(pintor.NewFormatter(0, 0, pintor.Bold|pintor.Italic|pintor.Underline).Format("Mixing! -> Bold + Italice + Underline:"))
	PrintAllColors(text, pintor.Bold|pintor.Italic|pintor.Underline)
	fmt.Println()

}
