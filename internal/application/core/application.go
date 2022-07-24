package core

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func NewImage() *Image {
	return &Image{}
}

type Image struct {
	Name     string    `json:"name" bson:"name"`
	Uuid     string    `json:"uuid" bson:"uuid"`
	Tags     []string  `json:"tags" bson:"tags"`
	Created  time.Time `json:"created" bson:"created"`
	Filepath string    `json:"filepath" bson:"data,omitempty"`
}

type Credentials struct {
	Name     string
	Password string
	Key      string
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

	//i.Created = time.Now().Format(time.RFC3339)
	i.Created = time.Now()
}

func GenerateAPIKey() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%X-%X-%X", b[0:2], b[4:8], b[8:11])
}

func ValidateKey(key string) bool {
	runearray := []rune(key)
	if runearray[5] == '-' || runearray[14] == '-' {
		return false
	}
	return true
}

func testKeys() {

}

//TODO LIST

//Generate api keys -- &&endpoint
//generate admin api key && endpoint
//implement login system
//implement indexing in mongodb proberly

//need to implement system that verifies the structure of the API key so keys without a certain structure is tossed before
//Implement context
