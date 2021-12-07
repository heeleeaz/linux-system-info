package data

import (
	"fmt"
	"testing"
)

func TestMemorySize(t *testing.T) {
	params := map[string]interface{}{"MemTotal: \t3221344 kB": "3221344", "bash not found": &FormatOutputError{}}

	for given, expected := range params {
		t.Run("when given command output="+given, func(t *testing.T) {
			commandExecutor = MockCommandExecutor{given, nil}
			if output, err := MemorySize(); err != nil {
				if fmt.Sprintf("%T", err) != fmt.Sprintf("%T", expected) {
					t.Errorf("expected: %T, actual: %T", expected, err)
				}
			} else {
				if expected != output {
					t.Errorf("expected: %s, actual: %s", expected, output)
				}
			}
		})
	}
}
