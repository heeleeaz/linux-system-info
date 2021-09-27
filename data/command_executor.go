package data

type CommandExecutor interface {
	executeCommand(command string) (string, error)
}
