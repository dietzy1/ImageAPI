package ports

import (
	"context"
	"mime/multipart"
)

// Implements the filedb methods
type FilePort interface {
	AddFile(ctx context.Context, uuid string, data multipart.File) error
	DeleteFile(ctx context.Context, uuid string) error
}
