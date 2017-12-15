package bamboo

type simpleError struct {
	message string
}

func (e *simpleError) Error() string {
	return e.message
}

func newSimpleError(message string) *simpleError {
	return &simpleError{
		message: message,
	}
}

type errBadURL struct {
	message string
}

func (e *errBadURL) Error() string {
	return e.message
}

func newErrBadURL(message string) *errBadURL {
	return &errBadURL{
		message: message,
	}
}
