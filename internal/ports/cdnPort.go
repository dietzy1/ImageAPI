package ports

import (
	"bytes"
	"context"

	"github.com/dietzy1/imageAPI/internal/application/core"
)

// Implements the filedb methods
type CdnPort interface {
	UploadFile(ctx context.Context, image core.Image, buf *bytes.Buffer) (string, error)
	DeleteFile(ctx context.Context, uuid string) error
	UpdateFile(ctx context.Context, image core.Image) error
}
