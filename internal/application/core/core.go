package core

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
	"image"
	_ "image/gif"
	"image/jpeg"
	"log"

	_ "image/png"
	"io"
	"strings"
	"time"

	"golang.org/x/image/draw"
	_ "golang.org/x/image/webp"

	"github.com/buckket/go-blurhash"
	"github.com/google/uuid"
	"github.com/vitali-fedulov/images4"
	"golang.org/x/crypto/bcrypt"
)

type Image struct {
	Title      string        `json:"title" bson:"title"`
	Uuid       string        `json:"uuid" bson:"uuid"`
	Tags       []string      `json:"tags" bson:"tags"`
	Created_At string        `json:"created_at" bson:"created_at"`
	Filepath   string        `json:"filepath" bson:"filepath"`
	Filesize   int64         `json:"filesize" bson:"filesize"`
	Width      int           `json:"width" bson:"width"`
	Height     int           `json:"height" bson:"height"`
	BlurHash   string        `json:"blurhash" bson:"blurhash"`
	Hash       images4.IconT `json:"-" bson:"hash"`
	Elo        float64       `json:"elo" bson:"elo"`
	Owner      string        `json:"owner" bson:"owner"`
}

type Credentials struct {
	Username     string   `json:"username" bson:"username"`
	Passwordhash string   `json:"passwordhash" bson:"passwordhash"`
	Uuid         string   `json:"uuid" bson:"uuid"`
	Key          string   `json:"key" bson:"key"`
	Created_At   string   `json:"created_at" bson:"created_at"`
	OwnedImages  []string `json:"owned_images" bson:"owned_images"`
}

// Simple validation againt the image struct that checks if name and tags are empty
func (i Image) Validate(image Image) error {
	if image.Title == "" {
		return errors.New("empty name")
	}
	if len(image.Tags) == 0 {
		return errors.New("empty tags")
	}
	if image.BlurHash == "" {
		return errors.New("blurhash error")
	}
	return nil
}

func (i Image) FindWidth(img image.Image) int {
	return img.Bounds().Max.X
}

func (i Image) FindHeight(img image.Image) int {
	return img.Bounds().Max.Y
}

// Returns a hashed icon hash of an image
func (i Image) HashSet(img image.Image) images4.IconT {
	return images4.Icon(img)
}

// function that scales down an image.image to 32x32px
func (i Image) ScaleDown(img image.Image) image.Image {
	// Set the desired width and height
	width := 32
	height := 32

	// Create a new image of the desired width and height
	newImage := image.NewRGBA(image.Rect(0, 0, width, height))

	// Scale the original image to fit the new image
	draw.CatmullRom.Scale(newImage, newImage.Bounds(), img, img.Bounds(), draw.Over, nil)

	return newImage
}

// returns the blurhash of an image
func (i Image) BlurHashing(img image.Image) string {
	blurhashed, err := blurhash.Encode(4, 3, img)
	if err != nil {
		log.Println("Error encoding image")
		return ""
	}
	return blurhashed
}

func (i Image) FileSize(buf bytes.Buffer) int64 {
	return int64(buf.Len())
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
	return time.Now().Format(time.RFC1123)
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
