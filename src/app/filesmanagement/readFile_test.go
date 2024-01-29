package filesmanagement

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	file, _ := os.Create("readfile.txt")
	file.Close()
	_,err :=	ReadFile("readfile.txt")
	if err != nil {
		t.Errorf("Error occurred: %v", err)
	}
	os.Remove("readfile.txt")
	
}

func TestReadFileWithoutPath(t *testing.T) {
	// read file without argument path
	_,err := ReadFile("")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}

func TestReadFileWithWrongPath(t *testing.T) {
	// read file with wrong path
	_,err := ReadFile("wrongpath/readfile.txt")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}

func TestReadFileWithEmptyPath(t *testing.T) {
	// read file with empty path
	_,err := ReadFile("")
	if err == nil {
		t.Errorf("Error occurred: %v", err)
	}

}
