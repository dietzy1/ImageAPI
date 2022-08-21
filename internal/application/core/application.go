package core

import (
	"crypto/rand"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"strings"
	"time"

	_ "golang.org/x/image/webp"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type Image struct {
	Name     string    `json:"name" bson:"name"`
	Uuid     string    `json:"uuid" bson:"uuid"`
	Tags     []string  `json:"tags" bson:"tags"`
	Created  time.Time `json:"created" bson:"created"`
	Filepath string    `json:"filepath" bson:"filepath"`
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
	if len(i.Tags) == 0 {
		fmt.Println("returning tags")
		return errors.New("")
	}
	fmt.Println("Validation ok")
	return nil
}

func (c Credentials) Validate(crreds Credentials) error {
	if c.Username == "" {
		return errors.New("username is required")
	}
	if c.Passwordhash == "" {
		return errors.New("password is required")
	}
	if c.Key == "" {
		return errors.New("key is required")
	}
	return nil
}

func (i *Image) NewUUID() string {
	return uuid.New().String()
}

// can add a generic method for setting time
func (i *Image) SetTime() time.Time {
	//i.Created = time.Now().Format(time.RFC3339)
	return time.Now()
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

// Input r.Form.Get("Tags")
func Split(input string) []string {
	input = strings.TrimSpace(input)
	return strings.Split(input, ",")
}

func (c *Credentials) Hash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		fmt.Println("Error hashing password")
	}
	return string(hashedPassword)
}

func (c *Credentials) CompareHash(storedpassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedpassword), []byte(password))
	return err == nil
}

func ConvertToJPEG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: 95})
}

func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
