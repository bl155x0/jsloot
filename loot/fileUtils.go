package loot

import (
	"bufio"
	"os"
	"slices"
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

// createFile creates the given file
func createFile(filename string) error {
	file, err := os.Create(filename)
	defer file.Close()
	return err
}

// ensureExists creates the given file if not already existing
func ensureFileExists(filename string) error {
	exists, err := fileExists(filename)
	if err != nil {
		return err
	}
	if exists == false {
		return createFile(filename)
	}
	return nil
}

// reads all the lines in the given file
func readLinesFromFile(filename string) ([]string, error) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0)
	if err != nil {
		return []string{}, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, nil
}

// appendLinesToFile appends the lines to the given file
func appendLinesToFile(filename string, lines []string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range lines {
		_, err := writer.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}
	return writer.Flush()
}

// appendLinesToFileIfNotExist appends the lines to the given file if not already existing
func appendLinesToFileIfNotExist(filename string, lines []string) error {
	existingLines, err := readLinesFromFile(filename)
	if err != nil {
		return err
	}

	var linesToWrite []string
	for _, lineToWrite := range lines {
		if slices.Contains(existingLines, lineToWrite) == false {
			linesToWrite = append(linesToWrite, lineToWrite)
		}
	}

	return appendLinesToFile(filename, linesToWrite)
}
