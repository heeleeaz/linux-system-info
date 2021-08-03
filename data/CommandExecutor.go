package data

import "os/exec"

const SHELL_TYPE = "bash"

func executeCommand(command string) (string, error) {
	out, err := exec.Command(SHELL_TYPE, "-c", command).Output()
	return string(out), err
}
