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

func ensureEndsWithPathSeparator(path string) string {
	// Check if the path already ends with os.PathSeparator
	if len(path) == 0 || path[len(path)-1] != os.PathSeparator {
		// If not, add os.PathSeparator to the end of the path
		path += string(os.PathSeparator)
	}
	return path
}

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

func createDirectory(dir string) error {
	return os.MkdirAll(dir, 0755)
}
