package ports

import "mime/multipart"

//Implements the filedb methods
type FilePort interface {
	AddFile(uuid string, data multipart.File) error
	DeleteFile(uuid string) error
}
