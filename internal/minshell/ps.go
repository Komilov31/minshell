package minshell

import (
	"fmt"

	"github.com/shirou/gopsutil/process"
)

func (ms *MinShell) ps() {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Could not get procceses", err)
		return
	}

	for _, p := range processes {
		name, err := p.Name()
		if err == nil {
			fmt.Printf("PID: %d, Name: %s\n", p.Pid, name)
		}
	}
}
