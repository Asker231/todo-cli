package menu

import (
	"fmt"
	"github.com/fatih/color"
	term "github.com/nsf/termbox-go"
)

func reset() {
	term.Sync()
}

func Select(sel int) {
	Menu(sel)
loop:
	for {
		ev := term.PollEvent()
		switch ev.Type {
		case term.EventKey:
			switch ev.Key {
			case term.KeyArrowDown:
				if sel > 2 {
					sel = 2
				}
				sel += 1
				reset()
				Menu(sel)
			case term.KeyArrowUp:
				if sel <= 1 {
					sel += 1
				}
				sel -= 1
				reset()
				Menu(sel)
			default:
				reset()
				break loop
			}
		}
	}
}

func Menu(sel int) {
	if sel == 1 {
		color.Cyan("Add todo: ➕\n")
		fmt.Print("\n")
		fmt.Print("Delete todo:🗑️\n")
		fmt.Print("\n")
		fmt.Print("List:📋\n")
	}

	if sel == 2 {
		fmt.Print("Add todo: ➕\n")
		fmt.Print("\n")
		color.Cyan("Delete todo:🗑️\n")
		fmt.Print("\n")
		fmt.Print("List:📋\n")
	}

	if sel == 3 {
		fmt.Print("Add todo: ➕\n")
		fmt.Print("\n")
		fmt.Print("Delete todo:🗑️\n")
		fmt.Print("\n")
		color.Cyan("List:📋\n")
	}
}
