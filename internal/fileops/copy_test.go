package fileops

import (
	"os"
	"testing"
)

func TestCopyFileOrDirectory(t *testing.T) {
	// Create file for copying
	srcFile := "source.txt"
	destFile := "destination.txt"
	file, err := os.Create(srcFile)
	if err != nil {
		t.Fatalf("Error creating test file: %s", err)
	}
	defer func() {
		file.Close()
		err := os.Remove(srcFile)
		if err != nil {
			t.Errorf("Error cleaning up test file: %s", err)
		}
	}()

	// Test copying a file
	err = CopyFileOrDirectory(srcFile, destFile)
	if err != nil {
		t.Fatalf("Error copying file: %s", err)
	}

	// Clean up destination file after the test
	defer os.Remove(destFile)
}
