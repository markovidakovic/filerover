package fileops

import (
	"fmt"
	"os"
	"time"
)

func DisplayFileInfo(filename string) error {
	// Get file information
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return fmt.Errorf("could not get file information '%s'", err)
	}

	// Extract file details
	fmt.Printf("Information for file: %s\n", filename)
	fmt.Println("-----------------------------------")
	fmt.Printf("Name: %s\n", fileInfo.Name())
	fmt.Printf("Size: %d bytes\n", fileInfo.Size())
	fmt.Printf("Permissions: %s\n", fileInfo.Mode())
	fmt.Printf("Last modified: %s\n", fileInfo.ModTime().Format(time.RFC3339))
	fmt.Printf("Is directory: %t\n", fileInfo.IsDir())

	return nil
}
