package adapter

import (
	"fmt"
	"os"
	"testing"
)

func TestFindFile(t *testing.T) {
	uuid := "testimage"
	a := NewFileAdapter()

	file, err := a.FindFile(uuid)
	if err != nil {
		t.Fail()
	}
	//FileDirectory()
	if file == nil {
		t.Fail()
	}
	t.Log(file)

}

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
