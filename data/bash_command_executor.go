package data

import "os/exec"

const ShellType = "bash"

type BashCommandExecutor struct{}

func NewBashCommandExecutor() BashCommandExecutor {
	return BashCommandExecutor{}
}

func (r BashCommandExecutor) executeCommand(command string) (string, error) {
	out, err := exec.Command(ShellType, "-c", command).Output()
	return string(out), err
}
