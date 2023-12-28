package fileops

import (
	"fmt"
	"os"
)

func ListDirectoryContents(dirname string) error {
	// Open the directory
	dir, err := os.Open(dirname)
	if err != nil {
		return fmt.Errorf("could not open directory '%s'", err)
	}
	defer dir.Close()

	// Read the contents of the directory
	fileInfos, err := dir.Readdir(-1)
	if err != nil {
		return fmt.Errorf("could not read directory contents '%s'", err)
	}

	// Display the list of files and directories in the specified directory
	fmt.Printf("contents of directory '%s':\n", dirname)
	for _, fileInfo := range fileInfos {
		fmt.Println(fileInfo.Name())
	}

	return nil
}
