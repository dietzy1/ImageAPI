package adapter

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

//Still need to implement stuff here
func (f *FileAdapter) FindFile(uuid string) ([]byte, error) {
	os.Chdir("/Users/martinvad/go/src/github.com/dietzy1/imageAPI/image-folder")
	file, err := os.ReadFile(uuid + ".jpg")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return file, nil
}

//Potentially dont need to decode the file
/* 	image, _, err := image.Decode(file)
if err != nil {
	return nil, err
}
return image, nil */

func (f *FileAdapter) AddFile(uuid string, data multipart.File) error {

	//Wrong format
	//Whatever input is given
	os.Chdir("/Users/martinvad/go/src/github.com/dietzy1/imageAPI/image-folder")
	file, err := os.OpenFile(uuid+".jpg", os.O_WRONLY|os.O_CREATE, 0666)
	defer file.Close()
	if err != nil {
		fmt.Println("unable to open the data")
		return err
	}
	io.Copy(file, data)

	/* 	input, format, err := image.Decode(file)
	   	if err != nil {
	   		fmt.Println(format)
	   		os.Remove(uuid + ".jpg")
	   		fmt.Println("unable to decode the input")
	   		return err
	   	} */
	//Need to ensure this is the correct path
	/* 	os.Chdir("/Users/martinvad/go/src/github.com/dietzy1/imageAPI/image-folder")
	   	file, err := os.Create(uuid + ".jpg")
	   	if err != nil {
	   		return err
	   	} */

	otp := jpeg.Options{
		Quality: 80,
	}
	target := image.NewRGBA(image.Rect(0, 0, 800, 800))
	err = jpeg.Encode(file, target, &otp)
	if err != nil {
		return err
	}
	return nil
}

func (f *FileAdapter) DeleteFile(uuid string) error {
	os.Chdir("/Users/martinvad/go/src/github.com/dietzy1/imageAPI/image-folder")
	err := os.Remove(uuid + ".jpg")
	if err != nil {
		return err
	}
	return nil
}

//Function to check if there are duplicates among the images
