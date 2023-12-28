package fileops

import (
	"os"
	"testing"
)

func TestCreateFile(t *testing.T) {
	filename := "testfile.txt"

	// Cleanup the file after the test
	defer func() {
		err := os.Remove(filename)
		if err != nil {
			t.Errorf("Error cleaning up test file: %s", err)
		}
	}()

	// Test creating a file
	err := CreateFile(filename)
	if err != nil {
		t.Errorf("Error creating file: %s", err)
	}

	// Test creating the same file again (should return an error)
	err = CreateFile(filename)
	if err == nil {
		t.Errorf("Expected an error, file already exists")
	}
}

func TestCreateDirectory(t *testing.T) {
	dirname := "testdir"

	// Clean up the directory after the test
	defer func() {
		err := os.RemoveAll(dirname)
		if err != nil {
			t.Errorf("Error cleaning up test directory: %s", err)
		}
	}()

	// Test creating a directory
	err := CreateDirectory(dirname)
	if err != nil {
		t.Errorf("Error creating directory: %s", err)
	}

	// Test creating the same directory again (should return an error)
	err = CreateDirectory(dirname)
	if err == nil {
		t.Errorf("Expected an error, directory already exists")
	}
}
