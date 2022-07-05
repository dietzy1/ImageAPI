package ports

import "mime/multipart"

//Implements the filedb methods
type FilePort interface {
	FindFile(uuid string) ([]byte, error)
	AddFile(uuid string, data multipart.File) error
	DeleteFile(uuid string) error
}
