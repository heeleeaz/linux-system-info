package data

import (
	"fmt"
	"strconv"
	"strings"
)

const CpuCountCommand = "nproc"

func CpuCount() (int, error) {
	if out, err := commandExecutor.executeCommand(CpuCountCommand); err != nil {
		return -1, err
	} else {
		return formatCpuCount(out)
	}
}

func formatCpuCount(output string) (int, error) {
	trimmed := strings.TrimSpace(output)
	if count, err := strconv.Atoi(trimmed); err != nil {
		return -1, &FormatOutputError{fmt.Sprintf("cannot convert %s to %T", output, count), err}
	} else {
		return count, nil
	}
}
