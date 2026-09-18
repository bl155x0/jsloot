package loot

import (
	"fmt"
	"os"
)

// fileExists reports if the given file exists
func directoryExists(filename string) (bool, error) {
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return stat.IsDir(), nil
}

// ensureDirectoryExists creates the given directory if not already present
func ensureDirectoryExists(dir string) error {
	fileExists, err := fileExists(dir)
	if err != nil {
		return err
	}
	if fileExists {
		return fmt.Errorf("Cannot create directory %s: File exists", dir)
	}
	directoryExists, err := directoryExists(dir)
	if err != nil {
		return err
	}
	if directoryExists == false {
		err := createDirectory(dir)
		if err != nil {
			return err
		}
	}
	return nil
}

// createDirectory creates the given directory
func createDirectory(dir string) error {
	return os.MkdirAll(dir, 0755)
}
