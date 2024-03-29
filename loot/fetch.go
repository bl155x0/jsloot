package loot

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
)

type downloadResult struct {
	fileName   string
	downloaded bool
}

func FetchFromFile(filename string, beautify bool) error {
	exists, err := fileExists(filename)
	if err != nil {
		return err
	}
	if exists == false {
		return fmt.Errorf("no such file %s", filename)

	}
	files, err := readLinesFromFile(filename)
	if err != nil {
		return err
	}

	targetDirectory := filepath.Dir(filename)
	err = ensureDirectoryExists(targetDirectory)
	if err != nil {
		return err
	}
	FetchAll(files, targetDirectory, beautify)
	return nil
}

func FetchAll(urls []string, targetDirectory string, beautify bool) {
	doIfVerbose(func() {
		fmt.Println("Fetching into " + targetDirectory + "...")
	})
	for _, file := range urls {
		Fetch(file, targetDirectory, beautify)
	}
	doIfVerbose(func() {
		fmt.Println("Finished")
	})
}
func Fetch(url string, targetDirectory string, beautify bool) {
	result, err := fetchInternal(url, targetDirectory)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	} else {
		if result.downloaded == true {
			if IsVerbose() {
				fmt.Printf("Download complete: %s -> %s\n", url, result.fileName)
			} else {
				fmt.Printf("%s\n", result.fileName)
			}
			if beautify {
				err := beautifyFile(result.fileName)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error: %v\n", err)
				}
			}
		}
	}
}

func fetchInternal(urlString string, targetDirectory string) (*downloadResult, error) {
	//ensure schema is present
	if strings.HasPrefix(strings.ToLower(urlString), "https://") == false &&
		strings.HasPrefix(strings.ToLower(urlString), "http://") == false {
		urlString = "https://" + urlString
	}

	//parse the URL
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		return nil, err
	}

	//create the local filename for it
	localFileName := filepath.Base(parsedUrl.Path)
	if localFileName[0] == '/' {
		localFileName = localFileName[1:]
	}
	localDir := filepath.Join(targetDirectory, parsedUrl.Host)
	err = ensureDirectoryExists(localDir)
	if err != nil {
		return nil, err
	}
	localFileName = filepath.Join(localDir, localFileName)
	absFileName, err := filepath.Abs(localFileName)
	if err != nil {
		return nil, err
	}

	//Check if the file exists
	exists, err := fileExists(absFileName)
	if err != nil {
		return nil, err
	}
	if exists == true {
		skippedResult := downloadResult{
			fileName:   absFileName,
			downloaded: false,
		}
		return &skippedResult, nil
	}

	doIfVerbose(func() {
		fmt.Printf("Downloading: %s\n", parsedUrl.String())
	})
	response, err := http.Get(parsedUrl.String())
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file %s: %s\n", parsedUrl.String(), response.Status)
	}

	//Create the file
	downloadFile, err := os.Create(absFileName)
	if err != nil {
		return nil, err
	}
	defer downloadFile.Close()

	//copy
	_, err = io.Copy(downloadFile, response.Body)
	if err != nil {
		return nil, err
	}

	result := downloadResult{
		fileName:   absFileName,
		downloaded: true,
	}
	return &result, nil
}
