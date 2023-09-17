package common

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/dattaray-basab/cks-clip-lib/globals"
)

var GetFirstFileInRootDir = func(filePath string) (string, error) {
	firstFile := ""
	err := filepath.WalkDir(filePath, func(s string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() {
			firstFile = s
			return filepath.SkipDir
		}
		return nil
	})
	return firstFile, err
}

var GetFirstLineOfFirstFile = func(templateMap map[string]string) (string, error) {
	storePath := filepath.Join(templateMap[globals.KEY_ALTER_PATH], globals.STORE_DIRNAME)
	if !IsDir(storePath) {
		err := fmt.Errorf("store-path %s does not exist", storePath)
		return "", err
	}

	firstFile, err := GetFirstFileInRootDir(storePath)
	if err != nil {
		return "", err
	}

	wordList, err := GetWordsFromFile(firstFile)
	if err != nil {
		return "", err
	}
	firstWordInFirstFile := wordList[0]

	return firstWordInFirstFile, nil
}

var GetWordsFromFile = func(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)

	// Scan all words from the file.
	var words []string
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}
