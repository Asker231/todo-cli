package colors

import "github.com/fatih/color"

func Cyan(title string){
	c := color.New(color.FgCyan).Add(color.Bold).Add(color.Underline)
	c.Print(title)
}