package minshell

import (
	"fmt"
	"io"
	"os/exec"
)

const (
	enoughProcessForPipeLine = 2
)

func (ms *MinShell) pipeline(commands []Command) {
	cmds := make([]*exec.Cmd, len(commands))
	writers := []*io.PipeWriter{}

	if len(commands) < enoughProcessForPipeLine {
		return
	}

	for i, command := range commands {
		cmds[i] = exec.Command(command.Name, command.Args...)

		if i > 0 {
			pipeReader, pipeWriter := io.Pipe()
			cmds[i-1].Stdout = pipeWriter
			cmds[i].Stdin = pipeReader

			writers = append(writers, pipeWriter)
		}

	}

	//запускаем команды и также горутины, которые при завершении команды закрою stdout
	for i := 0; i < len(cmds)-1; i++ {
		cmd := cmds[i]

		err := cmd.Start()
		if err != nil {
			fmt.Println(err)
			return
		}

		if i < len(writers) {
			go func() {
				_ = cmd.Wait()
				_ = writers[i].Close()
			}()
		}
	}

	lastCommand := cmds[len(cmds)-1]
	output, err := lastCommand.Output()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(output))
}
