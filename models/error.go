package models

type BlazingDocsError struct {
	Message string
}

func (e BlazingDocsError) Error() string {
	return e.Message
}
