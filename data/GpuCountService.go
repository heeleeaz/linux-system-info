package data

import (
	"regexp"
	"strings"
)

const FetchGpuCountCommand = "lspci | grep -i 'VGA\\|Display' | cut -d\" \" -f 1"
const LspciBusPortValidatorRegex = "(?:\\d{2}\\:){2}\\d{1,2}"

func GpuCount() (int, error) {
	if out, err := executeCommand(FetchGpuCountCommand); err != nil {
		return -1, err
	} else {
		return formatGpuCount(out)
	}
}

func formatGpuCount(output string) (int, *FormatOutputError) {
	trimmed := strings.TrimSpace(output)
	newlineSeperated := regexp.MustCompile(`\s+`).ReplaceAllString(trimmed, "\n")
	splitted := strings.Split(newlineSeperated, "\n")

	matchCount := 0
	for _, port := range splitted {
		matches, _ := regexp.MatchString(LspciBusPortValidatorRegex, port)

		if matches {
			matchCount += 1
		}
	}

	return matchCount, nil
}
