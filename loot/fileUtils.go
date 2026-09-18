package loot

import (
	"os"
)

// fileExists reports if the given file exitst
func fileExists(filename string) (bool, error) {
	stat, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false, nil
	} else if err != nil {
		return false, err
	}
	return stat.IsDir() == false, nil
}
