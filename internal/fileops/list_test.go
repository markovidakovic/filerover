package fileops

import (
	"os"
	"testing"
)

func TestListDirectoryContents(t *testing.T) {
	dirname := "testdir"
	err := os.Mkdir(dirname, 0755)
	if err != nil {
		t.Fatalf("Error creating test directory: %s", err)
	}
	defer os.RemoveAll(dirname)

	// Test listing directory contents
	err = ListDirectoryContents(dirname)
	if err != nil {
		t.Errorf("Error listing directory contents: %s", err)
	}
}
