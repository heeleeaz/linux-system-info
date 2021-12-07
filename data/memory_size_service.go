package data

import (
	"fmt"
	"regexp"
)

const FetchMemorySizeCommand = "grep -i MemTotal /proc/meminfo"

const MemorySizeFormatMatchCheckRegex = "(?:\\w+:)\\s+(?:\\d+)\\s+(kB)"
const MemorySizeExtractorRegex = "\\d+"

func MemorySize() (string, error) {
	if out, err := commandExecutor.executeCommand(FetchMemorySizeCommand); err != nil {
		return "", err
	} else {
		return formatMemorySize(out)
	}
}

func formatMemorySize(output string) (string, error) {
	if match, err := regexp.MatchString(MemorySizeFormatMatchCheckRegex, output); match {
		return regexp.MustCompile(MemorySizeExtractorRegex).FindString(output), nil
	} else {
		return "", &FormatOutputError{fmt.Sprintf("error formatting output: %s", output), err}
	}
}
