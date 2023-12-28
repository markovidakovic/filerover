package fileops

import (
	"fmt"
	"io"
	"os"
)

func CopyFile(source, destination string) error {
	// Open the source file
	srcFile, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("could not open source file '%s'", err)
	}
	defer srcFile.Close()

	// Create the destination file
	destFile, err := os.Create(destination)
	if err != nil {
		return fmt.Errorf("could not create destination file '%s'", err)
	}
	defer destFile.Close()

	// Copy the contents of the source file to the destination file
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("could not copy file contents '%s'", err)
	}

	fmt.Printf("file '%s' copied to '%s' successfully\n", source, destination)
	return nil
}

func CopyDirectory(source, destination string) error {
	// Open the source directory
	srcInfo, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("could not access source directory '%s'", err)
	}

	// Create the destination directory with the same permissions as the source directory
	if err := os.MkdirAll(destination, srcInfo.Mode()); err != nil {
		return fmt.Errorf("could not create destination directory '%s'", err)
	}

	// Read the contents of the source directory
	dirContents, err := os.ReadDir(source)
	if err != nil {
		return fmt.Errorf("could not read source directory contents '%s'", err)
	}

	// Copy files and subdirectories from the source to the destination directory
	for _, entry := range dirContents {
		srcPath := fmt.Sprintf("%s%s", source, entry.Name())
		destPath := fmt.Sprintf("%s%s", destination, entry.Name())

		if entry.IsDir() {
			if err := CopyDirectory(srcPath, destPath); err != nil {
				return err
			}
		} else {
			if err := CopyFile(srcPath, destPath); err != nil {
				return err
			}
		}
	}

	fmt.Printf("directory '%s' copied to '%s' successfully\n", source, destination)
	return nil
}

func CopyFileOrDirectory(source, destination string) error {
	// Stat the file/directory
	srcInfo, err := os.Stat(source)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file or directory '%s' does not exist", source)
		} else {
			return fmt.Errorf("could not access source file/directory '%s'", err)
		}
	}

	// Check if file or directory
	if srcInfo.Mode().IsRegular() {
		err := CopyFile(source, destination)
		if err != nil {
			return err
		}
		return nil
	} else if srcInfo.Mode().IsDir() {
		err := CopyDirectory(source, destination)
		if err != nil {
			return err
		}
		return nil
	} else {
		fmt.Println("source is neither a file nor a directory")
		return nil
	}
}
