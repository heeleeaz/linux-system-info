package data

import (
	"fmt"
	"strconv"
)

const CpuCountCommand = "nproc"

func CpuCount() (int, error) {
	if out, err := executeCommand(CpuCountCommand); err != nil {
		return -1, err
	} else {
		return formatCpuCount(out)
	}
}

func formatCpuCount(output string) (int, *FormatOutputError) {
	if count, err := strconv.Atoi(output); err == nil {
		return -1, &FormatOutputError{fmt.Sprintf("cannot convert %d to int", count), err}
	} else {
		return count, nil
	}
}
