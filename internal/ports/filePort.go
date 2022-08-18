package ports

import (
	"bytes"
	"context"

	"github.com/dietzy1/imageAPI/internal/application/core"
)

// Implements the filedb methods
type FilePort interface {
	UploadFile(ctx context.Context, image core.Image, buf *bytes.Buffer) (string, error)
	DeleteFile(ctx context.Context, image core.Image) error
	UpdateFile(ctx context.Context, image core.Image) error
}
