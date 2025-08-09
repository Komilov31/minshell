package minshell

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func (ms *MinShell) processCd(command Command) {
	if len(command.Args) > 1 {
		fmt.Println("minshell: cd: too many arguments")
		return
	}

	relativePath, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	if len(command.Args) == 0 {
		ms.currentDirectory = relativePath
		return
	}

	ms.cd(command.Args[0])
}

func (ms *MinShell) cd(relativePath string) {
	newCurrentDir := filepath.Join(ms.currentDirectory, relativePath)

	exists, err := dirExists(newCurrentDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	if !exists {
		fmt.Printf("minshell: cd: %s: No such file or directory\n", relativePath)
		return
	}

	ms.currentDirectory = filepath.Clean(newCurrentDir)
}
