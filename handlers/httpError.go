package handlers

type HttpError struct {
	message  string
	httpCode int
}

func (err HttpError) Error() string {
	return err.message
}

func newHttpError(message string, httpCode int) HttpError {
	return HttpError{message: message, httpCode: httpCode}
}
