package adapter

import (
	"fmt"
	"os"
	"testing"
)

func TestAddFile(t *testing.T) {

}

func TestDeleteFile(t *testing.T) {

}

func FileDirectory() {
	os.Chdir("../image-folder")
	newDir, err := os.Getwd()
	if err != nil {
	}
	fmt.Printf("Current Working Direcotry: %s\n", newDir)
}
