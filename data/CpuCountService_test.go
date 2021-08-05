package data

import (
	"fmt"
	"testing"
)

func TestCpuCount(t *testing.T) {
	params := map[string]interface{}{"4": 4, "6   \t\n": 6, "192n": &FormatOutputError{}}

	for given, expected := range params {
		t.Run("when given command output="+given, func(t *testing.T) {
			commandExecutor = MockCommandExecutor{given, nil}
			if output, err := CpuCount(); err != nil {
				if fmt.Sprintf("%T", err) != fmt.Sprintf("%T", expected) {
					t.Errorf("expected: %T, actual: %T", expected, err)
				}
			} else {
				if expected != output {
					t.Errorf("expected: %s, actual: %d", expected, output)
				}
			}
		})
	}
}
