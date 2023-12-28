package fileops

import (
	"fmt"
	"os"
)

func DeleteFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return fmt.Errorf("file '%s' does not exist", filename)
	}

	if err := os.Remove(filename); err != nil {
		return fmt.Errorf("could not delete file '%s'", err)
	}

	fmt.Printf("file '%s' deleted successfully\n", filename)
	return nil
}

func DeleteDirectory(dirname string) error {
	_, err := os.Stat(dirname)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory '%s' does not exist", dirname)
	}

	if err := os.RemoveAll(dirname); err != nil {
		return fmt.Errorf("could not delete directory '%s'", err)
	}

	fmt.Printf("directory '%s' deleted successfully\n", dirname)
	return nil
}
