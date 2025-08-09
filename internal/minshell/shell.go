package minshell

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type MinShell struct {
	input            *os.File
	currentDirectory string
}

func New() *MinShell {

	userHomeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not find home dir of user", err)
	}

	minshell := &MinShell{
		input:            os.Stdin,
		currentDirectory: userHomeDir,
	}

	return minshell
}

func (ms *MinShell) ProcessProgram() {
	reader := bufio.NewReader(ms.input)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)

	for {
		select {
		case <-sigs:
			fmt.Println()
			continue
		default:
			// fmt.Print(ms.currentDirectory + ":$ ")
			nextLine, err := reader.ReadString('\n')
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}

			if err == io.EOF {
				return
			}

			commands, err := parseArguments(nextLine)
			if err != nil {
				log.Fatal(err)
			}

			isPipeLine := strings.Contains(nextLine, "|")
			if isPipeLine {
				ms.pipeline(commands)
				continue
			}

			for _, command := range commands {
				ms.callCommand(command)
			}
		}
	}
}
