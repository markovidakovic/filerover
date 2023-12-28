package fileops

import (
	"os"
	"testing"
)

func TestMoveFileOrDirectory(t *testing.T) {
	// Create a file for moving
	srcFile := "source.txt"
	destFile := "destination.txt"
	file, err := os.Create(srcFile)
	if err != nil {
		t.Fatalf("Error creating test file: %s", err)
	}
	defer file.Close()

	// Test moving a file
	err = MoveFileOrDirectory(srcFile, destFile)
	if err != nil {
		t.Errorf("Error moving file: %s", err)
	}

	// Clean up destination file after the test
	defer os.Remove(destFile)
}
