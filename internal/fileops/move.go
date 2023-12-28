package fileops

import (
	"fmt"
	"os"
)

func MoveFile(source, destination string) error {
	// Check if the source file exists
	_, err := os.Stat(source)
	if os.IsNotExist(err) {
		return fmt.Errorf("source file '%s' does not exist", source)
	}

	// Check if the destination file already exists
	_, err = os.Stat(destination)
	if !os.IsNotExist(err) {
		return fmt.Errorf("destination file '%s' already exists", destination)
	}

	// Rename the source file to the destination file path
	err = os.Rename(source, destination)
	if err != nil {
		return fmt.Errorf("could not move file '%s'", err)
	}

	fmt.Printf("file '%s' moved to '%s' successfully\n", source, destination)
	return nil
}

func MoveDirectory(source, destination string) error {
	// Check if the source directory exists
	_, err := os.Stat(source)
	if os.IsNotExist(err) {
		return fmt.Errorf("source directory '%s' does not exist", source)
	}

	// Check if the destination directory already exists
	_, err = os.Stat(destination)
	if !os.IsNotExist(err) {
		return fmt.Errorf("destination directory '%s' already exists", destination)
	}

	// Rename the source directory to the destination directory path
	err = os.Rename(source, destination)
	if err != nil {
		return fmt.Errorf("could not move directory '%s'", err)
	}

	fmt.Printf("directory '%s' moved to '%s' successfully\n", source, destination)
	return nil
}

func MoveFileOrDirectory(source, destination string) error {
	// Check if the source file/directory exists
	_, err := os.Stat(source)
	if os.IsNotExist(err) {
		return fmt.Errorf("source file or directory '%s' does not exist", source)
	}

	// Check if the destination file/directory already exists
	_, err = os.Stat(destination)
	if err != nil {
		return fmt.Errorf("could not move file or directory '%s'", err)
	}

	// Rename the source directory to the destination path
	err = os.Rename(source, destination)
	if err != nil {
		return fmt.Errorf("could not move file or directory '%s'", err)
	}

	fmt.Printf("file or directory '%s' moved to '%s' successfully\n", source, destination)
	return nil
}
