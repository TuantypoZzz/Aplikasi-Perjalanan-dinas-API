package exception

type UnsupportedMediaError struct {
	Message string
}

func (unsupportedMediaError UnsupportedMediaError) Error() string {
	return unsupportedMediaError.Message
}
