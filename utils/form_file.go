package utils

import "os"

type FormFile struct {
	Name        string
	ContentType string
	Content     *os.File
}
