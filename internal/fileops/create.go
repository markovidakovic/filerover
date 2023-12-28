package fileops

import (
	"fmt"
	"os"
)

func CreateFile(filename string) error {
	_, err := os.Stat(filename)
	if !os.IsNotExist(err) {
		return fmt.Errorf("file '%s' already exists", filename)
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create file '%s'", filename)
	}
	defer file.Close()

	fmt.Printf("file '%s' created successfully\n", filename)
	return nil
}

func CreateDirectory(dirname string) error {
	_, err := os.Stat(dirname)
	if !os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' already exists", dirname)
	}

	// create the directory with read-write permissions for the current user
	err = os.Mkdir(dirname, 0755) // 0755 is the permission mode (rwxr-xr-x)
	if err != nil {
		return fmt.Errorf("could not create directory '%s'", dirname)
	}

	fmt.Printf("directory '%s' created successfully\n", dirname)
	return nil
}
