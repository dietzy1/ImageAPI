package core

import (
	"crypto/rand"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
	Username     string    `json:"username" bson:"username"`
	Passwordhash string    `json:"passwordhash" bson:"passwordhash"`
	Key          string    `json:"key" bson:"key"`
	Created      time.Time `json:"created" bson:"created"`
	Role         int       `json:"role" bson:"role"`
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

func (c Credentials) Validate(crreds Credentials) error {
	if c.Username == "" {
		return errors.New("No username")
	}
	if c.Passwordhash == "" {
		return errors.New("No password")
	}
	if c.Key == "" {
		return errors.New("No key")
	}
	return nil
}

func (i *Image) NewUUID() {
	i.Uuid = uuid.New().String()
}

//can add a generic method for setting time
func (i *Image) SetTime() {

	//i.Created = time.Now().Format(time.RFC3339)
	i.Created = time.Now()
}

func (c *Credentials) SetTime() {
	c.Created = time.Now()
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

//method that hashes a password
func (c *Credentials) Hash(creds Credentials) Credentials {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(creds.Passwordhash), 8)
	if err != nil {
		fmt.Println("Error hashing password")
	}
	c.Passwordhash = string(hashedPassword)
	return *c
}

//TODO LIST

//Generate api keys -- &&endpoint
//generate admin api key && endpoint
//implement login system
//implement indexing in mongodb proberly

//need to implement system that verifies the structure of the API key so keys without a certain structure is tossed before
//Implement context
