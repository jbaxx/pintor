package main

import (
	"fmt"

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

func main() {
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
