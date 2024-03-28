package loot

import (
	"path"
)

const (
	defaultLootBoxFile = "js.loot"
)

// AddToDir adds the given URLs to the default lootbox file in the given dir
func AddToDir(lootboxDir string, urls []string) error {
	if len(urls) == 0 {
		return nil
	}
	err := ensureDirectoryExists(lootboxDir)
	if err != nil {
		return err
	}
	fileName := path.Join(lootboxDir, defaultLootBoxFile)
	return AddToFile(fileName, urls)
}

// AddToFile adds the given URLs to the lootbox file
func AddToFile(lootboxFile string, urls []string) error {
	if len(urls) == 0 {
		return nil
	}
	err := ensureFileExists(lootboxFile)
	if err != nil {
		return err
	}
	return appendLinesToFileIfNotExist(lootboxFile, urls)
}
