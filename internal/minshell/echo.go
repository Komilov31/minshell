package minshell

import "fmt"

func (sh *MinShell) echo(args ...string) {
	for i, arg := range args {
		if i == len(args)-1 {
			fmt.Print(arg)
		} else {
			fmt.Print(arg + " ")
		}
	}
	fmt.Println()
}
