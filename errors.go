package bamboo

type simpleError struct {
	message string
}

func (e *simpleError) Error() string {
	return e.message
}
