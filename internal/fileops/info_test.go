package fileops

import (
	"os"
	"testing"
)

func TestDisplayFileInfo(t *testing.T) {
	// Create a test file
	filename := "testfile.txt"
	file, err := os.Create(filename)
	if err != nil {
		t.Fatalf("Error creating test file: %s", err)
	}
	defer func() {
		file.Close()
		err := os.Remove(filename)
		if err != nil {
			t.Errorf("Error cleaning up test file: %s", err)
		}
	}()

	// Test displaying file info
	err = DisplayFileInfo(filename)
	if err != nil {
		t.Errorf("Error displaying file info: %s", err)
	}
}
