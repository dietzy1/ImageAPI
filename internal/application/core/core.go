package core

import (
	"bytes"
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
	"github.com/vitali-fedulov/imagehash"
	"github.com/vitali-fedulov/images4"
	"golang.org/x/crypto/bcrypt"
)

// Used for image comparison
const (
	// Recommended hyper-space parameters for initial trials.
	epsPct     = 0.25
	numBuckets = 4
)

type Image struct {
	Name       string   `json:"name" bson:"name"`
	Uuid       string   `json:"uuid" bson:"uuid"`
	Tags       []string `json:"tags" bson:"tags"`
	Created_At string   `json:"created_at" bson:"created_at"`
	Filepath   string   `json:"filepath" bson:"filepath"`
	Filesize   int64    `json:"filesize" bson:"filesize"`
	Hash       uint64   `json:"hashset" bson:"hashset"`
}

type Credentials struct {
	Username     string `json:"username" bson:"username"`
	Passwordhash string `json:"passwordhash" bson:"passwordhash"`
	Key          string `json:"key" bson:"key"`
	Created_At   string `json:"created_at" bson:"created_at"`
	Role         int    `json:"role" bson:"role"`
}

// Simple validation againt the image struct that checks if name and tags are empty
func (i Image) Validate(image Image) error {
	if i.Name == "" {
		fmt.Println("returning name")
		return errors.New("empty name")
	}
	if len(i.Tags) == 0 {
		fmt.Println("returning tags")
		return errors.New("empty tags")
	}
	return nil
}

// Skip error handling to improve code structure
// Hashes the image and returns the hash // HashSet is the image that is being queried against the centralHashes
func (i Image) HashSet(buf *bytes.Buffer) uint64 {
	img, _, err := image.Decode(buf)
	if err != nil {
		return 0
	}
	icon := images4.Icon(img)

	return imagehash.HashSet(
		icon, imagehash.HyperPoints10, epsPct, numBuckets)[0]
}

// Hashes the image and returns the hash //CentralHash is all prior images that are being compared against
func (i Image) CentralHash(buf *bytes.Buffer) (uint64, error) {
	img, _, err := image.Decode(buf)
	if err != nil {
		return 0, err
	}
	icon := images4.Icon(img)

	return imagehash.CentralHash(
		icon, imagehash.HyperPoints10, epsPct, numBuckets), nil
}

func (i Image) FileSize(buf *bytes.Buffer) int64 {
	return int64(buf.Len())
}

// spawn 5 goroutines to use the function CompareImage
func batchProcess(image []Image) (*Image, bool) {
	return nil, true
}

// Simple validation againt the credentials struct that checks if username, password and key are empty strings
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

// Converts an error to a string
func Errconv(err error) string {
	return fmt.Sprintf("%s", err)
}

// Returns a newly generated uuid string
func (i *Image) NewUUID() string {
	return uuid.New().String()
}

// Returns time.now as a string in the format RFC1123
func (i *Image) SetTime() string {
	return time.Now().Format("RFC1123")
}

// Generates a custom API key.
func GenerateAPIKey() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%X-%X-%X", b[0:2], b[4:8], b[8:11])
}

// Initial validation to deter keys with wrong format
func ValidateKey(key string) bool {
	runearray := []rune(key)
	if runearray[5] == '-' || runearray[14] == '-' {
		return false
	}
	return true
}

// Input r.Form.Get("Tags")
// Splits a single string into an array of lowercase letters without any whitespace
func Split(input string) []string {
	return strings.Split(strings.ReplaceAll(strings.ToLower(input), " ", ""), ",")
}

func (c *Credentials) Hash(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		fmt.Println("Error hashing password")
	}
	return string(hashedPassword)
}

// Compares password from mongodb with input password
func (c *Credentials) CompareHash(storedpassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedpassword), []byte(password))
	return err == nil
}

// Accepts formats of webp, png, jpeg and gif
func ConvertToJPEG(w io.Writer, r io.Reader) error {
	img, _, err := image.Decode(r)
	if err != nil {
		return err
	}
	return jpeg.Encode(w, img, &jpeg.Options{Quality: 95})
}

// Validation of uuid
func IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}
