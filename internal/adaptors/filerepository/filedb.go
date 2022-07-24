package filerepository

import (
	//"image/jpeg"

	//"image/jpeg"

	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"os"

	_ "golang.org/x/image/webp"
)

type FileAdapter struct {
}

func NewFileAdapter() *FileAdapter {
	return &FileAdapter{}
}

func (f *FileAdapter) AddFile(uuid string, data multipart.File) error {
	os.Chdir(os.Getenv("FILE_DIR"))
	file, err := os.OpenFile(uuid+".jpg", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("unable to open the data")
		return err
	}
	otp := jpeg.Options{
		Quality: 80,
	}
	io.Copy(file, data)
	target := image.NewRGBA(image.Rect(0, 0, 800, 800))
	err = jpeg.Encode(file, target, &otp)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileAdapter) DeleteFile(uuid string) error {
	os.Chdir(os.Getenv("FILE_DIR"))
	err := os.Remove(uuid + ".jpg")
	if err != nil {
		return err
	}
	return nil
}
