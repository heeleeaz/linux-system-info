package data

import (
	"fmt"
	"strings"
)

const FetchMemorySizeCommand = "grep -i MemTotal /proc/meminfo"

func MemorySize() (string, error) {
	if out, err := executeCommand(FetchMemorySizeCommand); err != nil {
		return "", err
	} else {
		return formatMemorySize(out)
	}
}

func formatMemorySize(output string) (string, error) {
	stringSplit := strings.Split(output, ":")
	if len(stringSplit) == 2 {
		memorySize := strings.TrimSpace(stringSplit[1])
		return memorySize, nil
	} else {
		return "", &FormatOutputError{fmt.Sprintf("error formatting output: %s", output), nil}
	}
}
