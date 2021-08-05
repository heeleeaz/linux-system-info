package data

import (
	"fmt"
	"testing"
)

func TestGpuCount(t *testing.T) {
	params := map[string]interface{}{"03:03.1": 1, "03:01.1\n02:11.12": 2, "some=sks": &FormatOutputError{}}

	for given, expected := range params {
		t.Run("when given command output="+given, func(t *testing.T) {
			commandExecutor = MockCommandExecutor{given, nil}
			if output, err := GpuCount(); err != nil {
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
