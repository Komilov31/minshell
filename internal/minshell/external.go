// package minshell

// import (
// 	"fmt"
// 	"os"
// 	"os/exec"
// )

// func (sh *MinShell) external(command Command) {
// 	cmd := exec.Command(command.Name, command.Args...)
// 	cmd.Stdin = os.Stdin

// 	err := cmd.Start()
// 	if err != nil {
// 		fmt.Printf("%s: command not found\n", command.Name)
// 		return
// 	}
// 	fmt.Print(string(23))
// }

package minshell

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"
)

func (ms *MinShell) external(command Command) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	var cmd *exec.Cmd
	if command.Name == "ls" || command.Name == "dir" {
		if len(command.Args) == 0 {
			cmd = exec.Command(command.Name, ms.currentDirectory)
		} else {
			cmd = exec.Command(command.Name, command.Args...)
		}
	} else {
		cmd = exec.Command(command.Name, command.Args...)
	}

	go func() {
		for {
			<-sigs
			if cmd != nil && cmd.Process != nil {
				_ = cmd.Process.Signal(syscall.SIGINT)
			}
		}
	}()

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); !ok {
			fmt.Printf("%s: command not found\n", command.Name)
		}
		return
	}
	cmd = nil
}
