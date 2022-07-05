package core

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func NewImage() *Image {
	return &Image{}
}

type Image struct {
	Name    string    `json:"name" bson:"name"`
	Uuid    string    `json:"uuid" bson:"uuid"`
	Tags    []string  `json:"tags" bson:"tags"`
	Created time.Time `json:"created" bson:"created"`
	//	Data    multipart.File `json:"data,omitempty"  bson:"data,omitempty"`
}

func (i Image) Validate(image Image) error {
	if i.Name == "" {
		fmt.Println("returning")
		return errors.New("")
	}
	if len(i.Tags) < 0 {
		fmt.Println("returning tags")
		return errors.New("")
	}
	fmt.Println("Validation ok")
	return nil
}

func (i *Image) NewUUID() {
	i.Uuid = uuid.New().String()
}

func (i *Image) SetTime() {
	i.Created = time.Now()
}

//Might need to check if the pointers of the fields are equal to nil to find out if they are filled out.

//prolly need to define the image struct in here
