package minshell

import "fmt"

func (sh *MinShell) pwd() {
	fmt.Println(sh.currentDirectory)
}
