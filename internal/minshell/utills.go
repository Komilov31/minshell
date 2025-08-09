package minshell

import (
	"os"
)

// функция для проверки существует ли папка указанная по path
func dirExists(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}
		return false, err
	}

	return info.IsDir(), nil
}

func (ms *MinShell) callCommand(command Command) {
	switch command.Name {
	case "cd":
		ms.processCd(command)
	case "pwd":
		ms.pwd()
	case "echo":
		ms.echo(command.Args...)
	case "kill":
		ms.processKill(command)
	case "ps":
		ms.ps()
	default:
		ms.external(command)
	}
}
