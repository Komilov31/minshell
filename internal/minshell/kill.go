package minshell

import (
	"fmt"
	"os"
	"strconv"
)

func (ms *MinShell) kill(pid int) {
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = process.Kill()
	if err != nil {
		fmt.Printf("minshell: kill: (%d) - No such process\n", pid)
		return
	}
}

func (ms *MinShell) processKill(command Command) {
	if len(command.Args) < 1 {
		fmt.Println("kill: usage: kill pid | jobspec")
	}

	for _, p := range command.Args {
		pid, err := strconv.Atoi(p)
		if err != nil {
			fmt.Printf("minshell: kill: %s: arguments must be process or job ID\n", p)
			return
		}
		ms.kill(pid)
	}

}
