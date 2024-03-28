package loot

const (
	defaultLootBoxFile = "js.loot"
)

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
