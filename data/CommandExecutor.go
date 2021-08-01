package data

import "os/exec"

func executeCommand(command string) (string, error) {
	out, err := exec.Command(command).Output()
	return string(out), err
}
