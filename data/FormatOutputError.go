package data

type FormatOutputError struct {
	message string
	err     error
}

func (e *FormatOutputError) Error() string {
	return e.message + " : " + e.err.Error()
}
