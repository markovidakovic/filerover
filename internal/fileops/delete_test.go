package fileops

import (
	"os"
	"testing"
)

func TestDeleteFile(t *testing.T) {
	filename := "testfile.txt"

	// Create a file for deletion
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Error creating test file: %s", err)
	}
	defer file.Close()

	// Test deleting a file
	err = DeleteFile(filename)
	if err != nil {
		t.Errorf("Error deleting file: %s", err)
	}

	// Test deleting a non-existing file (should return an error)
	err = DeleteFile(filename)
	if err == nil {
		t.Errorf("Expected an error, file does not exist")
	}
}

func TestDeleteDirectory(t *testing.T) {
	dirname := "testdir"

	// Create a directory for deletion
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		t.Fatalf("Error creating test directory: %s", err)
	}
	defer os.RemoveAll(dirname)

	// Test deleting a directory
	err = DeleteDirectory(dirname)
	if err != nil {
		t.Errorf("Error deleting directory: %s", err)
	}

	// Test deleting a non-existent directory (should return an error)
	err = DeleteDirectory(dirname)
	if err == nil {
		t.Errorf("Expected an error, directory does not exist")
	}
}
