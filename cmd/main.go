package main

import (
	"fmt"
	menu "github.com/Asker231/todo-cli.git/internal/menu"
	term "github.com/nsf/termbox-go"
)
func main() {
	sel := 1;
	err := term.Init()
	if err != nil{
		fmt.Println(err)
	}
	defer term.Close()
	menu.Select(sel)
}

