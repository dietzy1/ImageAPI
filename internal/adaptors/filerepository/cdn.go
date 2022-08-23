package filerepository

import (
	//"image/jpeg"

	//"image/jpeg"

	"bytes"
	"context"
	"fmt"

	"io"
	"os"
	"strings"

	"github.com/dietzy1/imageAPI/internal/application/core"
	"github.com/imagekit-developer/imagekit-go"
	"github.com/imagekit-developer/imagekit-go/api/media"
	"github.com/imagekit-developer/imagekit-go/api/uploader"
)

type FileAdapter struct {
	client *imagekit.ImageKit
}

func NewImageKitClientAdapter() (*FileAdapter, error) {

	client := imagekit.NewFromParams(imagekit.NewParams{
		PrivateKey:  os.Getenv("PRIVATE_KEY"),
		PublicKey:   os.Getenv("PUBLIC_KEY"),
		UrlEndpoint: os.Getenv("URL_ENDPOINT"),
	})
	a := &FileAdapter{client: client}
	return a, nil
}

func (f *FileAdapter) UploadFile(ctx context.Context, image core.Image, buf *bytes.Buffer) (string, error) {
	tags := strings.Join(image.Tags, ",")
	tags = strings.TrimSpace(tags)

	params := uploader.UploadParam{
		FileName:          image.Uuid + ".jpg",
		UseUniqueFileName: newFalse(),
		Tags:              tags,
		Folder:            "/pepes/",
		IsPrivateFile:     newFalse(),
		ResponseFields:    "filepath",
	}
	res, err := f.client.Uploader.Upload(ctx, io.ByteReader(buf), params)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return res.Data.Url, nil
}

func (f *FileAdapter) DeleteFile(ctx context.Context, uuid string) error {
	fileid, err := f.GetFile(ctx, uuid)
	if err != nil {
		return err
	}
	_, err = f.client.Media.DeleteFile(ctx, fileid)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileAdapter) GetFile(ctx context.Context, uuid string) (string, error) {
	query := fmt.Sprintf(`name = "%s"`, uuid+".jpg")
	res, err := f.client.Media.Files(ctx, media.FilesParam{
		SearchQuery: query,
	})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return res.Data[0].FileId, nil
}

func (f *FileAdapter) UpdateFile(ctx context.Context, image core.Image) error {
	fileid, err := f.GetFile(ctx, image.Uuid)
	if err != nil {
		return err
	}
	_, err = f.client.Media.UpdateFile(ctx, fileid, media.UpdateFileParam{
		Tags: image.Tags,
	})
	if err != nil {
		return err
	}
	return nil
}

func newFalse() *bool {
	b := false
	return &b
}
