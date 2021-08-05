package data

var commandExecutor CommandExecutor = NewBashCommandExecutor()

type MockCommandExecutor struct {
	output string
	err    error
}

func (r MockCommandExecutor) executeCommand(command string) (string, error) {
	return r.output, r.err
}
