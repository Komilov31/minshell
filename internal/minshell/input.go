package minshell

import (
	"fmt"
	"strings"
)

// структура для хранения информации о команде
type Command struct {
	Name string
	Args []string
}

func parseArguments(nextLine string) ([]Command, error) {
	nextLine = strings.TrimSuffix(nextLine, "\n")
	cmds := strings.Split(nextLine, " | ")

	var commands []Command
	for _, cmd := range cmds {
		command := Command{}
		cmdSplited := strings.Split(cmd, " ")
		if len(cmdSplited) == 0 {
			return nil, fmt.Errorf("minshell: syntax error")
		}

		command.Name = cmdSplited[0]
		if len(cmdSplited) > 1 {
			command.Args = cmdSplited[1:]
		}

		commands = append(commands, command)
	}

	return commands, nil
}
