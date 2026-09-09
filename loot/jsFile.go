package loot

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

// JSFile is a JavaScript file to store, together with the URL it was downloaded from
type JSFile struct {
	URL     string `json:"URL"`
	Content string `json:"content"`
}

// ParseJSFile reads and parses a JSFile as JSON from the given stream
func ParseJSFile(r io.Reader) (JSFile, error) {
	var jsFile JSFile
	err := json.NewDecoder(r).Decode(&jsFile)
	if err != nil {
		return jsFile, err
	}
	err = validateJSFile(jsFile)
	return jsFile, err
}

// validateJSFile ensures all mandatory fields of the JSFile are set
func validateJSFile(f JSFile) error {
	if f.URL == "" {
		return fmt.Errorf("missing mandatory field: URL")
	}
	if f.Content == "" {
		return fmt.Errorf("missing mandatory field: content")
	}
	return nil
}

// StoreJSFile writes the JSFile's content to the local path its URL resolves to under rootDirectory,
// the same path Fetch would have downloaded it to. Returns the path written to.
// An existing file at that path is overwritten. If beautify is true, the stored file is run through
// js-beautify, the same as Fetch does for downloaded files.
func StoreJSFile(f JSFile, rootDirectory string, beautify bool) (string, error) {
	_, absFileName, err := resolveLocalFile(f.URL, rootDirectory)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(absFileName, []byte(f.Content), 0644)
	if err != nil {
		return "", err
	}

	if beautify {
		err = beautifyFile(absFileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		}
	}
	return absFileName, nil
}
