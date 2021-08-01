package domain

type SystemError struct {
	message string
	err     error
}

func (e *SystemError) Error() string {
	return e.message + " : " + e.err.Error()
}
