package loot

import (
	"net/url"
	"path/filepath"
	"strings"
)

// resolveLocalFile parses urlString and returns the local file path it maps to under rootDirectory (by host), creating that directory if needed
func resolveLocalFile(urlString string, rootDirectory string) (*url.URL, string, error) {
	//ensure schema is present
	if strings.HasPrefix(strings.ToLower(urlString), "https://") == false &&
		strings.HasPrefix(strings.ToLower(urlString), "http://") == false {
		urlString = "https://" + urlString
	}

	//parse the URL
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return nil, "", err
	}

	//create the local filename for it
	localFileName := filepath.Base(parsedUrl.Path)
	if localFileName[0] == '/' {
		localFileName = localFileName[1:]
	}
	localDir := filepath.Join(rootDirectory, parsedUrl.Host)
	err = ensureDirectoryExists(localDir)
	if err != nil {
		return nil, "", err
	}
	localFileName = filepath.Join(localDir, localFileName)
	absFileName, err := filepath.Abs(localFileName)
	if err != nil {
		return nil, "", err
	}
	return parsedUrl, absFileName, nil
}
